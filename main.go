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
	fmt.Println("before  crate ")
	fmt.Println(s2)
	//for i2, j2 := 0, len(s2)-1; i2 < j2; i2, j2 = i2+1, j2-1 {
	//	s2[i2], s2[j2] = s2[j2], s2[i2]
	//}

	fmt.Println("before  crate ")
	fmt.Println(s2)
	for _, v2 := range s2 {
		for i2, val2 := range v2 {
			if i2%4 < 3 {
				crate2 += string(val2)
			}

			if i2%4 == 3 {
				if crate2 != "   " {
					crates2[crate_index2] = append(crates2[crate_index2], crate2)
				}
				crate2 = ""
				crate_index2++

			}
		}
		if crate2 != "   " {
			crates2[crate_index2] = append(crates2[crate_index2], crate2)
		}
		crate2 = ""
		crate_index2 = 0

	}
	//fmt.Println(crates)
	fmt.Println("before  crate ")
	for _, cr2 := range crates2 {

		fmt.Println(cr2)
	}

	for _, in2 := range instructions2 {
		//in := instructions[0]
		ins_split := strings.Split(in2, " ")
		amount, _ := strconv.Atoi(ins_split[1])
		from, _ := strconv.Atoi(ins_split[3])
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
