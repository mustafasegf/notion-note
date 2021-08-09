package util

import "github.com/spf13/viper"

type Config struct {
	NotionToken      string `mapstructure:"NOTION_TOKEN"`
	NotionDatabaseID string `mapstructure:"NOTION_DATABASE_ID"`
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

	return
}
