package mongo

import (
	"context"
	"github.com/claudioontheweb/go-dao-pattern/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func getConnection() *mongo.Database {

	config.GetConfig()

	var DB_NAME = viper.GetString("DB.MYSQL.DB_NAME")
	/*var DB_USER = viper.GetString("DB.MYSQL.DB_USER")
	var DB_PASS = viper.GetString("DB.MYSQL.DB_PASS")*/
	var DB_HOST = viper.GetString("DB.MYSQL.DB_HOST")

	clientOptions := options.Client().ApplyURI("mongodb://" + DB_HOST + ":27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	return client.Database(DB_NAME)
}