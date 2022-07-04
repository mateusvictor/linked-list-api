package main

import ( 
	"fmt"
	"linked-list-api/src/linked_list"
)

var students = []linked_list.Student {
	{ID: "1", FirstName: "Mateus", LastName: "Silva", Age: 17},
	{ID: "2", FirstName: "Jorge", LastName: "Vieira", Age: 82},
	{ID: "3", FirstName: "Cecilia", LastName: "Assis", Age: 213},
	{ID: "4", FirstName: "João", LastName: "Carlos", Age: 21},
	{ID: "5", FirstName: "Mario", LastName: "Santa", Age: 13},
	{ID: "6", FirstName: "Roberto", LastName: "Castro", Age: 23},
}

func main() {
	head := linked_list.Node{}
	head.Student = students[0]

	fmt.Println("Appending nodes")

	for _, v := range students[1:] {
		linked_list.Append(&head, v)
		fmt.Println("Student's node:", v, "appended")
	}

	fmt.Println()
	fmt.Println("Printing linked list")
	
	linked_list.Print(&head)
}