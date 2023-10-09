package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/EdwardKerckhof/gohtmx/config"
)

type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type logger struct {
	config      *config.Config
	sugarLogger *zap.SugaredLogger
}

func New(config *config.Config) *logger {
	return &logger{config: config}
}

func (l *logger) getLoggerLevel(config *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[config.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *logger) InitLogger() {
	logLevel := l.getLoggerLevel(l.config)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderconfig zapcore.EncoderConfig
	if l.config.Api.Mode == "development" {
		encoderconfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderconfig = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderconfig.LevelKey = "LEVEL"
	encoderconfig.CallerKey = "CALLER"
	encoderconfig.TimeKey = "TIME"
	encoderconfig.NameKey = "NAME"
	encoderconfig.MessageKey = "MESSAGE"

	if l.config.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderconfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderconfig)
	}

	encoderconfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

func (l *logger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	log.Printf(template, args...)
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
