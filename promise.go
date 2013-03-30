package main

import (
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	//"flag"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

var promises = make([]PromiseTicket, 0, 100)

const (
	HTTP_GET  = "GET"
	HTTP_POST = "POST"
)

type PromiseTicket struct {
	Id           bson.ObjectId
	Name         string
	ExecuteAfter time.Time

	//callback HttpCallback
}

type HttpCallback struct {
	id  bson.ObjectId
	url string
}

func main() {
	log.Println("started at " + time.Now().String())

	promises = append(promises, PromiseTicket{
		Id:           bson.NewObjectId(),
		Name:         "ticket 1",
		ExecuteAfter: time.Now(),
	})

	promises = append(promises, PromiseTicket{
		Id:           bson.NewObjectId(),
		Name:         "ticket 2",
		ExecuteAfter: time.Now(),
	})

	log.Println("start serving...")

	http.Handle("/promise", http.HandlerFunc(handlePromise))
	http.ListenAndServe(":8080", nil)
}

func handlePromise(response http.ResponseWriter, request *http.Request) {
	log.Println("handling promise response: " + request.RequestURI)

	switch request.Method {
	case HTTP_GET:
		wire, err := json.Marshal(promises)
		if err != nil {
			log.Println("error while marshalling promises: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			log.Println("returning: " + string(wire))
			response.Write(wire)
		}

	case HTTP_POST:
		wire, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("couldn't read request body: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			var promise PromiseTicket
			err = json.Unmarshal(wire, &promise)

			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte(err.Error()))
			} else {
				promises = append(promises, promise)
			}
		}

	default:
		log.Printf("unsupported http request method: %v", request.Method)
		response.WriteHeader(http.StatusBadRequest)
	}

}
