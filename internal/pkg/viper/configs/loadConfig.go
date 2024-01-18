package configs

import "github.com/spf13/viper"

type conf struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
}

func LoadConfig() (*conf, error) {

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg := &conf{
		DBDriver: viper.GetString("DB_DRIVER"),
		DBUser:   viper.GetString("DB_USER"),
		DBPass:   viper.GetString("DB_PASS"),
		DBHost:   viper.GetString("DB_HOST"),
		DBPort:   viper.GetString("DB_PORT"),
		DBName:   viper.GetString("DB_NAME"),
	}

	return cfg, err
}
