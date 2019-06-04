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

func (pc *Service) Create(p Participant) (Participant, error) {
	ps, err := pc.readFile()
	if err != nil {
		return Participant{}, err
	}
	if findParticipant(p.Name, ps) >= 0 {
		return Participant{}, errors.New(NameAlreadyTakenErrorMessage)
	}

	ps = append(ps, p)

	if err = pc.writeFile(ps); err != nil {
		return Participant{}, err
	} else {
		return p, nil
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

func (pc *Service) IncreaseScore(name string, amount uint16) (Participant, error) {
	return pc.updateParticipant(name, func(p *Participant) error {
		p.IncreaseScore(amount)
		return nil
	})
}

func (pc *Service) DecreaseScore(name string, amount uint16) (Participant, error) {
	return pc.updateParticipant(name, func(p *Participant) error {
		p.DecreaseScore(amount)
		return nil
	})
}

func (pc *Service) UpdateScore(name string, newScore uint16) (Participant, error) {
	return pc.updateParticipant(name, func(p *Participant) error {
		p.SetScore(newScore)
		return nil
	})
}

func (pc *Service) UpdateName(oldName string, newName string) (Participant, error) {
	return pc.updateParticipant(oldName, func(p *Participant) error {
		p.Name = newName
		return nil
	})
}

func (pc *Service) updateParticipant(name string, updater func(*Participant) error) (Participant, error) {
	ps, err := pc.readFile()
	if err != nil {
		return Participant{}, err
	}

	i := findParticipant(name, ps)
	if i < 0 {
		return Participant{}, errors.New(ParticipantNotFoundErrorMessage)
	}

	if err = updater(&ps[i]); err != nil {
		return Participant{}, err
	}

	if err = pc.writeFile(ps); err != nil {
		return Participant{}, err
	}
	return ps[i], nil
}

func (pc *Service) readFile() ([]Participant, error) {
	ps := []Participant{}
	if _, err := os.Stat(pc.filePath); os.IsNotExist(err) {
		return ps, pc.writeFile(ps)
	}

	jps, err := ioutil.ReadFile(pc.filePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jps, &ps)
	return ps, err
}

func (pc *Service) writeFile(ps []Participant) error {
	jps, err := json.Marshal(ps)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pc.filePath, jps, 664)
	return err
}
