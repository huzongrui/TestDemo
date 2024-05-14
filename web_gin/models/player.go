package models

import (
	"web_gin/dao"

	"github.com/jinzhu/gorm"
)

type Player struct {
	Id          int
	Aid         int
	Ref         string
	Nickname    string
	Declaration string
	Avatar      string
	Score       int `json:"score"`
}

func (Player) TableName() string {
	return "player"
}

func GetPlayers(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid=?", aid).Order(sort).Find(&players).Error
	return players, err
}
func GetPlayerInfoById(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id=?", id).Find(&player).Error
	return player, err
}

func UpdatePlayer(id int) {

	// dao.Db.Model(&player).Where("id=?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))

	var player Player
	//dao.Db.Model(&player).Updates(Player{Id: id, Score: player.Score + 1})
	dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + 1"))

}
