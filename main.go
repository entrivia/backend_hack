package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Comment string `json:"comment"`
}

var users []User
var tasks []Task

func main() {
	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/task", createTask)
	r.GET("/task/:id", getTask)
	r.POST("/task/:id/comment", addComment)
	r.GET("/workers", getWorkers)
	r.GET("/tasks", getAllTasks)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, user)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}
func createTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func getTask(c *gin.Context) {
	id := c.Param("id")
	for _, task := range tasks {
		if strconv.Itoa(task.ID) == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func addComment(c *gin.Context) {
	id := c.Param("id")
	var comment struct {
		Comment string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, task := range tasks {
		if strconv.Itoa(task.ID) == id {
			tasks[i].Comment = comment.Comment
			c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func getWorkers(c *gin.Context) {
	// Assuming that workers are the users who have created tasks
	var workers []User
	for _, task := range tasks {
		for _, user := range users {
			if task.Username == user.Username {
				workers = append(workers, user)
				break
			}
		}
	}
	c.JSON(http.StatusOK, workers)
}

func getAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}
