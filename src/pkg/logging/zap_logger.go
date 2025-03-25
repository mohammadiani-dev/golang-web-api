package logging

import (
	"golang-web-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLoggerInstance *zap.SugaredLogger

type zapLogger struct {
	cfg *config.Logger
	logger *zap.SugaredLogger
}

var LogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func newZapLogger(cfg *config.Logger) *zapLogger {
	Logger := &zapLogger{cfg: cfg}
	Logger.Init()
	return Logger
}

func (l *zapLogger) getLogLevel(level string) zapcore.Level {
	if logLevel, ok := LogLevelMap[level]; ok {
		return logLevel
	}
	return zapcore.DebugLevel
}

func (l *zapLogger) Init() {

	once.Do(func() {
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename: l.cfg.FilePath,
			MaxSize: 1,
			MaxBackups: 10,
			MaxAge: 5,
			Compress: true,
		})

		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			writer,
			zap.NewAtomicLevelAt(l.getLogLevel(l.cfg.Level)),
		)

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1) , zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

		zapLoggerInstance = logger.With("Appname", "car app","LoggerName", "zaplogger")
	})
	l.logger = zapLoggerInstance
	
}

func (l *zapLogger) Info(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	params := prepareLogKeys(extera, cat, sub)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}	

func (l *zapLogger) Debug(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	params := prepareLogKeys(extera, cat, sub)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Warn(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	params := prepareLogKeys(extera, cat, sub)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}


func (l *zapLogger) Error(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	params := prepareLogKeys(extera, cat, sub)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}


func (l *zapLogger) Fatal(cat Category, sub subCategory, msg string, extera map[ExtraKey]interface{}) {
	params := prepareLogKeys(extera, cat, sub)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}



func prepareLogKeys(extra map[ExtraKey]interface{},cat Category,sub subCategory) []interface{} {
	if(extra == nil) {
		extra = make(map[ExtraKey]interface{})
	}
	extra["category"] = cat
	extra["subCategory"] = sub
	params := mapTozapParams(extra)

	return params
}