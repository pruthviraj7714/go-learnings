package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{}

var nextId = 1

func addUserHandler(c *gin.Context) {
	var UserInput User

	err := c.ShouldBindBodyWithJSON(&UserInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	UserInput.Id = nextId
	nextId++

	// user := User{
	// 	Id:   nextId,
	// 	Name: UserInput.Name,
	// 	Age:  UserInput.Age,
	// }

	users = append(users, UserInput)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Successfully created",
		"User":    UserInput,
	})
}

func getSingleUser(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id Passed",
		})
		return
	}

	for _, user := range users {
		if user.Id == i {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user with given id not found"})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	userId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Id Passed",
		})
		return
	}

	var Input User

	if err := c.ShouldBindBodyWithJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, u := range users {
		if u.Id == userId {
			users[i].Age = Input.Age
			users[i].Name = Input.Name
			c.JSON(http.StatusOK, users[i])
			return
		}

	}

	c.JSON(http.StatusNotFound, gin.H{"message": "user with given id not found"})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"Message": "User successfully deleted",
			})
			return

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func listUsers(c *gin.Context) {

	if len(users) > 0 {
		c.JSON(http.StatusOK, users)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "No users found",
	})
}

func main() {
	router := gin.Default()

	router.POST("/users", addUserHandler)
	router.GET("/users", listUsers)
	router.GET("/user/:id", getSingleUser)
	router.PUT("/user/:id", updateUser)
	router.DELETE("/user/:id", deleteUser)

	router.Run(":8080")

}
