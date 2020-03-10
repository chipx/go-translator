package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
	"transaltor/datasource"
)

func NewRedisStore(redis *redis.Client, source datasource.DataSource) (*Redis, error) {
	store := &Redis{redis: redis, source: source}
	store.update()
	store.mutex = sync.RWMutex{}

	return store, nil
}

const (
	lastModifiedKey = "last-modified"
	itemKeyTemplate = "%s:%s"
)

type Redis struct {
	redis  *redis.Client
	source datasource.DataSource
	mutex  sync.RWMutex
}

func (r *Redis) update() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	lastModified := r.getLastModified()
	if !r.source.GetLastModified().After(lastModified) {
		return nil
	}

	data, err := r.source.LoadLast(lastModified)
	if err != nil {
		return err
	}

	for lang := range data {
		for key, message := range data[lang].AsMap() {
			r.set(lang, key, message)
		}
	}

	return nil
}

func (r *Redis) getLastModified() time.Time {
	cmd := r.redis.Get(lastModifiedKey)
	unixTimeSec, err := cmd.Int64()
	if err != nil {
		return time.Unix(0, 0)
	}

	return time.Unix(unixTimeSec, 0)
}

func (r *Redis) set(lang string, key string, data interface{}) error {
	storeKey := fmt.Sprintf(itemKeyTemplate, lang, key)
	return r.redis.Set(storeKey, data, 0).Err()
}

func (r *Redis) Get(lang string, key string) (interface{}, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	storeKey := fmt.Sprintf(itemKeyTemplate, lang, key)
	msg, err := r.redis.Get(storeKey).Result()
	if err == nil {
		return msg, nil
	}

	msg, err = r.source.Get(lang, key)
	if err != nil {
		return nil, err
	}

	r.set(lang, key, msg)

	return msg, nil
}

func (r *Redis) Set(lang string, key string, data interface{}) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.source.Set(lang, key, data.(string)); err != nil {
		return err
	}

	return r.set(lang, key, data)
}

func (r *Redis) MarkAsUntranslated(lang string, key string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.set(lang, key, key)
	return r.source.MarkAsUntranslated(lang, key)
}
