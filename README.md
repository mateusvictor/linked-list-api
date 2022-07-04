# linked-list-api
First RESTful API made in Go Lang with the Gin framework that implements the Linked List Data Structure.

This projects was designed to have a basic understanding of the Go Language including its basic syntax, packages, structs, pointers and the Gin framework.

In order to practice pointers and memory manipulation, a basic implementation of an in-memory linked list was made that stores information about the main entity: Student.

## Basic Usage: Linked List impl 
```shell
➜  linked-list-api git:(main) ✗ go run cmd/api/main.go

Appending nodes
Student's node: {2 Jorge Vieira 82} appended
Student's node: {3 Cecilia Assis 213} appended
Student's node: {4 João Carlos 21} appended
Student's node: {5 Mario Santa 13} appended
Student's node: {6 Roberto Castro 23} appended

Printing linked list
Student: {1 Mateus Silva 17}
Next: {{2 Jorge Vieira 82} 0xc00007e0c0}

Student: {2 Jorge Vieira 82}
Next: {{3 Cecilia Assis 213} 0xc00007e140}

Student: {3 Cecilia Assis 213}
Next: {{4 João Carlos 21} 0xc00007e1c0}

Student: {4 João Carlos 21}
Next: {{5 Mario Santa 13} 0xc00007e240}

Student: {5 Mario Santa 13}
Next: {{6 Roberto Castro 23} <nil>}

Student: {6 Roberto Castro 23}
Next: <nil>
```