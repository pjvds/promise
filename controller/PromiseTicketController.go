package controller

import (
	"github.com/pjvds/promise/data"
	"github.com/pjvds/promise/model"
	"io/ioutil"
	"log"
	"net/http"
)

type PromiseTicketController struct {
	repository data.PromiseTicketRepository
	marshaller data.Marshaller
}

func NewPromiseTicketController(repository data.PromiseTicketRepository, marshaller data.Marshaller) *PromiseTicketController {
	return &PromiseTicketController{
		repository: repository,
		marshaller: marshaller,
	}
}

func (ctr *PromiseTicketController) Handle(response http.ResponseWriter, request *http.Request) {
	log.Println("handling promise response: " + request.RequestURI)

	repository := ctr.repository
	marshaller := ctr.marshaller

	switch request.Method {
	case "GET":
		promises := repository.All()
		wire, err := marshaller.Marshal(promises)
		if err != nil {
			log.Println("error while marshalling promises: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(wire)
		}

	case "POST":
		wire, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println("couldn't read request body: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
		} else {
			var promise model.PromiseTicket
			err = marshaller.Unmarshal(wire, &promise)

			if err != nil {
				log.Printf("bad request: ", err)

				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte(err.Error()))
			} else {
				repository.Add(promise)
			}
		}

	default:
		log.Printf("unsupported http request method: %v", request.Method)
		response.WriteHeader(http.StatusBadRequest)
	}
}
