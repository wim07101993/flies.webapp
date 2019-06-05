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
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jv)
}
