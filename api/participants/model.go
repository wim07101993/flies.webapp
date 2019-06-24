package participants

import (
	"math"
)

type Participant struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Score uint16 `json:"score"`
}

func (p *Participant) IncreaseScore(amount uint16) (newScore uint16) {
	if p.Score > math.MaxUint16-amount {
		p.Score = math.MaxInt16
	} else {
		p.Score += amount
	}
	return p.Score
}

func (p *Participant) DecreaseScore(amount uint16) (newScore uint16) {
	if p.Score < amount {
		p.Score = 0
	} else {
		p.Score -= amount
	}
	return p.Score
}

func (p *Participant) SetScore(value uint16) (newScore uint16) {
	p.Score = value
	return p.Score
}

func (p *Participant) GetScore() uint16 {
	return p.Score
}
