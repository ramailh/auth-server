package props

import (
	"os"
)

var (
	Port, MongoHost, MongoUsername, MongoPassword, DBName, ClientID, ClientSecret string
)

func Setup() {
	DBName = os.Getenv("DB_NAME")
	Port = os.Getenv("PORT")
	MongoHost = os.Getenv("MONGO_HOST")
	MongoUsername = os.Getenv("MONGO_USERNAME")
	MongoPassword = os.Getenv("MONGO_PASSWORD")
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
}
