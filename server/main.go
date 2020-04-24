package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Email string
}

func main() {

	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Do the following:
	// In a command window:
	// set MONGOLAB_URI=mongodb://heroku_x6vpdmxg:Password@ds135514.mlab.com:35514/heroku_x6vpdmxg
	// heroku_x6vpdmxg is my username, replace the same with yours. Type in your password.
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})

	if xxx, err := sess.DatabaseNames(); err == nil {
		fmt.Println("I AM DATABASE NAMES:", xxx[0])
	}

	collection := sess.DB("heroku_x6vpdmxg").C("users")

	err = collection.Insert(&Person{"Stefan Klaste", "klaste@posteo.de"},
		&Person{"Nishant Modak", "modak.nishant@gmail.com"},
		&Person{"Prathamesh Sonpatki", "csonpatki@gmail.com"},
		&Person{"murtuza kutub", "murtuzafirst@gmail.com"},
		&Person{"aniket joshi", "joshianiket22@gmail.com"},
		&Person{"Michael de Silva", "michael@mwdesilva.com"},
		&Person{"Alejandro Cespedes Vicente", "cesal_vizar@hotmail.com"})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	result := Person{}
	err = collection.Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}

	fmt.Println("Email Id:", result.Email)
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": result.Name,
		})
	})

	r.Run()
}
