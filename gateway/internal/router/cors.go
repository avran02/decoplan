package router

import (
	"github.com/avran02/decoplan/gateway/internal/config"
	"github.com/go-chi/cors"
)

func getCorsOptions(corsConf config.CORS) cors.Options {
	return cors.Options{
		AllowedOrigins:   corsConf.AllowedOrigins,
		AllowedMethods:   corsConf.AllowedMethods,
		AllowedHeaders:   corsConf.AllowedHeaders,
		AllowCredentials: corsConf.AllowCredentials,
		Debug:            true,
	}
}
