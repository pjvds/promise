package main

import (
	log "code.google.com/p/log4go"
	"github.com/pjvds/promise/controller"
	"github.com/pjvds/promise/data"
	"github.com/pjvds/promise/data/mongo"
	"labix.org/v2/mgo"
	"net/http"
	"time"
)

func main() {
	server, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	database := server.DB("promise")
	session := mongo.NewMongoPromiseSession(server, database)
	repoFac := mongo.NewMongoPromiseRepositoryFactory(session)
	repo := repoFac.CreatePromiseTicketRepository()
	marsh := data.NewJsonMarshaller()
	ctrl := controller.NewPromiseTicketController(repo, marsh)

	Serve(&ServeInfo{
		uri: ":8080",
		promiseTicketController: ctrl,
	})
}

type ServeInfo struct {
	uri                     string
	promiseTicketController *controller.PromiseTicketController
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func Serve(info *ServeInfo) {
	ctrl := info.promiseTicketController

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/v1/promise", func(w http.ResponseWriter, r *http.Request) {
		ctrl.Handle(w, r)
	})

	log.Info("Serving at " + info.uri)
	log.Critical(http.ListenAndServe(info.uri, Log(http.DefaultServeMux)))

	// Let log4go flush
	time.Sleep(time.Second)
}
