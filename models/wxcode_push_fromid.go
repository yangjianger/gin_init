package models

type WxcodePushFromid struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	FromId     string `json:"from_id" xorm:"comment('微信fromid') VARCHAR(120)"`
	OpenId     string `json:"open_id" xorm:"comment('微信openid') index VARCHAR(120)"`
	Addtime    int    `json:"addtime" xorm:"comment('添加时间') index INT(11)"`
	Status     int    `json:"status" xorm:"default 0 comment('0初始状态，1待发送，2已使用，或过期') TINYINT(1)"`
	PushId     int    `json:"push_id" xorm:"comment('push_id') index INT(11)"`
	IsArrive   int    `json:"is_arrive" xorm:"default 0 comment('推送是否到达0默认 1已到达') TINYINT(1)"`
	IsClick    int    `json:"is_click" xorm:"default 0 comment('是否点击 0没有 1点击') TINYINT(1)"`
	ClickTime  int    `json:"click_time" xorm:"default 0 comment('点击时间') INT(11)"`
	ArriveTime int    `json:"arrive_time" xorm:"default 0 comment('到达时间') INT(11)"`
	Days       int    `json:"days" xorm:"INT(11)"`
	Uid        int    `json:"uid" xorm:"comment('用户id') index INT(11)"`
	Path       string `json:"path" xorm:"comment('小程序路径') VARCHAR(255)"`
}
