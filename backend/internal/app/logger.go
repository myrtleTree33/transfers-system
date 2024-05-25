package app

import (
	"context"

	"go.uber.org/zap"
)

type ILogger interface {
	InfofContext(ctx context.Context, format string, args ...interface{})
	WarnfContext(ctx context.Context, format string, args ...interface{})
	ErrorfContext(ctx context.Context, format string, args ...interface{})
	FatalfContext(ctx context.Context, format string, args ...interface{})
	PanicfContext(ctx context.Context, format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type LoggerT struct {
	sugaredLogger *zap.SugaredLogger
}

func NewLogger() ILogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	ll := &LoggerT{
		sugaredLogger: logger.Sugar(),
	}

	return ll
}

func (l *LoggerT) InfofContext(ctx context.Context, format string, args ...interface{}) {
	fmtStr := "[context=%s] " + format
	l.sugaredLogger.Infof(fmtStr, append([]interface{}{ctx.Value("context")}, args...)...)
}

func (l *LoggerT) WarnfContext(ctx context.Context, format string, args ...interface{}) {
	fmtStr := "[context=%s] " + format
	l.sugaredLogger.Warnf(fmtStr, append([]interface{}{ctx.Value("context")}, args...)...)
}

func (l *LoggerT) ErrorfContext(ctx context.Context, format string, args ...interface{}) {
	fmtStr := "[context=%s] " + format
	l.sugaredLogger.Errorf(fmtStr, append([]interface{}{ctx.Value("context")}, args...)...)
}

func (l *LoggerT) FatalfContext(ctx context.Context, format string, args ...interface{}) {
	fmtStr := "[context=%s] " + format
	l.sugaredLogger.Fatalf(fmtStr, append([]interface{}{ctx.Value("context")}, args...)...)
}

func (l *LoggerT) PanicfContext(ctx context.Context, format string, args ...interface{}) {
	fmtStr := "[context=%s] " + format
	l.sugaredLogger.Panicf(fmtStr, append([]interface{}{ctx.Value("context")}, args...)...)
}

func (l *LoggerT) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *LoggerT) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *LoggerT) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *LoggerT) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *LoggerT) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

var Logger ILogger = nil
