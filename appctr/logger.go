package appctr

import (
	"go.uber.org/zap"
	"log"
)

func Log() *zap.Logger {
	return &lg
}

var lg zap.Logger

func prepareLog() {
	var err error
	var l *zap.Logger

	l, err = zap.NewDevelopment()

	if err != nil {
		log.Fatalf("logger can't be instanced")
	}

	lg = *l
}
