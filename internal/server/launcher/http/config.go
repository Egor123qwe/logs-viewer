package http

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port         int
	ShutdownTime time.Duration
	ReadTime     time.Duration
}

func NewConfig() Config {
	return Config{
		Port:         viper.GetInt("http.port"),
		ShutdownTime: viper.GetDuration("http.shutdown_time"),
		ReadTime:     viper.GetDuration("http.read_time"),
	}
}
