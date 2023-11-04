package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data")
	fmt.Println("Part 1")
	s := data[0:8]
	instructions := data[10:]
	crate := ""
	crate_index := 0
	var crates = [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	//reverse array
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	for _, v := range s {
		for i, val := range v {
			if i%4 < 3 {
				crate += string(val)
			}

			if i%4 == 3 {
				if crate != "   " {
					crates[crate_index] = append(crates[crate_index], crate)
				}
				crate = ""
				crate_index++

			}
		}
		if crate != "   " {
			crates[crate_index] = append(crates[crate_index], crate)
		}
		crate = ""
		crate_index = 0

	}
	//fmt.Println(crates)
	fmt.Println("before  crate ")
	for _, cr := range crates {

		fmt.Println(cr)
	}

	for _, in := range instructions {
		//in := instructions[0]
		ins_split := strings.Split(in, " ")
		amount, _ := strconv.Atoi(ins_split[1])
		from, _ := strconv.Atoi(ins_split[3])
		from--
		to, _ := strconv.Atoi(ins_split[5])
		to--
		for index := 0; index < amount; index++ {
			crates[to] = append(crates[to], crates[from][len(crates[from])-1])
			crates[from] = crates[from][:len(crates[from])-1]
		}
	}
	fmt.Println("Final crate ")

	for _, cr := range crates {

		fmt.Println(cr)
	}

	fmt.Println("Part 2")
	s2 := data[0:8]
	instructions2 := data[10:]
	crate2 := ""
	crate_index2 := 0
	var crates2 = [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	//reverse array
	// since go perform inplace there is no need to revert the array again
	//for i2, j2 := 0, len(s2)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
	//	s2[i2], s2[j2] = s2[j2], s2[i2]
	//}

	for _, v2 := range s2 {
		for i2, val2 := range v2 {
			//the first 3 characters is our crate
			if i2%4 < 3 {
				crate2 += string(val2)
			}

			// the 4 indicate split between crates, so we save our crate into the right index and reset the variable for the next crate value.

			if i2%4 == 3 {
				//check for empty crates
				if crate2 != "   " {
					crates2[crate_index2] = append(crates2[crate_index2], crate2)
				}
				crate2 = ""
				crate_index2++

			}
		}
		//since there is no last space the last crate should be appended outside the loop
		if crate2 != "   " {
			crates2[crate_index2] = append(crates2[crate_index2], crate2)
		}
		crate2 = ""
		crate_index2 = 0

	}
	fmt.Println("before instructions  crate ")
	for _, cr2 := range crates2 {
		fmt.Println(cr2)
	}

	for _, in2 := range instructions2 {
		ins_split := strings.Split(in2, " ")
		amount, _ := strconv.Atoi(ins_split[1])
		from, _ := strconv.Atoi(ins_split[3])
		//since the crates goes from 1 to 9 we need to sub 1 to turn them into indexes
		from--
		to, _ := strconv.Atoi(ins_split[5])
		to--
		var removed = crates2[from][len(crates2[from])-amount:]
		crates2[to] = append(crates2[to], removed...)
		crates2[from] = crates2[from][:len(crates2[from])-amount]
	}
	fmt.Println("Final crate ")
	var result = ""
	for _, cr2 := range crates2 {

		fmt.Println(cr2)
		result = result + cr2[len(cr2)-1]
	}
	fmt.Printf("Final result %s ", result)
}
