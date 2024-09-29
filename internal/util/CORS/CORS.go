package CORS

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
)

func New(router http.Handler) http.Handler {
	cors := handlers.CORS(
		handlers.AllowedOrigins(viper.GetStringSlice("allowed_origins")),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"application/json", "Content-Type", "text/plain"}),
		handlers.AllowCredentials(),
	)

	return cors(router)
}
