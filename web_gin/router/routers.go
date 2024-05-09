package router

import (
	"web_gin/config"
	"web_gin/controllers"
	"web_gin/pkg/logger"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
		// user.POST("/list", controllers.UserController{}.GETList)
		// user.POST("/add", controllers.UserController{}.AddUserTest)
		// user.POST("/update", controllers.UserController{}.UpdateUserTest)
		// user.POST("/delete", controllers.UserController{}.DeleleUserTest)
		// user.POST("/list/test", controllers.UserController{}.GetUserListTest)
	}
	orer := r.Group("/order")
	{
		orer.POST("/list", controllers.OrderController{}.GETList)
	}
	return r
}
