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

const (
	ServiceServerUrl = "host.go_service_server.url"
	GinMode          = "gin_mode"
	Profile          = "profile"
	ServerAddress    = "server_address"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	Profile        string `mapstructure:"PROFILE"`
	ContextPath    string `mapstructure:"CONTEXT_PATH"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
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
