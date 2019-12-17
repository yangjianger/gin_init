package models

type PayLog struct {
	Id      int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Orderid string `json:"orderid" xorm:"comment('订单id') VARCHAR(32)"`
	Meta    string `json:"meta" xorm:"comment('支付返回信息') TEXT"`
	ErrMsg  string `json:"err_msg" xorm:"default '' comment('错误消息') VARCHAR(255)"`
}
