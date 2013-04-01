package controller

import (
	log "code.google.com/p/log4go"
	//"errors"
	"github.com/pjvds/promise/data"
	"github.com/pjvds/promise/messaging"
	"github.com/pjvds/promise/model"
	"github.com/pjvds/promise/serialization"
	"io/ioutil"
	"net/http"
	"time"
)

type PromiseTicketController struct {
	repository data.PromiseTicketRepository
	marshaller serialization.Marshaller
	bus        messaging.Bus
}

func NewPromiseTicketController(repository data.PromiseTicketRepository, bus messaging.Bus, marshaller serialization.Marshaller) *PromiseTicketController {
	return &PromiseTicketController{
		repository: repository,
		marshaller: marshaller,
		bus:        bus,
	}
}

func (ctr *PromiseTicketController) Handle(response http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		ctr.get(response, request)
	case "POST":
		ctr.post(response, request)

	default:
		log.Error("unsupported http request method: %v", request.Method)
		response.WriteHeader(http.StatusBadRequest)
	}
}

func (ctr *PromiseTicketController) get(response http.ResponseWriter, request *http.Request) {
	repository := ctr.repository
	marshaller := ctr.marshaller

	promises := repository.All()
	wire, err := marshaller.Marshal(promises)
	if err != nil {
		log.Error("error while marshalling promises: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(wire)
	}
}

func (ctr *PromiseTicketController) post(response http.ResponseWriter, request *http.Request) {
	marshaller := ctr.marshaller

	wire, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error("couldn't read request body: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		var doc map[string]interface{}
		err = marshaller.Unmarshal(wire, &doc)
		if err != nil {
			log.Warn("bad request: ", err)

			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(err.Error()))
		} else {

			name := doc["Name"].(string)
			url := doc["Url"].(string)
			when := time.Now() // todo: just stamp it with now due unmarshalling errors :-/

			if err != nil {
				log.Warn("invalid time: %v, %v", doc["When"], err.Error())
				response.WriteHeader(http.StatusBadRequest)
			} else {
				ticket := model.NewTicket(name, url, when)

				log.Info(ticket)

				if err != nil {
					log.Error("unable to publish message with bus: %v", err.Error())
					response.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
	}
}
