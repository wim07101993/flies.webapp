package participants

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	BadJsonErrorMessage   = "The given object could not be interpreted by the server."
	BadAmountErrorMessage = "The given amount could not be interpreted by the server."
	BadScoreErrorMessage  = "The given score could not be interpreted by the server."
	NameParameter         = "name"
	AmountParameter       = "amount"
	ScoreParameter        = "score"
	NewNameParamter       = "newName"
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
	if err != nil {
		checkError(w, errors.New(BadJsonErrorMessage))
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

func (pc *Controller) IncreaseScore(w http.ResponseWriter, r *http.Response, ps httprouter.Params) {
	name := ps.ByName(NameParameter)
	sAmount := ps.ByName(AmountParameter)
	amount, err := strconv.ParseUint(sAmount, 10, 16)
	if err != nil {
		checkError(w, errors.New(BadAmountErrorMessage))
		return
	}

	p, err := pc.service.IncreaseScore(name, uint16(amount))
	if checkError(w, err) {
		return
	}

	writeJson(w, p)
}

func (pc *Controller) DecreaseScore(w http.ResponseWriter, r *http.Response, ps httprouter.Params) {
	name := ps.ByName(NameParameter)
	sAmount := ps.ByName(AmountParameter)
	amount, err := strconv.ParseUint(sAmount, 10, 16)
	if err != nil {
		checkError(w, errors.New(BadAmountErrorMessage))
		return
	}

	p, err := pc.service.DecreaseScore(name, uint16(amount))
	if checkError(w, err) {
		return
	}

	writeJson(w, p)
}

func (pc *Controller) UpdateScore(w http.ResponseWriter, r *http.Response, ps httprouter.Params) {
	name := ps.ByName(NameParameter)
	sScore := ps.ByName(ScoreParameter)
	score, err := strconv.ParseUint(sScore, 10, 16)
	if err != nil {
		checkError(w, errors.New(BadScoreErrorMessage))
		return
	}

	p, err := pc.service.UpdateScore(name, uint16(score))
	if checkError(w, err) {
		return
	}

	writeJson(w, p)
}

func (pc *Controller) UpdateName(w http.ResponseWriter, r *http.Response, ps httprouter.Params) {
	name := ps.ByName(NameParameter)
	newName := ps.ByName(NewNameParamter)

	p, err := pc.service.UpdateName(name, newName)
	if checkError(w, err) {
		return
	}

	writeJson(w, p)
}
