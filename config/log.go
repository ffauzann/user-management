package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logPath = "./log/go.log"
)

// SetupLogger used to configure CLI log
func SetupLogger() {
	_, err := os.Stat("log")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("log", 0755)
		if errDir != nil {
			zap.S().Fatal(err)
		}

	}

	_, err = os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		zap.S().Error(err)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.OutputPaths = []string{"stdout", logPath}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}
