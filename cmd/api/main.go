package main

import ( 
	"linked-list-api/src/linked_list"
	"net/http"
	"github.com/gin-gonic/gin"
)

var headPointer = linked_list.Init()

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, linked_list.ToArray(headPointer))
}

func reverseLinkedList(c *gin.Context) {
	headPointer = linked_list.Reverse(headPointer)
	c.IndentedJSON(http.StatusOK, linked_list.ToArray(headPointer))
}

func getStudentById(c *gin.Context){
	student := linked_list.FindByID(headPointer, c.Param("id"))
	
	if student == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, student)
}

func createStudent(c *gin.Context){
	var newStudent linked_list.Student
	err := c.BindJSON(&newStudent)

	if err != nil {
		return
	}
	
	linked_list.Append(headPointer, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/reverse", reverseLinkedList)
	router.GET("/students/:id", getStudentById)
	router.POST("/students", createStudent)
	router.Run("localhost:8080")
}