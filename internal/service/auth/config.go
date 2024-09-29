package auth

import (
	model "github.com/Egor123qwe/logs-viewer/internal/model/auth"
	"github.com/spf13/viper"
)

func newConfig() model.Credentials {
	return model.Credentials{
		Username: viper.GetString("auth.username"),
		Password: viper.GetString("auth.password"),
	}
}
