package datasource

import (
	"github.com/chipx/go-translator/internal"
	"time"
)

type DataSource interface {
	LoadAll(criteria Criteria) (map[string]*internal.Vocabulary, error)
	GetLastModified() time.Time
	Get(lang string, key string) (string, error)
	Set(lang string, key string, msg string) error
	Delete(lang string, key string) error
	MarkAsUntranslated(lang string, key string) error
}

type Criteria struct {
	KeyPrefix string
	Since     time.Time
	Langs     []string
	Translated Translated
	Limit int
	Page int
}

func (c *Criteria) GetOffset() int {
	if c.Page > 0 && c.Limit > 0 {
		return c.Page * c.Limit
	}

	return 0
}

type Translated int

const (
	NOT_SET Translated = iota
	YES
	NO
)
