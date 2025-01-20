package main

import "fmt"

func main() {

	var todo1 = "Learning Go!"
	var todo2 = "Building a project in Go"
	var todo3 = "Writing a blog on each project"

	var todoItems = []string{todo1, todo2, todo3}

	fmt.Println("#### Welcome to my Go Learning Project ####")

	fmt.Println("List of My Todo Lists")

	for index, task := range todoItems {
		//fmt.Println(index, task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}
