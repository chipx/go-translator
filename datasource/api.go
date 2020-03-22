package datasource

import (
	"context"
	"errors"
	"github.com/chipx/go-translator/api"
	"github.com/chipx/go-translator/internal"
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const defaultApiQueryLimit = 1000

var ErrorApiClientNotInitialized = errors.New("Api client did not initialized ")
var ErrorApiRequest = errors.New("Api request failed ")

func NewApiDataSource(path string, keepAlive time.Duration, grpcOpts *[]grpc.DialOption) DataSource {
	source := &apiDataSource{
		path:      path,
		keepAlive: keepAlive,
	}

	if grpcOpts != nil {
		source.opts = *grpcOpts
	}

	return source
}

type apiDataSource struct {
	path      string
	opts      []grpc.DialOption
	client    api.TranslatorClient
	keepAlive time.Duration
	connMutex sync.Mutex
	wg        sync.WaitGroup
}

func (a *apiDataSource) getClient() api.TranslatorClient {
	a.connMutex.Lock()
	defer a.connMutex.Unlock()

	a.wg.Add(1)

	if a.client != nil {
		return a.client
	}

	conn, err := grpc.Dial(a.path, a.opts...)
	if err != nil {
		log.WithError(err).Fatalf("Dial %s with opts %v failed", a.path, a.opts)
		return nil
	}

	log.Debugf("Connection to %s created", a.path)

	a.client = api.NewTranslatorClient(conn)

	if a.keepAlive > 0 {
		go func() {
			timer := time.NewTimer(a.keepAlive)
			<-timer.C
			a.connMutex.Lock()
			a.wg.Wait()
			defer a.connMutex.Unlock()

			err := conn.Close()
			if err != nil {
				log.WithError(err).Errorf("Close connection failed")
			} else {
				log.Debugf("Connection to %s closed", a.path)
			}
			a.client = nil
		}()
	}

	return a.client
}

func (a *apiDataSource) LoadAll(criteria Criteria) (map[string]*internal.Vocabulary, error) {
	defer a.wg.Done()

	client := a.getClient()
	if client == nil {
		return nil, ErrorApiClientNotInitialized
	}

	req := &api.GetAllRequest{
		KeyPrefix:  criteria.KeyPrefix,
		Langs:      criteria.Langs,
		Trunslated: int32(criteria.Translated),
		Limit:      int32(criteria.Limit),
		Page:       int32(criteria.Page),
	}

	if !criteria.Since.IsZero() {
		req.Since = &timestamp.Timestamp{
			Seconds: criteria.Since.Unix(),
			Nanos:   0,
		}
	}

	neePagination := false
	if criteria.Limit < 1 {
		req.Limit = defaultApiQueryLimit
		req.Page = 0
		neePagination = true
	}

	ctl := make(map[string]*internal.Vocabulary)

	for {
		resp, err := client.GetAll(context.Background(), req)
		if err != nil {
			log.WithError(err).Errorf("Get all by {%v} failed", req)
			return nil, ErrorApiRequest
		}

		if len(resp.List) == 0 {
			break
		}

		for v := 0; v < len(resp.List); v++ {
			if _, ok := ctl[resp.List[v].Lang]; !ok {
				ctl[resp.List[v].Lang] = internal.NewVocabulary(resp.List[v].Lang)
			}

			for i := 0; i < len(resp.List[v].Data); i++ {
				ctl[resp.List[v].Lang].Set(resp.List[v].Data[i].Key, resp.List[v].Data[i].Message)
			}
		}

		if neePagination {
			req.Page += 1
		} else {
			break
		}
	}

	return ctl, nil
}
func (a *apiDataSource) GetLastModified() time.Time {
	defer a.wg.Done()

	client := a.getClient()
	if client == nil {
		return time.Unix(0, 0)
	}

	modifiedTime, err := client.GetLastModified(context.Background(), &api.SimpleRequest{})
	if err != nil {
		log.WithError(err).Errorf("Get last modified failed")
		return time.Unix(0, 0)
	}
	log.Debugf("modifiedTime - %v", modifiedTime.Seconds)
	return time.Unix(modifiedTime.Seconds, 0)
}
func (a *apiDataSource) Get(lang string, key string) (string, error) {
	defer a.wg.Done()

	client := a.getClient()
	if client == nil {
		return "", ErrorApiClientNotInitialized
	}

	resp, err := client.Get(context.Background(), &api.TranslateRequest{
		Lang: lang,
		Key:  key,
	})

	if err != nil {
		log.WithError(err).Errorf("Get translate for lang %s by key %s failed", lang, key)
		return "", ErrorApiRequest
	}

	return resp.Message, nil
}

func (a *apiDataSource) Set(lang string, key string, msg string) error {
	defer a.wg.Done()

	client := a.getClient()
	if client == nil {
		return ErrorApiClientNotInitialized
	}

	_, err := client.Set(context.Background(), &api.SetTranslateRequest{
		Lang:    lang,
		Key:     key,
		Message: msg,
	})

	if err != nil {
		log.WithError(err).Errorf("Set for lang %s by key %s failed", lang, key)
		return ErrorApiRequest
	}

	return nil
}
func (a *apiDataSource) Delete(lang string, key string) error {
	log.Errorf("Delete translate for remote data source not implemented")
	return nil
}
func (a *apiDataSource) MarkAsUntranslated(lang string, key string) error {
	defer a.wg.Done()

	client := a.getClient()
	if client == nil {
		return ErrorApiClientNotInitialized
	}

	_, err := client.MarkAsUntranslated(context.Background(), &api.TranslateRequest{
		Lang: lang,
		Key:  key,
	})

	if err != nil {
		log.WithError(err).Errorf("Mark as untranslated for lang %s by key %s failed", lang, key)
		return ErrorApiRequest
	}

	return nil
}
