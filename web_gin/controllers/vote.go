package controllers

import (
	"strconv"
	"time"
	"web_gin/cache"
	"web_gin/models"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {

	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)
	if userId == 0 || playerId == 0 {
		returnError(c, 4001, "请输入正确的信息")
		return
	}
	user, _ := models.GetUserInfoById(userId)
	if user.Id == 0 {
		returnError(c, 4001, "投票用户不存在")
		return
	}
	player, _ := models.GetPlayerInfoById(playerId)
	if player.Id == 0 {
		returnError(c, 4001, "参赛选手不存在")
		return
	}
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		returnError(c, 4001, "已经投过票了")
		return
	}
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		models.UpdatePlayer(playerId)
		//fmt.Println(UPdateErr)
		redisKey := "rank:" + strconv.Itoa(player.Aid)
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerId))
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		returnSuccess(c, 0, "投票成功", rs, 1)
		return
	}
	returnError(c, 4004, "投票错误")
	//return
}
