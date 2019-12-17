package models

type LotteryUvData struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Lid     int    `json:"lid" xorm:"comment('抽奖id') index INT(11)"`
	Uid     int    `json:"uid" xorm:"comment('用户id') index INT(11)"`
	Ip      string `json:"ip" xorm:"comment('ip') VARCHAR(20)"`
	Hash    string `json:"hash" xorm:"comment('hash') unique VARCHAR(32)"`
	AddTime int    `json:"add_time" xorm:"default 0 comment('查看时间') INT(11)"`
}
