package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

//DB Env config
type envConfig struct {
	Version    string `env:"VERSION" envDefault:"0.0.1"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`

	// DBName  string `env:"DB_Name,required"`
	// HOST string `env:"URL_HOST,required"`
	// PORT string `env:"PORT,required"`
	// UrlInfoCollectionName string `env:"Url_Info_Collection_Name,required"`
}

var (
	// Env is the config
	Env = envConfig{}
)

// Setup setup config function
func Setup() {
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Printf("%+v\n", Env)
}
