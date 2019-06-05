package participants

import (
	"encoding/json"
	"net/http"
)

func findParticipant(name string, participants []Participant) int {
	for i, participant := range participants {
		if participant.Name == name {
			return i
		}
	}
	return -1
}

func writeJson(w http.ResponseWriter, v interface{}) {
	jv, err := json.Marshal(v)
	if checkError(w, err) {
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jv)
}

func checkError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	errMes := err.Error()
	if errMes == ParticipantNotFoundErrorMessage {
		http.Error(w, errMes, http.StatusNotFound)
	} else if errMes == NameAlreadyTakenErrorMessage ||
		errMes == BadJsonErrorMessage {
		http.Error(w, errMes, http.StatusBadRequest)
	} else {
		http.Error(w, errMes, http.StatusInternalServerError)
	}

	return true
}
