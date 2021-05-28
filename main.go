package main

import (
	"fmt"
	"os"

	"github.com/ad3n/dice-game/game"
)

func main() {
	nPlayer := 0
	nDice := 0

	fmt.Print("Masukkan Jumlah Pemain: ")
	_, err := fmt.Scanln(&nPlayer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Print("Masukkan Jumlah Dadu: ")
	_, err = fmt.Scanln(&nDice)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	game := game.Game{}
	game.Play(nPlayer, nDice)
}
