package models

import (
	"time"
	"web_gin/dao"
)

type Vote struct {
	Id       int
	UserId   int
	PlayerId int
	AddTime  int64
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId int, playerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id=? AND player_id=?", userId, playerId).First(&vote).Error
	return vote, err
}
func AddVote(userId int, playerId int) (int, error) {
	vote := Vote{UserId: userId, PlayerId: playerId, AddTime: time.Now().Unix()}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
