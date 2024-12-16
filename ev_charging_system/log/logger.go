package log

import (
	"ev_charging_system/config"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *CustomLogger

type CustomLogger struct {
	*zap.SugaredLogger

	writer io.Writer
}

func InitLogger(name string) {
	writeSyncer := getLogWriter(name)
	encoder := getEncoder()
	if config.ChargeConfig.LogLevel == "" {
		panic("LogLevel null")
	}
	var core zapcore.Core
	switch config.ChargeConfig.LogLevel {
	case "Debug":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	case "Info":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	case "Warn":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.WarnLevel)
	case "Error":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.ErrorLevel)
	case "DPanic":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DPanicLevel)
	case "Panic":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.PanicLevel)
	case "Fatal":
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.FatalLevel)
	}

	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = &CustomLogger{l.Sugar(), writeSyncer}

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	// 如果做日志分析系统就用JSON格式
	//return zapcore.NewJSONEncoder(encoderConfig)

}

func getLogWriter(name string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./%s.log", name),
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     15,
		Compress:   false,
	}

	var syncer zapcore.WriteSyncer
	if config.IsConsoleMode() {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	} else {
		syncer = zapcore.AddSync(lumberJackLogger)
	}

	return syncer
}

func Sync() {
	logger.Sync()
}

func Error(args ...interface{}) {
	logger.SugaredLogger.Error(args)
}

func Errorf(template string, args ...interface{}) {
	logger.SugaredLogger.Errorf(template, args...)
}
func Info(args ...interface{}) {
	logger.SugaredLogger.Info(args)
}

func Infof(template string, args ...interface{}) {
	logger.SugaredLogger.Infof(template, args...)
}

func Debug(args ...interface{}) {
	logger.SugaredLogger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	logger.SugaredLogger.Debugf(template, args...)
}

func Warning(args ...interface{}) {
	logger.SugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.SugaredLogger.Warnf(template, args...)
}

func LoggerWriter() io.Writer {
	return logger.writer
}

func CronPanicLogger() *CustomLogger {
	return logger
}

func (c *CustomLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Infow(msg, keysAndValues)
}

func (c *CustomLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.SugaredLogger.Errorw(err.Error(), msg, keysAndValues)
}
