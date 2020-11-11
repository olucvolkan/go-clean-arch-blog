package main

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/olucvolkan/go-clean-arch-blog/config"
)


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


func main(){

	config := config.New()

	ensureDBExists(config)

	fmt.Println(config.DBUrl())
	gormDB, err := gorm.Open("mysql", config.DBUrl())
	if err != nil {
		fmt.Println(fmt.Errorf("Can't connect to database, err: %v", err))
		return
	}
	gormDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Todo{})
	defer gormDB.Close()
}