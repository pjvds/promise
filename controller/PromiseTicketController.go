package controller

import (
	log "code.google.com/p/log4go"
	"github.com/pjvds/promise/data"
	"github.com/pjvds/promise/model"
	"io/ioutil"
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
	repository := ctr.repository
	marshaller := ctr.marshaller

	wire, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error("couldn't read request body: " + err.Error())
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		var promise model.PromiseTicket
		err = marshaller.Unmarshal(wire, &promise)

		if err != nil {
			log.Warn("bad request: ", err)

			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(err.Error()))
		} else {
			repository.Add(&promise)
		}
	}
}
