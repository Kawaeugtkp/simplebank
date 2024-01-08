package util

import "github.com/spf13/viper"

// Config stores all configration of the application.
// The values are read by viper from a config file or environment variables.
// mapstructureでapp.envで作成したものを指定する
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// loadConfig reads configration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	// 以下二行はhaapp.envファイルを使用できるようにしたもの
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
