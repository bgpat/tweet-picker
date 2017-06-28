package controllers

import (
	"net/http"

	"github.com/bgpat/tweet-picker/database"
	"github.com/bgpat/tweet-picker/models"
	"gopkg.in/gin-gonic/gin.v1"
)

var db = database.Default()

func GetTweets(c *gin.Context) {
	tweets := []*models.Tweet{}
	err := db.Unscoped().Find(&tweets)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tweets)
}

func GetTweet(c *gin.Context) {
	id := c.Param("id")
	tweet := models.Tweet{}
	err := db.Unscoped().First(&tweet, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, &tweet)
}

func GetUsers(c *gin.Context) {
	users := []*models.User{}
	err := db.Find(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("user_id")
	user := models.User{}
	err := db.First(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetUserTweets(c *gin.Context) {
	userID := c.Param("user_id")
	tweets := []*models.Tweet{}
	err := db.Unscoped().Find(&tweets, "user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tweets)
}

func GetUserTweet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Param("user_id")
	tweet := models.Tweet{}
	err := db.Unscoped().First(&tweet, "id = ? AND user_id = ?", id, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, &tweet)
}
