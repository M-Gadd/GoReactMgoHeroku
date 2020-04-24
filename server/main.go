package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Email string
}

func main() {

	r := gin.Default()

	// Do the following:
	// In a command window:
	// set MONGOLAB_URL=mongodb://IndianGuru:dbpassword@ds051523.mongolab.com:51523/godata
	// IndianGuru is my username, replace the same with yours. Type in your password.
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

	collection := sess.DB("godata").C("user")

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
