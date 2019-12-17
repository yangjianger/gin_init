package models

type LotterySponsorWxcode struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	WxcodeName  string `json:"wxcode_name" xorm:"default '' comment('小程序名称') VARCHAR(32)"`
	WxcodePath  string `json:"wxcode_path" xorm:"default '' comment('小程序路径') VARCHAR(255)"`
	WxcodeAppid string `json:"wxcode_appid" xorm:"default '' comment('小程序appid') VARCHAR(120)"`
	SponsorName string `json:"sponsor_name" xorm:"default '' comment('商家名称') VARCHAR(120)"`
}
