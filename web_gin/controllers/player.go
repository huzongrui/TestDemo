package controllers

import (
	"strconv"
	"time"
	"web_gin/cache"
	"web_gin/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		returnError(c, 4004, "没有相关的信息")
		return
	}
	returnSuccess(c, 0, "sucesss", rs, 1)
}
func (P PlayerController) GetPlayerRank(c *gin.Context) {

	// err := cache.Rdb.Set(cache.Rctx, "name", "zhangfei", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	//var redisKey string
	redisKey := "rank:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfoById(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		returnSuccess(c, 0, "sucesss", players, 1)

		return
	}

	rsDb, errDb := models.GetPlayers(aid, "score desc")
	if errDb == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		returnSuccess(c, 0, "sucesss", rsDb, 1)
		return
	}
	returnError(c, 4004, "没有相关的信息")

}
