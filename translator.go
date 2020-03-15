package transaltor

import (
	"fmt"
	"github.com/chipx/go-translator/store"
	"github.com/sirupsen/logrus"
)

var defaultStore store.Store

func InitTranslator(store store.Store) {
	defaultStore = store
}

func Translate(lang string, message string) string {
	storedTranslate, err := defaultStore.Get(lang, message)
	if err != nil {
		logrus.WithError(err).Warningf("Get message %s:%s from store failed", lang, message)
	} else if storedTranslate != "" {
		return storedTranslate
	}

	if err := defaultStore.MarkAsUntranslated(lang, message); err != nil {
		logrus.WithError(err).Errorf("Mark as untranslated message %s:%s failed", lang, message)
	}

	return message
}

func Translatef(lang string, message string, args ...interface{}) string {
	storedMessage := Translate(lang, message)
	return fmt.Sprintf(storedMessage, args...)
}
