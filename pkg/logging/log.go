package logging

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Logger interface {
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

type GormLogger struct {
	logger        zap.Logger
	slowThreshold time.Duration
}

func NewGormLogger(logger zap.Logger, slowThreshold time.Duration) *GormLogger {
	return &GormLogger{
		logger:        logger,
		slowThreshold: slowThreshold,
	}
}

// 暂时先不提供更改日志等级的改动
func (ml *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return ml
}

func (ml *GormLogger) Info(ctx context.Context, format string, v ...interface{}) {
	ml.logger.Info(format)
}

func (ml *GormLogger) Warn(ctx context.Context, format string, v ...interface{}) {
	ml.logger.Warn(format)
}

func (ml *GormLogger) Error(ctx context.Context, format string, v ...interface{}) {
	ml.logger.Error(format)
}

const (
	Sql    = "sql"
	CostMS = "cost_ms"
	Row    = "row"
)

func (ml *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil:
		ml.logger.Info("SQL Print ",
			zap.Float64(CostMS, float64(elapsed.Nanoseconds())/1e6),
			zap.String(Sql, sql),
			zap.Int64(Row, rows))
	case ml.slowThreshold != 0 && elapsed > ml.slowThreshold:
		ml.logger.Info("Slow SQL",
			zap.Float64(CostMS, float64(elapsed.Nanoseconds())/1e6),
			zap.String(Sql, sql),
			zap.Int64(Row, rows))
	default:
		ml.logger.Info("Error SQL Print ",
			zap.Float64(CostMS, float64(elapsed.Nanoseconds())/1e6),
			zap.String(Sql, sql),
			zap.Int64(Row, rows))
	}
}
