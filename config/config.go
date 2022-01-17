package config

import (
	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
)

type Config struct {
	AppName  string
	IsDebug  bool
	Server   ServerConfig
	Database Database
}

func New() (*Config, error) {
	config := &Config{}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("gora")

	const defaultPort = 8080

	viper.SetDefault("APPNAME", "gora")
	viper.SetDefault("ISDEBUG", true)
	viper.SetDefault("PORT", defaultPort)
	viper.SetDefault("DB_HOSTNAME", "localhost")
	viper.SetDefault("DB_USERNAME", "fake")
	viper.SetDefault("DB_PASSWORD", "fakepassword")
	viper.SetDefault("DB_PORT", 5432)

	config.AppName = viper.GetString("APPNAME")
	config.IsDebug = viper.GetBool("ISDEBUG")
	config.Server = ServerConfig{Port: viper.GetInt("PORT")}
	config.Database = Database{
		Hostname: viper.GetString("DB_HOSTNAME"),
		Username: viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		Port:     viper.GetInt("DB_PORT"),
	}

	return config, nil
}

func (c *Config) Print() {
	_, _ = pp.Println(c)
}
