package participants

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const (
	ParticipantNotFoundErrorMessage = "Participant not found"
	NameAlreadyTakenErrorMessage    = "Name already taken"
)

type Service struct {
	filePath string
}

func NewService(filePath string) Service {
	return Service{
		filePath: filePath,
	}
}

func (pc *Service) Create(participant Participant) (Participant, error) {
	participants, err := pc.readFile()
	if err != nil {
		return Participant{}, err
	}
	if findParticipant(participant.Name, participants) >= 0 {
		return Participant{}, errors.New(NameAlreadyTakenErrorMessage)
	}

	participants = append(participants, participant)

	if err = pc.writeFile(participants); err != nil {
		return Participant{}, err
	} else {
		return participant, nil
	}
}

func (pc *Service) GetAll() ([]Participant, error) {
	return pc.readFile()
}

func (pc *Service) Get(name string) (Participant, error) {
	ps, err := pc.readFile()
	if err != nil {
		return Participant{}, err
	}

	i := findParticipant(name, ps)
	if i < 0 {
		return Participant{}, errors.New(ParticipantNotFoundErrorMessage)
	}

	return ps[i], nil
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
