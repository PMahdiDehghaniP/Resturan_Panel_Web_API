package logging

import (
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ZapLogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type ZapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *ZapLogger {
	logger := &ZapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func prepareLogKeys(extra map[ExtraKey]interface{}, category Category, subCategory SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}

	extra["Category"] = category
	extra["SubCategory"] = subCategory

	params := mapToZapParams(extra)
	return params
}

func (logger *ZapLogger) GetLogLevel() zapcore.Level {
	level, exists := ZapLogLevelMap[logger.cfg.Logger.Level]
	if !exists {
		level = zapcore.DebugLevel
	}
	return level
}

func (logger *ZapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logger.cfg.Logger.FilePath,
		MaxSize:    10,
		MaxAge:     5,
		LocalTime:  true,
		MaxBackups: 10,
		Compress:   true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w, logger.GetLogLevel(),
	)
	customLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel)).Sugar()
	customLogger = customLogger.With("App Name", "Restaurant Web App", "LoggerName", "ZapLog")
	logger.logger = customLogger
}

func (logger *ZapLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	logger.logger.Debugw(message, params...)
}

func (logger *ZapLogger) DebugF(template string, arg ...interface{}) {
	logger.logger.Debugf(template, arg...)
}

func (logger *ZapLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	logger.logger.Infow(message, params...)
}

func (logger *ZapLogger) InfoF(template string, arg ...interface{}) {
	logger.logger.Infof(template, arg...)
}

func (logger *ZapLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	logger.logger.Warnw(message, params...)
}

func (logger *ZapLogger) WarnF(template string, arg ...interface{}) {
	logger.logger.Warnf(template, arg...)
}

func (logger *ZapLogger) Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	logger.logger.Errorw(message, params...)
}

func (logger *ZapLogger) ErrorF(template string, arg ...interface{}) {
	logger.logger.Errorf(template, arg...)
}

func (logger *ZapLogger) Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	logger.logger.Fatalw(message, params...)
}

func (logger *ZapLogger) FatalF(template string, arg ...interface{}) {
	logger.logger.Fatalf(template, arg...)
}
