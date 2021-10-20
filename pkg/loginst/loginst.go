package loginst

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"gitlab.xinghuolive.com/Backend-Go/octopus/instance/alerterinst"
)

const (
	RequestId = "request_id"
)

type Logger struct {
	*log.Logger
}

var logger *Logger

func Inst() *Logger {
	return logger
}

// LogHook log hook模板
type LogHook func(entry *log.Entry)

func init() {
	logger = &Logger{Logger: log.New()}
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stderr)
	logger.SetLevel(log.InfoLevel)
	AddLogHook(alerterinst.NotifyLogEntry)
}

// AddLogHook add a log reporter.
func AddLogHook(f LogHook) {
	m := NewMonitor(f)
	logger.AddHook(m)
}

/*
	这个方法不要在预发布和线上环境用
*/
func SetDebugLevel() {
	logger.SetLevel(log.DebugLevel)
}

// Monitor 信息监控
type Monitor struct {
	Callback LogHook
}

// NewMonitor 返回一个实例
func NewMonitor(l LogHook) *Monitor {
	m := new(Monitor)
	m.Callback = l
	return m
}

// Levels 这些级别的日志会被回调
func (m *Monitor) Levels() []log.Level {
	return []log.Level{
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
	}
}

// Fire 实际执行了回调
func (m *Monitor) Fire(entry *log.Entry) error {
	m.Callback(entry)
	return nil
}

func (l *Logger) WithRequestId(ctx context.Context) *log.Entry {
	return l.WithField(RequestId, GetRequestId(ctx))
}

func GetRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	requestId, ok := ctx.Value(RequestId).(string)
	if !ok {
		return ""
	}
	return requestId
}
