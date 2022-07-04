package linked_list

import "fmt"

type Student struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int32 `json:"age"`
}

type Node struct {
	Student Student `json:"student"`
	Next *Node `json:"next"`
}

var students = []Student {
	{ID: "1", FirstName: "Mateus", LastName: "Silva", Age: 17},
	{ID: "2", FirstName: "Jorge", LastName: "Vieira", Age: 82},
	{ID: "3", FirstName: "Cecilia", LastName: "Assis", Age: 213},
	{ID: "4", FirstName: "João", LastName: "Carlos", Age: 21},
	{ID: "5", FirstName: "Mario", LastName: "Santa", Age: 13},
	{ID: "6", FirstName: "Roberto", LastName: "Castro", Age: 23},
}

func Init() (*Node) {
	head := Node{}
	head.Student = students[0]

	fmt.Println("Appending nodes")

	for _, v := range students[1:] {
		Append(&head, v)
		fmt.Println("Student's node:", v, "appended")
	}
	
	fmt.Println()

	return &head
}

func Append(head *Node, new_std Student){
	current := head

	for (*current).Next != nil {
		current = current.Next
	}

	new_node := Node{
		Student: new_std, Next: nil,
	}
	
	current.Next = &new_node
}

func Print(head *Node) {
	current := head

	for current != nil {
		print_node(current)
		current = current.Next
	}
}

func ToArray(head *Node) ([]Student) {
	var students []Student
	current := head

	for current != nil {
		students = append(students, current.Student)
		current = current.Next
	}

	return students
}

func FindByID(head *Node, id string) (node *Node){
	current := head

	for current != nil {
		if (*current).Student.ID == id {
			return current
		}
		current = current.Next
	}

	return nil
}

func Reverse(head *Node) (*Node) {
	current := head
	var next, prev *Node

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func print_node(node *Node){
	fmt.Println("Student:", node.Student)
	fmt.Print("Next: ")
	
	if node.Next != nil {
		fmt.Println(*node.Next)
	 } else {
		fmt.Println(nil)
	}

	fmt.Println()
}