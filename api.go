package transaltor

import (
	"context"
	"errors"
	"fmt"
	"github.com/chipx/go-translator/api"
	"github.com/chipx/go-translator/datasource"
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"time"
)

var ErrorNotFound = errors.New("Not found translate ")
var ErrorInternal = errors.New("Internal error ")

func NewApiServer(dataSource datasource.DataSource) api.TranslatorServer {
	return &apiServer{dataSource: dataSource}
}

type apiServer struct {
	dataSource datasource.DataSource
}

func (s *apiServer) Get(ctx context.Context, req *api.TranslateRequest) (*api.SimpleResponse, error) {
	translate, err := s.dataSource.Get(req.Lang, req.Key)
	if err != nil {
		log.WithError(err).Errorf("Get data for %v failed", *req)
		return nil, ErrorNotFound
	}

	return &api.SimpleResponse{
		Success: true,
		Message: translate,
	}, nil
}

func (s *apiServer) GetAll(ctx context.Context, req *api.GetAllRequest) (*api.GetAllResponse, error) {
	criteria := datasource.Criteria{
		KeyPrefix: req.KeyPrefix,
		Langs:     req.Langs,
		Limit:     int(req.GetLimit()),
		Page:      int(req.GetPage()),
	}

	if req.GetSince() != nil {
		criteria.Since = time.Unix(req.GetSince().Seconds, 0)
	}

	if req.Trunslated == int32(datasource.YES) || req.Trunslated == int32(datasource.NO) {
		criteria.Translated = datasource.Translated(req.Trunslated)
	}

	data, err := s.dataSource.LoadAll(criteria)
	if err != nil {
		log.WithError(err).Errorf("Load data for request %v failed", *req)
		return nil, ErrorNotFound
	}

	resp := &api.GetAllResponse{
		List: make([]*api.Vocabulary, 0, len(data)),
	}

	for langName := range data {
		voc := &api.Vocabulary{
			Lang: langName,
			Data: make([]*api.Translate, 0, len(data[langName].AsMap())),
		}

		for key, message := range data[langName].AsMap() {
			voc.Data = append(voc.Data, &api.Translate{
				Lang:    langName,
				Key:     key,
				Message: message,
			})
		}

		resp.List = append(resp.List, voc)
	}
	fmt.Println("-------")
	fmt.Println(resp)
	fmt.Println("-------")
	return resp, nil
}

func (s *apiServer) MarkAsUntranslated(ctx context.Context, req *api.TranslateRequest) (*api.SimpleResponse, error) {
	err := s.dataSource.MarkAsUntranslated(req.Lang, req.Key)
	if err != nil {
		log.WithError(err).Errorf("Mark as untranslated for %v failed", *req)
		return nil, ErrorInternal
	}

	return &api.SimpleResponse{
		Success: true,
		Message: req.Key,
	}, nil
}

func (s *apiServer) GetLastModified(ctx context.Context, req *api.SimpleRequest) (*timestamp.Timestamp, error) {
	log.Debug("LastModified: ", s.dataSource.GetLastModified())
	return &timestamp.Timestamp{
		Seconds: s.dataSource.GetLastModified().Unix(),
		Nanos:   0,
	}, nil
}
