package main

import (
"fmt"
"github.com/claudioontheweb/go-dao-pattern/dao/factory"
	"github.com/claudioontheweb/go-dao-pattern/models"
	"log"
)

func main() {

	// Get User DAO for MYSQL
	userDao := factory.FactoryDao("mysql")

	// Create User
	var newUser models.User
	newUser.Firstname = "Hans"
	newUser.Lastname = "Lustig"

	err := userDao.Create(&newUser)
	 if err != nil {
	 	log.Fatal(err)
	 	return
	 }

	// Get All Users
	users, err := userDao.GetAll()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(users)

	// Get User By ID
	user, err := userDao.GetById(10)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user)

	// Update User
	var updatedUser models.User
	updatedUser.ID = 9
	updatedUser.Firstname = "Hansss"
	updatedUser.Lastname = "Lustigggg"

	err = userDao.Update(&updatedUser)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Delete User
	err = userDao.Delete(6)
	if err != nil {
		log.Fatal(err)
		return
	}

}