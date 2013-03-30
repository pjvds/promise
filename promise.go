package main

import (
	//"labix.org/v2/mgo"
	"flag"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"time"
)

var (
	promiseRepository = NewPromiseRepository(99)
	address           = *flag.String("address", ":8080", "address to server on, eq: 127.0.0.1:80")
	marshaller        = NewJsonMarshaller()
)

const (
	HTTP_GET  = "GET"
	HTTP_POST = "POST"
)

func main() {
	flag.Parse()

	promiseRepository.Add(PromiseTicket{
		Id:           bson.NewObjectId(),
		Name:         "ticket 1",
		ExecuteAfter: time.Now(),
	})

	promiseRepository.Add(PromiseTicket{
		Id:           bson.NewObjectId(),
		Name:         "ticket 2",
		ExecuteAfter: time.Now(),
	})

	http.Handle("/promise", http.HandlerFunc(handlePromise))

	log.Println("Serving at " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handlePromise(response http.ResponseWriter, request *http.Request) {
	log.Println("handling promise response: " + request.RequestURI)

	switch request.Method {
	case HTTP_GET:
		promises := promiseRepository.All()
		wire, err := marshaller.Marshal(promises)
		if err != nil {
			log.Println("error while marshalling promises: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(wire)
		}

	case HTTP_POST:
		wire, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("couldn't read request body: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			var promise PromiseTicket
			err = marshaller.Unmarshal(wire, &promise)

			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte(err.Error()))
			} else {
				promiseRepository.Add(promise)
			}
		}

	default:
		log.Printf("unsupported http request method: %v", request.Method)
		response.WriteHeader(http.StatusBadRequest)
	}
}
