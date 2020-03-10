package datasource

import (
	"github.com/chipx/go-translator/internal"
	"time"
)

type DataSource interface {
	LoadAll() (map[string]*internal.Vocabulary, error)
	LoadLast(since time.Time) (map[string]*internal.Vocabulary, error)
	GetLastModified() time.Time
	Get(lang string, key string) (string, error)
	Set(lang string, key string, msg string) error
	Delete(lang string, key string) error
	MarkAsUntranslated(lang string, key string) error
}
