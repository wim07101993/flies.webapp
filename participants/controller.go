package participants

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	BadJsonErrorMessage = "The given object could not be interpreted by the server."
	NameParameter       = "name"
)

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (pc *Controller) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jp, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if checkError(w, err) {
		return
	}

	var p Participant
	err = json.Unmarshal(jp, &p)
	if checkError(w, errors.New(BadJsonErrorMessage)) {
		return
	}

	p, err = pc.service.Create(p)
	if checkError(w, err) {
		return
	}

	writeJson(w, p)
}

func (pc *Controller) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ps, err := pc.service.GetAll()
	if checkError(w, err) {
		return
	}

	writeJson(w, ps)
}

func (pc *Controller) Get(w http.ResponseWriter, r *http.Response, ps httprouter.Params) {
	name := ps.ByName(NameParameter)
	p, err := pc.service.Get(name)
	if checkError(w, err) {
		return
	}

	writeJson(w, p)
}
