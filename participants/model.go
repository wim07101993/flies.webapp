package participants

import (
	"math"
)

type Participant struct {
	Name  string
	score uint16
}

func (p *Participant) IncreaseScore(amount uint16) (newScore uint16) {
	if p.score > math.MaxUint16-amount {
		p.score += amount
	} else {
		p.score = math.MaxInt16
	}
	return p.score
}

func (p *Participant) DecreaseScore(amount uint16) (newScore uint16) {
	if p.score < amount {
		p.score = 0
	} else {
		p.score -= amount
	}
	return p.score
}

func (p *Participant) SetScore(value uint16) (newScore uint16) {
	p.score = value
	return p.score
}

func (p *Participant) GetScore() uint16 {
	return p.score
}
