package transaltor

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"transaltor/store"
)

var defaultStore store.Store

func InitTranslator(store store.Store) {
	defaultStore = store
}

func Translate(lang string, message string) string {
	storedTranslate, err := defaultStore.Get(lang, message)
	if err != nil {
		logrus.WithError(err).Warningf("Get message %s:%s from store failed", lang, message)
	}

	if storedTranslate != nil {
		switch res := storedTranslate.(type) {
		case string:
			return res
		default:
			logrus.Errorf("Can not translate cached message (%v) of type %T", storedTranslate, res)
			return message
		}
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
