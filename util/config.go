package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LineSecret    string `mapstructure:"LINE_SECRET"`
	LineToken     string `mapstructure:"LINE_TOKEN"`
	MongoHost     string `mapstructure:"MONGO_HOST"`
	MongoUsername string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	MongoPassword string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	MongoURI      string
	ServerPort    string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	config.makeMongoURI()
	return
}

func (c *Config) makeMongoURI() {
	c.MongoURI = fmt.Sprintf("mongodb://%s:%s@%s:27017/", c.MongoUsername, c.MongoPassword, c.MongoHost)
}
