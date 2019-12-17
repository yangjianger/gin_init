package models

type AdminLog struct {
	Uid        int    `json:"uid" xorm:"not null INT(11)"`
	Controller string `json:"controller" xorm:"not null VARCHAR(120)"`
	Action     string `json:"action" xorm:"not null VARCHAR(120)"`
	Url        string `json:"url" xorm:"VARCHAR(320)"`
	AddTime    int    `json:"add_time" xorm:"not null INT(10)"`
	RequestIp  string `json:"request_ip" xorm:"default '' comment('请求ip') VARCHAR(20)"`
	Content    string `json:"content" xorm:"TEXT"`
}
