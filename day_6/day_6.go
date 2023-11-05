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

	data, err := input.Strings(2022, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data")
	fmt.Println("Part 1")
	//part 1 package of 4 non repeating digits
	//package_strings := data[0][:4]

	//part 2 package of 14 non repeating digits
	// the logic is the same
	package_strings := data[0][:14]

	var last_package = package_strings
	fmt.Printf("package %s\n", package_strings)
	for i, c := range data[0][:30] {
		last_package = package_strings
		package_strings = package_strings[1:] + string(c)
		duplicate := false
		duplicate = duplicate_rune(package_strings)
		if !duplicate {
			fmt.Printf("old Pacakge %s\n", last_package)
			fmt.Printf("new Pacakge %s\n", package_strings)
			fmt.Printf("found in index %d", i)
			break
		}
	}

}
func duplicate_rune(s string) bool {
	for _, r := range s {
		rune_count := strings.Count(s, string(r))
		if rune_count > 1 {
			return true
		}
	}
	return false
}
