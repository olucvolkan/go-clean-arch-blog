package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HTTPort    string
	DBHost     string
	DBUser     string
	DBName     string
	DBPassword string
	DBPort     string
}



func New() *Config {
	return &Config{
		HTTPort:    getConfigFileValues("server.address", "9090"),
		DBHost:     getConfigFileValues("database.host", "localhost"),
		DBUser:     getConfigFileValues("database.user", "root"),
		DBName:     getConfigFileValues("database.name", ""),
		DBPassword: getConfigFileValues("database.pass", ""),
	}
}

// DBUrl returns connection string for DB connection
func (c *Config) DBUrl() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func getConfigFileValues(key, defaultVal string) string {
	val := viper.GetString(key)
	if val == "" {
		return defaultVal
	}

	return val
}
