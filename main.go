package main

import (
	//"mongorestaurantsample-master/domain"
	"fmt"
	"os"
	"strings"

	//	"bson"
	//	dbrepo "github.com/priteshgudge/mongosample/dbrepository"
	dbrepo "./dbrepository"

	//	mongoutils "github.com/priteshgudge/mongosample/utils"
	mongoutils "./utils"
)

//	importing domain obbject restaurant
//import "mongorestaurantsample-master/domain"
//import "gopkg.in/mgo.v2/bson"

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	//	fmt.Println(repoaccess)

	//Run sample commands

	document := os.Args[1]
	data := os.Args[2]
	s := strings.Split(data, "=")
	s[0] = s[0][2:]
	var count int
	if document == "find" {
		if s[0] == "type_of_food" {
			result, _ := repoaccess.FindByTypeOfFood(s[1])
			fmt.Println(result)
		} else if s[0] == "postcode" {
			result, _ := repoaccess.FindByTypeOfPostCode(s[1])
			fmt.Println(result)
		} else {
			fmt.Println("Invalid arguments")
		}
	} else if document == "count" {
		if s[0] == "type_of_food" {
			count, _ = repoaccess.CountRestaurantByFood(s[1])
			fmt.Println("count Restaurant ByFood ", count)
		} else {
			fmt.Println("Invalid arguments")
		}
	} else {
		fmt.Println("Invalid arguments")
	}
}
