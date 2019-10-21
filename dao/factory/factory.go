package factory

import (
	"github.com/claudioontheweb/go-dao-pattern/dao/mysql"
	"github.com/claudioontheweb/go-dao-pattern/dao/repository"
	"log"
)

func FactoryDao(e string) repository.UserDao {
	var i repository.UserDao

	switch e {
	case "mysql":
		i = mysql.UserImplMysql{}

	default:
		log.Fatal("No Implementation ")
		return nil
	}

	return i
}