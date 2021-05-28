package game

import "math/rand"

type Dice struct {
	Result int
}

func (d *Dice) Throw() {
	d.Result = rand.Intn(5) + 1
}
