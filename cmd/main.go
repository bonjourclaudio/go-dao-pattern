package main

import (
	"fmt"
	"github.com/claudioontheweb/go-dao-pattern/dao/factory"
	"github.com/claudioontheweb/go-dao-pattern/models"
	"log"
)

func main() {

	// Get User DAO for MYSQL
	userDao := factory.FactoryDao("mongo")

	// Create User
	var newUser models.User
	newUser.Firstname = "Peter"
	newUser.Lastname = "Witzig"

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
	user, err := userDao.GetById(9)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user)

	/*
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
	*/

	// Delete User
	err = userDao.Delete(20)
	if err != nil {
		log.Fatal(err)
		return
	}

}