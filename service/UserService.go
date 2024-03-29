package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

// Get User
func FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

// Post User
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully posted")
}

// delete User
func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	for _, user := range userList {
		log.Println(user)
		userList = append(userList[:userId], userList[userId+1:]...)
		c.JSON(http.StatusOK, "Successfully deleted")
		return
	}
	c.JSON(http.StatusNotFound, "Error")
}

func PutUser(c *gin.Context) {
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	for key, user := range userList {
		if userId == user.Id {
			userList[key] = beforeUser
			log.Println(userList[key])
			c.JSON(http.StatusOK, "Successfully")
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}
