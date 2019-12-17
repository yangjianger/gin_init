package models

type RequestLog struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Url     string `json:"url" xorm:"comment('错误请求参数') TEXT"`
	AddTime int    `json:"add_time" xorm:"comment('请求时间') INT(11)"`
	Method  string `json:"method" xorm:"index VARCHAR(120)"`
	Error   string `json:"error" xorm:"comment('错误信息') TEXT"`
}
