package bootstrap

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Env struct {
	AppEnv           string `mapstructure:"APP_ENV"`
	ContextPath      string `mapstructure:"CONTEXT_PATH"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout   int    `mapstructure:"CONTEXT_TIMEOUT"`
	GinMode          string `mapstructure:"GIN_MODE"`
	ServiceServerUrl string `mapstructure:"SERVICE_SERVER_URL"`
	Properties       map[string]interface{}
}

type property struct {
	Id          uint   `json:"id"`
	Application string `json:"application"`
	Profile     string `json:"profile"`
	Key         string `json:"key"`
	Value       string `json:"value"`
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

	resp, err := http.Get("http://localhost:8100/go-centralize-configuration/deposit/default")
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}

	var properties []property
	json.Unmarshal(body, &properties)
	propertyMap := map[string]any{}
	for _, property := range properties {
		key := property.Key
		value := property.Value
		propertyMap[key] = value
		viper.Set(key, value)
	}
	env.Properties = propertyMap

	return &env
}
