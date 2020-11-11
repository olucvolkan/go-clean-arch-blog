package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/olucvolkan/go-clean-arch-blog/config"
	"log"
)


func main(){

	c := config.New()

	ensureDBExists(c)


	fmt.Println(c.DBUrl())
	dbConn, err := sql.Open("mysql", c.DBUrl())

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	log.Fatal(e.Start(c.HTTPort))
}

func ensureDBExists(config * config.Config) {
	db, err := sql.Open("mysql", config.DBUrlWithoutDBName() )

	if err != nil {
		fmt.Println("can't connect database for creating table")
		return
	}

	defer db.Close()
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBName + ";")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database or updated")

	}
}