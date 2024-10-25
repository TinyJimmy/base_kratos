/*
*description: 初始化日志
*creator: zhangce
*date: 2024-10-23
 */
package zlog

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ log.Logger = (*ZapLogger)(nil)
var ServiceLogger *zap.Logger
var RouterLogger *zap.Logger

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	slog *zap.Logger
	rlog *zap.Logger
	Sync func() error
}

func InitZapLogger() *ZapLogger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger := NewZapLogger(cfg,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2))
	ServiceLogger = logger.slog
	RouterLogger = logger.rlog
	return logger
}

// NewZapLogger return a zap logger.
func NewZapLogger(encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), level)
	zapLogger := zap.New(core, opts...)
	serviceWriteSyncer := getServiceLogWriter()
	routerWriteSyncer := getRouterLogWriter()
	serviceCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), serviceWriteSyncer, zapcore.DebugLevel)
	routerCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), routerWriteSyncer, zapcore.DebugLevel)
	serviceLogger := zap.New(serviceCore, zap.AddCaller())
	routerLogger := zap.New(routerCore, zap.AddCaller())
	return &ZapLogger{log: zapLogger, slog: serviceLogger, rlog: routerLogger, Sync: zapLogger.Sync}
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}

func getServiceLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("../../internal/zlog/service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return zapcore.AddSync(file)
}

func getRouterLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("../../internal/zlog/router.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return zapcore.AddSync(file)
}
