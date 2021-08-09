package util

import "github.com/spf13/viper"

type Config struct {
	NotionToken      string `mapstructure:"NOTION_TOKEN"`
	NotionDatabaseID string `mapstructure:"NOTION_DATABASE_ID"`
	LineSecret       string `mapstructure:"LINE_SECRET"`
	LineToken        string `mapstructure:"LINE_TOKEN"`
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
