package main

import (
	logger "log"
	"mongorestaurantsample-master/dbrepository"
	handlerlib "mongorestaurantsample-master/restapplication/packages/httphandlers"
	"mongorestaurantsample-master/restapplication/restaurantcrudhandler"
	mongoutils "mongorestaurantsample-master/utils"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	// "github.com/priteshgudge/gohttpexamples/sample4/dbrepo/userrepo"
	// handlerlib "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/httphandlers"
	// "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/usercrudhandler"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {

	pingHandler := &handlerlib.PingHandler{}
	dbname := "restaurant"
	mongosession, err := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	if err != nil {
		logger.Fatal("Error in creating mongoSession")
	}
	dbsession := dbrepository.NewMongoRepository(mongosession, dbname)
	handler := restaurantcrudhandler.NewRestaurantCrudHandler(dbsession)
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/restaurant/ping", pingHandler)
	h.Handle("/restaurant/", handler)
	h.Handle("/restaurant/{id}", handler)
	h.Handle("/restaurant/", handler)
	logger.Fatal(http.ListenAndServe(":8080", h))
}
