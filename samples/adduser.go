package main

import "fmt"
import "github.com/timburks/agents"
import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password string
}

func getMongoSession() (mongoSession *mgo.Session) {
	mongoSession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)
	return mongoSession
}


func main() {
	username := "test"
	password := "123"

	fmt.Printf("Setting password for %v with password %v.\n", username, password)
	
	
	saltedPassword := agents.MD5HashWithSalt(password, "agent.io")
	
	fmt.Printf("salted password is %v", saltedPassword)
	
	mongoSession := getMongoSession()
	defer mongoSession.Close()
	
	usersCollection := mongoSession.DB("accounts").C("users")
	
	
	var user User
	user.Username = username
	user.Password = saltedPassword
	usersCollection.Insert(&user)
	
}
