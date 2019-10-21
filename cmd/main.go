package main

import (
	"fmt"
	"github.com/claudioontheweb/go-dao-pattern/dao/factory"
	"log"
)

func main() {

	userDao := factory.FactoryDao("mysql")

	users, err := userDao.GetAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(users)

}