package api

import "github.com/spf13/viper"

type config struct {
	logs logsConfig
}

type logsConfig struct {
	host string
	port int
}

func newConfig() config {
	return config{
		logs: logsConfig{
			host: viper.GetString("api.logs_storage.host"),
			port: viper.GetInt("api.logs_storage.port"),
		},
	}
}
