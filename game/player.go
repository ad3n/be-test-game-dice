package game

type Player struct {
	TotalDice int
	ToNext    int
	Dices     []*Dice
	Point     int
}

func (p *Player) AddDices(n int) {
	for i := 1; i <= n; i++ {
		p.Dices = append(p.Dices, &Dice{})
	}

	p.TotalDice = len(p.Dices)
}

func (p *Player) ThrowDices() {
	for _, m := range p.Dices {
		m.Throw()
	}
}

func (p *Player) Result() []int {
	result := []int{}
	for _, d := range p.Dices {
		switch d.Result {
		case 1:
			p.TotalDice--
			p.ToNext++
		case 6:
			p.Point++
			p.TotalDice--
		}

		result = append(result, d.Result)
	}

	return result
}

func (p *Player) Statistic() []int {
	result := []int{}
	for _, d := range p.Dices {
		if d.Result != 6 && d.Result != 1 {
			result = append(result, d.Result)
		}
	}

	return result
}
