package models

type SmsLog struct {
	Tel      string `json:"tel" xorm:"not null VARCHAR(12)"`
	Ip       string `json:"ip" xorm:"not null VARCHAR(15)"`
	Num      int    `json:"num" xorm:"not null default 0 INT(2)"`
	Lasttime int    `json:"lasttime" xorm:"not null INT(10)"`
}
