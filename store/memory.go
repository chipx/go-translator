package store

import (
	"fmt"
	"github.com/chipx/go-translator/datasource"
	"github.com/chipx/go-translator/internal"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func NewMemoryStore(source datasource.DataSource, updatePeriod time.Duration) (*Memory, error) {
	var err error
	store := &Memory{source: source}
	store.catalog, err = source.LoadAll(datasource.Criteria{})
	if err != nil {
		return nil, err
	}

	if updatePeriod > 0 {
		lastModifiedFromSource := source.GetLastModified()
		go func() {
			ticker := time.NewTicker(updatePeriod)
			for {
				<-ticker.C

			}
		}()
	}

	return store, nil
}

type Memory struct {
	source       datasource.DataSource
	catalog      map[string]*internal.Vocabulary
	mutex        sync.RWMutex
	lastModified time.Time
}

func (m *Memory) Get(lang string, key string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if voc, ok := m.catalog[lang]; ok {
		if data, ok := voc.Lookup(key); ok {
			return data, nil
		}
	}

	data, err := m.source.Get(lang, key)
	if err == nil {
		m.set(lang, key, data)
		return data, nil
	}

	return "", fmt.Errorf(ErrorStoreDataNotFound)
}

func (m *Memory) update() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	modifiedFromSource := m.source.GetLastModified()
	if !modifiedFromSource.After(m.lastModified) {
		return nil
	}

	data, err := m.source.LoadAll(datasource.Criteria{Since: m.lastModified})
	if err != nil {
		return err
	}

	m.lastModified = modifiedFromSource
	m.catalog = data

	return nil
}

func (m *Memory) set(lang string, key string, data string) {
	if _, ok := m.catalog[lang]; !ok {
		m.catalog[lang] = internal.NewVocabulary(lang)
	}

	m.catalog[lang].Set(key, data)
}

func (m *Memory) Set(lang string, key string, data string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	err := m.source.Set(lang, key, data)
	if err != nil {
		return err
	}

	m.set(lang, key, data)
	return nil
}

func (m *Memory) MarkAsUntranslated(lang string, key string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	err := m.source.MarkAsUntranslated(lang, key)
	if err != nil {
		return err
	}

	m.set(lang, key, key)
	return nil
}
