package global

import (
	"github.com/caarlos0/env/v6"
)

type GlobalConfig struct {
	Name string `env:"NAME,required,notEmpty"`
	Port int    `env:"PORT" envDefault:"3000"`
}

var Config *GlobalConfig

func Init() {
	Config = &GlobalConfig{}
	env.Parse(Config)
}
