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

func writeJsonParticipant(w http.ResponseWriter, p Participant) {
	jp, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(jp)
}
