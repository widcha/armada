package configs

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost        string `mapstructure:"db_host"`
	DBMaxConn     int    `mapstructure:"db_max_conn"`
	DBMaxIdleConn int    `mapstructure:"db_max_idle_conn"`
	DBMaxTTLConn  int    `mapstructure:"db_max_ttl_conn"`
	DBName        string `mapstructure:"db_name"`
	DBPassword    string `mapstructure:"db_password"`
	DBPort        string `mapstructure:"db_port"`
	DBUsername    string `mapstructure:"db_username"`
	Port          string `mapstructure:"port"`
	AppName       string `mapstructure:"app_name"`
	AppVersion    string `mapstructure:"app_version"`
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	if cfg == nil {
		cfg = &Config{}
	}

	return cfg
}

func Load() {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()

		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName(".env")

		err := v.ReadInConfig()
		if err != nil {
			fmt.Println("config file not found: ", err)
		}

		config := new(Config)
		err = v.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		cfg = config
	})
}
