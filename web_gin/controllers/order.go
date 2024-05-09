package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}
type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (u OrderController) GETList(c *gin.Context) {
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wanglll")

	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	search := &Search{}
	err := c.BindJSON(&search)

	if err == nil {
		returnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	returnError(c, 0, gin.H{"err": err})
}
