package participants

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Service struct {
	filePath string
}

func NewService(filePath string) Service {
	return Service{
		filePath: filePath,
	}
}

func (pc *Service) readFile() ([]Participant, error) {
	participants := []Participant{}
	if _, err := os.Stat(pc.filePath); os.IsNotExist(err) {
		return participants, pc.writeFile(participants)
	}

	bParticipants, err := ioutil.ReadFile(pc.filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bParticipants, &participants)
	return participants, err
}

func (pc *Service) writeFile(participants []Participant) error {
	jParticipants, err := json.Marshal(participants)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pc.filePath, jParticipants, 664)
	return err
}
