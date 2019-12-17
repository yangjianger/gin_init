package models

type WxcodeToken struct {
	Id    int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Token string `json:"token" xorm:"not null index VARCHAR(512)"`
}
