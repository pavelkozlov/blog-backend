package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"blog/pkg/config"
)

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

type logger struct {
	zap *zap.SugaredLogger
}

func (l logger) Info(msg string, keysAndValues ...interface{}) {
	l.zap.Infow(msg, keysAndValues...)
}

func (l logger) Debug(msg string, keysAndValues ...interface{}) {
	l.zap.Debugw(msg, keysAndValues...)
}

func (l logger) Warn(msg string, keysAndValues ...interface{}) {
	l.zap.Warnw(msg, keysAndValues...)
}

func (l logger) Error(msg string, keysAndValues ...interface{}) {
	l.zap.Errorw(msg, keysAndValues)
}

func NewLogger(config *config.Config) Logger {

	level := zap.InfoLevel
	if config.Server.DebugMode {
		level = zap.DebugLevel
	}

	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, err := conf.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal("Can not init logger: " + err.Error())
	}

	sugar := zapLogger.Sugar()

	return &logger{
		zap: sugar,
	}
}
