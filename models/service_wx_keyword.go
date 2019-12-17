package models

type ServiceWxKeyword struct {
	Id          int    `json:"id" xorm:"not null pk autoincr MEDIUMINT(8)"`
	Type        int    `json:"type" xorm:"not null default 0 TINYINT(1)"`
	Class       string `json:"class" xorm:"not null default 'text' VARCHAR(20)"`
	Admin       string `json:"admin" xorm:"not null VARCHAR(30)"`
	Addtime     int    `json:"addtime" xorm:"not null default 0 INT(10)"`
	Starttime   int    `json:"starttime" xorm:"not null default 0 INT(10)"`
	Endtime     int    `json:"endtime" xorm:"not null default 0 INT(10)"`
	Isshow      int    `json:"isshow" xorm:"not null default 1 TINYINT(1)"`
	Keyword     string `json:"keyword" xorm:"not null VARCHAR(60)"`
	Data        string `json:"data" xorm:"not null MEDIUMTEXT"`
	Description string `json:"description" xorm:"not null default ' ' VARCHAR(120)"`
}
