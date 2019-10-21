package mysql

import (
	"database/sql"
	"github.com/claudioontheweb/go-dao-pattern/config"
	"github.com/spf13/viper"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {

	config.GetConfig()

	var DB_NAME = viper.GetString("DB.MYSQL.DB_NAME")
	var DB_USER = viper.GetString("DB.MYSQL.DB_USER")
	var DB_PASS = viper.GetString("DB.MYSQL.DB_PASS")
	var DB_HOST = viper.GetString("DB.MYSQL.DB_HOST")


	db, err := sql.Open("mysql", DB_USER + ":" + DB_PASS + "@tcp("+ DB_HOST + ":3306)/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db

}