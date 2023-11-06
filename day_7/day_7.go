package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

// define the file type
type file struct {
	size int
	name string
}

// define the dir type
type dir struct {
	childs []*dir
	files  []file
	parent *dir
	size   int
	name   string
}

var current_dir *dir
var current_dir_to_remove *dir

var dir_tree = make([]dir, 0)

const COMMAND = 1
const DIR = 2
const FILE = 3
const MAX_SIZE = 100000
const NEED_SIZE = 9192532

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data")
	fmt.Println(data[0])
	//loop over inputs
	for _, in := range data[:] {
		//decodify input
		decode_input(in)

	}
	get_results(current_dir)

}
func decode_input(in string) {
	//split the input
	input_arr := strings.Split(in, " ")
	//get input type
	if input_arr[0] == "$" {
		if input_arr[1] == "cd" {
			//if is cd change the current_dir to the folder specified, and add the folder to the folder array if it doesn exist
			current_dir = validate_and_change_dir(input_arr[2])
		}
	} else if input_arr[0] == "dir" {
		create_new_dir(input_arr[1], current_dir)
	} else if v, err := strconv.Atoi(input_arr[0]); err == nil {
		add_file(input_arr[1], v)
	}

}
func validate_and_change_dir(directory_name string) *dir {
	//check if is need to go to the parent folder
	if directory_name == ".." {
		return current_dir.parent

	}
	//check if directory alread exists
	if current_dir != nil {
		for _, directory := range current_dir.childs {
			if directory.name == directory_name {
				return directory
			}
		}
	}
	// if doesnt exist create new dir append to current_dir child and append to the tree.
	return create_new_dir(directory_name, current_dir)

}

func create_new_dir(directory_name string, current *dir) *dir {
	var new_dir = &dir{
		name:   directory_name,
		parent: current,
	}
	if current != nil {
		current.childs = append(current.childs, new_dir)
	}
	return new_dir
}

func add_file(file_name string, file_size int) {
	new_file := file{
		name: file_name,
		size: file_size,
	}
	current_dir.files = append(current_dir.files, new_file)
	current_dir.size += file_size
	add_size_to_parent(file_size)
}

func add_size_to_parent(size int) {
	c_dir := current_dir

	for {

		if c_dir.parent == nil {
			break
		}
		c_dir = c_dir.parent
		c_dir.size += size
	}
}

func get_results(current *dir) {
	c_dir := current_dir

	for {

		if c_dir.parent == nil {

			break
		}
		c_dir = c_dir.parent
	}
	result := make([]*dir, 0)
	tree_traversal(c_dir, &result)
	total := 0
	for _, r := range result {
		total += r.size
	}
	fmt.Printf("----- Part 1 total %d  directories size under  %d -----\n", len(result), MAX_SIZE)
	fmt.Printf("Total part 1  %d\n", total)
	fmt.Println(" ---- Part 2 -----")
	fmt.Printf(" top level with a size %d\n", c_dir.size)
	fmt.Printf(" and a space left of (70000000-size) %d\n", 70000000-c_dir.size)
	fmt.Printf(" and a space required for (30000000) %d\n", 30000000-(70000000-c_dir.size))
	fmt.Printf("directory name to remove to free enought space (part 2)  %s\n", current_dir_to_remove.name)
	fmt.Printf("directory size to remove to free enought space (part 2)  %d\n", current_dir_to_remove.size)

}

func tree_traversal(curr *dir, results *[]*dir) *[]*dir {

	//base case ; stop recursing
	if curr == nil {
		return results
	}

	if curr.size < MAX_SIZE {
		*results = append(*results, curr)
	}
	if curr.size >= NEED_SIZE {
		if current_dir_to_remove == nil {
			current_dir_to_remove = curr
		} else if current_dir_to_remove.size > curr.size {
			current_dir_to_remove = curr
		}
	}
	for _, child := range curr.childs {
		tree_traversal(child, results)
	}

	return results
}
