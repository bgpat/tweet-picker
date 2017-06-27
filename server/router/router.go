package router

import (
	"github.com/bgpat/tweet-picker/database"
	"github.com/bgpat/tweet-picker/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func Initialize(r *gin.Engine) {
	db := database.Default()

	r.GET("/tweets", func(c *gin.Context) {
		tweets := []*models.Tweet{}
		db.Unscoped().Find(&tweets)
		c.JSON(200, tweets)
	})
	r.GET("/users", func(c *gin.Context) {
		users := []*models.User{}
		db.Find(&users)
		c.JSON(200, users)
	})
	r.GET("/users/:id", func(c *gin.Context) {
		user := models.User{}
		id := c.Param("id")
		db.First(&user, id)
		c.JSON(200, &user)
	})
	r.GET("/users/:id/tweets", func(c *gin.Context) {
		tweets := []*models.Tweet{}
		userID := c.Param("id")
		db.Unscoped().Find(&tweets, "user_id = ?", userID)
		c.JSON(200, tweets)
	})
}
