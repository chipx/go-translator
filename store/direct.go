package store

import (
	"fmt"
	"github.com/chipx/go-translator/datasource"
	"sync"
)

func NewDirectStore(source datasource.DataSource) *Direct {
	store := &Direct{source: source}
	store.mutex = sync.RWMutex{}

	return store
}

type Direct struct {
	source datasource.DataSource
	mutex  sync.RWMutex
}

func (m *Direct) Get(lang string, key string) (interface{}, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	data, err := m.source.Get(lang, key)
	if err == nil {
		return data, nil
	}

	return nil, fmt.Errorf(ErrorStoreDataNotFound)
}

func (m *Direct) Set(lang string, key string, data interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	err := m.source.Set(lang, key, data.(string))
	if err != nil {
		return err
	}

	return nil
}

func (m *Direct) MarkAsUntranslated(lang string, key string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	err := m.source.MarkAsUntranslated(lang, key)
	if err != nil {
		return err
	}

	return nil
}
