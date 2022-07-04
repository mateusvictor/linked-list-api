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