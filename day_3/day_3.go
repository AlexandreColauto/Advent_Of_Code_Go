package main

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

	data, err := input.Strings(2022, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data")
	fmt.Println(data[2])

	total_priority := 0
	for i := 0; i < len(data)-2; i += 3 {

		fmt.Println(i)
		v := data[i]
		partener_1 := data[i+1]
		partener_2 := data[i+2]

		for _, char := range v {
			if strings.ContainsRune(partener_1, char) && strings.ContainsRune(partener_2, char) {
				if rune('a') <= char && char <= rune('z') {
					priority := char - rune('a') + 1
					total_priority += int(priority)
					fmt.Println(string(char))
					break
				} else {

					priority := char - rune('A') + 27
					total_priority += int(priority)
					fmt.Println(string(char))
					break
				}

			}
		}
	}

	fmt.Println(total_priority)

}
