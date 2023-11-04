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

	data, err := input.Strings(2022, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data")
	fmt.Println(data[len(data)-1])
	fmt.Println("Part 1")
	fmt.Println(data[len(data)-1])
	overlaping_groups := 0
	for _, v := range data {
		elf_group := strings.Split(v, ",")

		first_group_bounds := strings.Split(elf_group[0], "-")
		second_group_bounds := strings.Split(elf_group[1], "-")
		s1, err := strconv.Atoi(first_group_bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		e1, err := strconv.Atoi(first_group_bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		s2, err := strconv.Atoi(second_group_bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		e2, err := strconv.Atoi(second_group_bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		if (s1 <= s2 && e2 <= e1) || (s2 <= s1 && e1 <= e2) {
			overlaping_groups += 1
		}
	}
	fmt.Println("Total")
	fmt.Println(overlaping_groups)
	fmt.Println("Part 2")
	overlaping_groups = 0
	for _, v := range data {
		elf_group := strings.Split(v, ",")

		first_group_bounds := strings.Split(elf_group[0], "-")
		second_group_bounds := strings.Split(elf_group[1], "-")
		s1, err := strconv.Atoi(first_group_bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		e1, err := strconv.Atoi(first_group_bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		s2, err := strconv.Atoi(second_group_bounds[0])
		if err != nil {
			log.Fatal(err)
		}
		e2, err := strconv.Atoi(second_group_bounds[1])
		if err != nil {
			log.Fatal(err)
		}
		if e1 >= s2 && s1 <= e2 {
			overlaping_groups += 1
		}
	}
	fmt.Println("Total")
	fmt.Println(overlaping_groups)
}
