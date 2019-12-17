package models

type ServiceWxBottomMenu struct {
	Id        int    `json:"id" xorm:"not null pk autoincr SMALLINT(5)"`
	Listorder int    `json:"listorder" xorm:"not null default 0 SMALLINT(5)"`
	Parentid  int    `json:"parentid" xorm:"not null default 0 SMALLINT(5)"`
	Isshow    int    `json:"isshow" xorm:"not null default 1 TINYINT(1)"`
	Type      string `json:"type" xorm:"not null default 'view' VARCHAR(32)"`
	Class     string `json:"class" xorm:"not null default 'button' VARCHAR(32)"`
	Name      string `json:"name" xorm:"not null VARCHAR(60)"`
	Url       string `json:"url" xorm:"not null VARCHAR(120)"`
}
