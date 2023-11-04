package day_1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/echojc/aocutil"
)

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 1)
	if err != nil {
		log.Fatal(err)
	}

	biggest := 0
	second := 0
	third := 0
	sum := 0
	for _, value := range data {
		if value == "" {
			if biggest < sum {
				third = second
				second = biggest
				biggest = sum

			} else if second < sum {
				third = second
				second = sum
			} else if third < sum {
				third = sum
			}
			sum = 0
		} else {
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		}

	}
	fmt.Println("biggest")
	fmt.Println(biggest)
	fmt.Println("second")
	fmt.Println(second)
	fmt.Println("third")
	fmt.Println(third)
	fmt.Println("result")
	fmt.Println(third + second + biggest)
}
