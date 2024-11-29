package init

import (
	"movie-festival-app/entity"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupMainConfig
func SetupMainConfig() (config *entity.Config) {
	viper.SetConfigFile("main.json")
	zap.S().Info("Environment: ", "Local")

	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Fatal(err)
	}
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err = viper.Unmarshal(&config)
	if err != nil {
		zap.S().Fatal(err)
	}
	return
}

// Setup Logger
func SetupLogger() *zap.Logger {
	_, err := os.Stat("./log")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./log", 0755)
		if errDir != nil {
			zap.S().Fatal(err)
		}
	}

	_, err = os.OpenFile("./log/go.log", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		zap.S().Fatal(err)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000000Z")
	config.OutputPaths = []string{"stdout", "./log/go.log"}
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	return logger
}
