package database

import (
"fmt"
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
"github.com/joho/godotenv"
	"log"
	"os"
)

var MYSQLDBGORM *gorm.DB

func ConnectMySQL() {

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return
	}

	databaseUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_host"), os.Getenv("db_port"), os.Getenv("db_name"))
	MYSQLDBGORM, err = gorm.Open("mysql", databaseUrl)

	if err != nil {
		log.Println(err)
		return
	}

	err = MYSQLDBGORM.DB().Ping()

	if err != nil {
		log.Println(err)
		return
	}

	MYSQLDBGORM.LogMode(true)
	log.Println("Successful connection to MySQL DB ")

}

func CloseMySQL() error {
	return MYSQLDBGORM.Close()
}
