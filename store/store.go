package store

const (
	ErrorStoreDataNotFound = "Data did not found in store and data source "
)

type Store interface {
	Get(lang string, key string) (string, error)
	Set(lang string, key string, data string) error
	MarkAsUntranslated(lang string, key string) error
}
