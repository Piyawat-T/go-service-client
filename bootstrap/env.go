package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv           string `mapstructure:"APP_ENV"`
	ContextPath      string `mapstructure:"CONTEXT_PATH"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout   int    `mapstructure:"CONTEXT_TIMEOUT"`
	GinMode          string `mapstructure:"GIN_MODE"`
	ServiceServerUrl string `mapstructure:"SERVICE_SERVER_URL"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
