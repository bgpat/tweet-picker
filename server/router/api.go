package router

import (
	"github.com/bgpat/tweet-picker/server/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

func InitializeAPI(r *gin.RouterGroup) {
	r.GET("/tweets", controllers.GetTweets)
	r.GET("/tweets/:id", controllers.GetTweet)

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:user_id", controllers.GetUser)

	r.GET("/users/:user_id/tweets", controllers.GetUserTweets)
	r.GET("/users/:user_id/tweets/:id", controllers.GetUserTweet)
}
