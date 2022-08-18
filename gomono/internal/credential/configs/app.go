package configs

import (
	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	GraphQL struct {
		Host       string `env:"APP_GQL_HOST,required,notEmpty"`
		AuthHeader string `env:"APP_GQL_AUTH_HEADER,required,notEmpty" envDefault:"X-Hasura-Admin-Secret"`
		AuthSecret string `env:"APP_GQL_AUTH_SECRET,required,notEmpty"`
	}
	JWT struct {
		Secret string `env:"APP_JWT_SECRET,required,notEmpty"`
	}
}

var App *AppConfig

func Init() {
	App = &AppConfig{}
	env.Parse(App)
}
