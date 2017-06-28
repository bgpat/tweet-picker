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
	errs := db.Unscoped().Order("deleted_at desc").Find(&tweets).GetErrors()
	if ReturnErrors(c, errs) {
		return
	}
	c.JSON(http.StatusOK, tweets)
}

func GetTweet(c *gin.Context) {
	id := c.Param("id")
	tweet := models.Tweet{}
	errs := db.Unscoped().First(&tweet, id).GetErrors()
	if ReturnErrors(c, errs) {
		return
	}
	c.JSON(http.StatusOK, &tweet)
}

func GetUsers(c *gin.Context) {
	users := []*models.User{}
	errs := db.Order("updated_at desc").Find(&users).GetErrors()
	if ReturnErrors(c, errs) {
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("user_id")
	user := models.User{}
	errs := db.First(&user, id).GetErrors()
	if ReturnErrors(c, errs) {
		return
	}
	c.JSON(http.StatusOK, &user)
}

func GetUserTweets(c *gin.Context) {
	userID := c.Param("user_id")
	tweets := []*models.Tweet{}
	errs := db.Unscoped().Order("deleted_at desc").Find(&tweets, "user_id = ?", userID).GetErrors()
	if ReturnErrors(c, errs) {
		return
	}
	c.JSON(http.StatusOK, tweets)
}

func GetUserTweet(c *gin.Context) {
	id := c.Param("id")
	userID := c.Param("user_id")
	tweet := models.Tweet{}
	errs := db.Unscoped().First(&tweet, "id = ? AND user_id = ?", id, userID).GetErrors()
	if ReturnErrors(c, errs) {
		return
	}
	c.JSON(http.StatusOK, &tweet)
}
