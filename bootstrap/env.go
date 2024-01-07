package bootstrap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const ServiceServerUrl = "host.go-service-server.url"

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ContextPath    string `mapstructure:"CONTEXT_PATH"`
	Profile        string `mapstructure:"PROFILE"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	GinMode        string `mapstructure:"GIN_MODE"`
	CloudConfigUrl string `mapstructure:"CLOUD_CONFIG_URL"`
	Properties     map[string]interface{}
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

	configUrl := fmt.Sprintf("%s/%s/%s", env.CloudConfigUrl, env.ContextPath, env.Profile)

	resp, err := http.Get(configUrl)
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
