package day_2

import (
	"fmt"
	"log"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data")
	fmt.Println(data[1])
	fmt.Println(strings.Split(data[1], " ")[0])
	score := 0
	win := 6
	draw := 3
	lose := 0
	rock := 1
	papper := 2
	scisors := 3
	for _, v := range data {
		move := strings.Split(v, " ")
		enemy := move[0]
		player := move[1]
		switch player {
		case "X":
			score += lose
			switch enemy {
			case "A":
				score += scisors
			case "B":
				score += rock
			case "C":
				score += papper
			}
		case "Y":
			score += draw
			switch enemy {
			case "A":
				score += rock
			case "B":
				score += papper
			case "C":
				score += scisors
			}
		case "Z":
			score += win
			switch enemy {
			case "A":
				score += papper
			case "B":
				score += scisors
			case "C":
				score += rock
			}
		}

	}

	fmt.Println(score)
}
