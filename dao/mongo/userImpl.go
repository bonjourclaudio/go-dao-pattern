package mongo

import (
	"context"
	"github.com/claudioontheweb/go-dao-pattern/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type UserImplMongo struct {}

// Create User
func (dao UserImplMongo) Create(u *models.User) error {

	db := getConnection()
	defer db.Client().Disconnect(context.TODO())

	users, err := dao.GetAll()
	if err != nil {
		log.Fatal(err)
		return err
	}

	highestId := uint(0)

	for _,v := range users {

		if highestId < v.ID {
			highestId = v.ID
		}

	}

	u.ID = highestId + 1


	_, err = db.Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Get User By Id
func (dao UserImplMongo) GetById(id int) (models.User, error) {

	db := getConnection()
	defer db.Client().Disconnect(context.TODO())

	var user models.User

	filter := bson.M{"id": id}

	err := db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, nil
}

// Get All Users
func (dao UserImplMongo) GetAll() ([]models.User, error) {
	var users []models.User

	db := getConnection()
	defer db.Client().Disconnect(context.TODO())

	filter := options.Find()

	cur, err := db.Collection("users").Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error Cursor: ", err)
		return users, err
	}

	for cur.Next(context.TODO()) {

		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal("Error decoding", err)

			return users, err
		}

		users = append(users, user)

	}

	if err := cur.Err(); err != nil {
		log.Fatal("Error Cursor end", err)
		return users, err
	}

	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("Error closing cursor", err)
		return users, err
	}

	return users, nil
}

// Update User
func (dao UserImplMongo) Update(u *models.User) error {
	return nil
}

// Delete User
func (dao UserImplMongo) Delete(id int) error {
	db := getConnection()
	defer db.Client().Disconnect(context.TODO())

	filter := bson.M{"id": id}

	_, err := db.Collection("users").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}