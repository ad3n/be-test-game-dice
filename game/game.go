package game

import (
	"fmt"
	"strings"
)

type Game struct {
	Players []*Player
	Point   map[int]int
	nThrown int
	finish  bool
}

func (g *Game) Play(nPlayers int, nDices int) {
	g.finish = false
	g.nThrown = 0
	g.Point = map[int]int{}
	for i := 1; i <= nPlayers; i++ {
		g.Players = append(g.Players, &Player{})
	}

	for _, p := range g.Players {
		p.AddDices(nDices)
	}

	g.doPlay()
}

func (g *Game) result() {
	fmt.Println("=========================================================")

	for i, p := range g.Players {
		fmt.Printf("Pemain #%d (%d): ", (i + 1), g.Point[i])
		numbers := p.Result()
		result := ""
		for _, v := range numbers {
			result = fmt.Sprintf("%s%d,", result, v)
		}

		fmt.Printf("%s\n", strings.Trim(result, ","))
	}

	fmt.Println("=========================================================")
}

func (g *Game) statistic() {
	fmt.Println("=========================================================")

	addjustment := map[int]int{}
	for i, p := range g.Players {
		g.Point[i] = g.Point[i] + p.Point

		idx := i + 1
		if i == len(g.Players)-1 {
			idx = 0
		}

		addjustment[idx] = g.Players[i].ToNext
	}

	for i, p := range g.Players {
		numbers := p.Statistic()
		fmt.Printf("Pemain #%d (%d): ", (i + 1), g.Point[i])
		result := ""
		for _, v := range numbers {
			result = fmt.Sprintf("%s%d,", result, v)
		}

		for j := 0; j < addjustment[i]; j++ {
			result = fmt.Sprintf("%s1,", result)
		}

		fmt.Printf("%s\n", strings.Trim(result, ","))
	}

	fmt.Println("=========================================================")
}

func (g *Game) doPlay() {
	if !g.finish {
		g.throw()
		g.statistic()

		g.reset()

		emptyDicePlayer := 0
		for _, p := range g.Players {
			if p.TotalDice == 0 {
				emptyDicePlayer++
			}
		}

		if emptyDicePlayer == len(g.Players)-1 {
			g.finish = true

			winner := Winner{}
			for i, p := range g.Point {
				if i == 0 {
					winner.Index = i
					winner.Point = p

					continue
				}

				if winner.Point >= p {
					continue
				}

				winner.Index = i
				winner.Point = p
			}

			if winner.Point == 0 {
				fmt.Println("Permainan Seri")
			} else {
				fmt.Printf("Pemenang: Pemain #%d (%d)\n", (winner.Index + 1), winner.Point)
			}

		}

		g.doPlay()
	}
}

func (g *Game) throw() {
	g.nThrown++

	fmt.Printf("Lempar #%d\n", g.nThrown)
	for _, p := range g.Players {
		p.ThrowDices()
	}

	g.result()

	fmt.Printf("Hasil #%d\n", g.nThrown)
}

func (g *Game) reset() {
	length := len(g.Players)
	temp := []*Player{}

	for i := 1; i <= length; i++ {
		temp = append(temp, &Player{})
	}

	addjustment := map[int]int{}
	for i := range temp {
		idx := i + 1
		if i == len(temp)-1 {
			idx = 0
		}

		addjustment[idx] = g.Players[i].ToNext
	}

	for i, p := range temp {
		p.AddDices(g.Players[i].TotalDice + addjustment[i])
	}

	g.Players = temp
}
