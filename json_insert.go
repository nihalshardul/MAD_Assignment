package main

import (
	//"mongorestaurantsample-master/domain"
	"fmt"
	"os"
	"bufio"
	"encoding/json"
//	"bson"
//	dbrepo "github.com/priteshgudge/mongosample/dbrepository"
	dbrepo "./dbrepository"

//	mongoutils "github.com/priteshgudge/mongosample/utils"
	mongoutils "./utils"
)
//	importing domain obbject restaurant
import "mongorestaurantsample-master/domain"
//import "gopkg.in/mgo.v2/bson"

func main() {
	//pass mongohost through the environment
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))

	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
//	fmt.Println(repoaccess)

	//Run sample commands

	data, err := os.Open("./restaurant.json")

	if err!=nil{
		fmt.Println("File reading error",err)
		return
	}//else{
//		fmt.Println("Content in file : ",string(data))
//	}

	scanner := bufio.NewScanner(data)

	var res domain.Restaurant

//	fmt.Println(res)
	str:=""

	for scanner.Scan(){
		str=scanner.Text()
		json.Unmarshal([]byte(str), &res)
	//	fmt.Println(res)
		res.DBID = domain.NewID()
	//	fmt.Println(res)
		repoaccess.Store(&res)
	}

}
