package loginst

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Level      string   `json:"level"`
	Encoding   string   `json:"encoding"`
	OutputPath []string `json:"output_path"`
}

var log *zap.Logger

func Inst() *zap.Logger {
	return log
}

func init() {

	conf := zap.NewProductionConfig()
	conf.OutputPaths = []string{"stdout"}
	conf.EncoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
	})

	var err error
	log, err = conf.Build()
	if err != nil {
		panic(err)
	}
}
