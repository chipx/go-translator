package store

import (
	"fmt"
	"github.com/chipx/go-translator/datasource"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func NewRedisStore(redis *redis.Client, source datasource.DataSource, updatePeriod time.Duration) (*Redis, error) {
	store := &Redis{redis: redis, source: source}
	store.update()
	go func() {
		ticker := time.NewTicker(updatePeriod)
		for {
			<-ticker.C
			store.update()
		}
	}()

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

	modifiedFromSource := r.source.GetLastModified()
	if !modifiedFromSource.After(lastModified) {
		return nil
	}

	data, err := r.source.LoadAll(datasource.Criteria{Since: lastModified})
	if err != nil {
		return err
	}

	for lang := range data {
		for key, message := range data[lang].AsMap() {
			r.set(lang, key, message)
		}
	}

	r.redis.Set(lastModifiedKey, modifiedFromSource.Unix(), 0)

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

func (r *Redis) set(lang string, key string, data string) error {
	storeKey := fmt.Sprintf(itemKeyTemplate, lang, key)
	return r.redis.Set(storeKey, data, 0).Err()
}

func (r *Redis) Get(lang string, key string) (string, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	storeKey := fmt.Sprintf(itemKeyTemplate, lang, key)
	msg, err := r.redis.Get(storeKey).Result()
	if err == nil {
		return msg, nil
	}

	msg, err = r.source.Get(lang, key)
	if err != nil {
		return "", err
	}

	r.set(lang, key, msg)

	return msg, nil
}

func (r *Redis) Set(lang string, key string, data string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.source.Set(lang, key, data); err != nil {
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
