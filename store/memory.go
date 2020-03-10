package store

import (
	"fmt"
	"go-transaltor/datasource"
	"go-transaltor/internal"
	"sync"
)

func NewMemoryStore(source datasource.DataSource) (*Memory, error) {
	var err error
	store := &Memory{source: source}
	store.catalog, err = source.LoadAll()
	if err != nil {
		return nil, err
	}
	store.mutex = sync.RWMutex{}

	return store, nil
}

type Memory struct {
	source  datasource.DataSource
	catalog map[string]*internal.Vocabulary
	mutex   sync.RWMutex
}

func (m *Memory) Get(lang string, key string) (interface{}, error) {
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

	return nil, fmt.Errorf(ErrorStoreDataNotFound)
}

func (m *Memory) set(lang string, key string, data interface{}) {
	if _, ok := m.catalog[lang]; !ok {
		m.catalog[lang] = internal.NewVocabulary(lang)
	}

	m.catalog[lang].Set(key, data.(string))
}

func (m *Memory) Set(lang string, key string, data interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	err := m.source.Set(lang, key, data.(string))
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
