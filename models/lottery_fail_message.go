package models

type LotteryFailMessage struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Lid     int    `json:"lid" xorm:"default 0 comment('开奖id') index INT(11)"`
	Message string `json:"message" xorm:"comment('错误信息') TEXT"`
	Addtime string `json:"addtime" xorm:"comment('添加时间') VARCHAR(32)"`
}
