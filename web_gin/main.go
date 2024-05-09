package main

import (
	"web_gin/router"
)

func main() {
	r := router.Router()
	// r := gin.Default()
	// user := r.Group("/user")
	// {
	// 	user.POST("/list", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "user list")
	// 	})

	// 	user.PUT("/add", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "user add")
	// 	})
	// 	user.DELETE("/delete", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "user delele ")
	// 	})
	// }

	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	//panic("11")

	//fmt.Println(4)
	r.Run(":8888")
	//panic("11")
}
