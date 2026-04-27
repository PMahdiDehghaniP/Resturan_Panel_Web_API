package logging

import (
	"sync"

	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
)

type Logger interface {
	Init()

	Debug(category Category, subCategory SubCategory, message string, extraKey map[ExtraKey]interface{})
	DebugF(template string, args ...interface{})

	Info(category Category, subCategory SubCategory, message string, extraKey map[ExtraKey]interface{})
	InfoF(template string, args ...interface{})

	Warn(category Category, subCategory SubCategory, message string, extraKey map[ExtraKey]interface{})
	WarnF(template string, args ...interface{})

	Error(category Category, subCategory SubCategory, message string, extraKey map[ExtraKey]interface{})
	ErrorF(template string, args ...interface{})

	Fatal(category Category, subCategory SubCategory, message string, extraKey map[ExtraKey]interface{})
	FatalF(template string, args ...interface{})
}

var (
	logger Logger
	once   sync.Once
)

func NewLogger(cfg *config.Config) Logger {
	once.Do(func() {
		switch cfg.Logger.LoggerName {
		case "Zap":
			logger = NewZapLogger(cfg)
		case "Zero":
			logger = NewZeroLogger(cfg)
		default:
			panic("invalid logger name")
		}
	})
	return logger
}
