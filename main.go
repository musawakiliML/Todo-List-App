package main

import (
	"fmt"
)

func main() {
	// Variable declaration
	var todo1 = "Learning Go!"
	var todo2 = "Building a project in Go"
	var todo3 = "Writing a blog on each project"

	var todoItems = []string{todo1, todo2, todo3}

	fmt.Printf("#### Welcome to my Go Learning Project ####\n")

	// Call function for Displaying Todos
	printTodos(todoItems)

	todoItems = addTodo(todoItems, "Solve Programming Challenge")

	fmt.Printf("Updated List!")

	printTodos(todoItems)

}

func printTodos(todoItems []string) {
	// Function to Print Todo List
	fmt.Println("List of My Todo Lists")

	for index, task := range todoItems {
		//fmt.Println(index, task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTodo(todoItems []string, newTodo string) []string {
	// Function to add Todo list
	var updateTodoList = append(todoItems, newTodo)

	return updateTodoList

}
