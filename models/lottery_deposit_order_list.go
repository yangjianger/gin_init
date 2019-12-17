package models

type LotteryDepositOrderList struct {
	Orderid         string `json:"orderid" xorm:"not null pk default '' comment('订单号') VARCHAR(32)"`
	Lid             int    `json:"lid" xorm:"default 0 comment('抽奖id') index INT(11)"`
	Deposit         string `json:"deposit" xorm:"default 0.00 comment('押金') DECIMAL(8,2)"`
	Uid             int    `json:"uid" xorm:"default 0 comment('用户id') INT(11)"`
	PayTime         int    `json:"pay_time" xorm:"default 0 comment('支付时间') INT(11)"`
	PayErrorCode    string `json:"pay_error_code" xorm:"default '' comment('支付错误码') VARCHAR(32)"`
	PayErrorMsg     string `json:"pay_error_msg" xorm:"default '' comment('支付退款信息') VARCHAR(120)"`
	PayMeta         string `json:"pay_meta" xorm:"comment('支付返回字符串') TEXT"`
	PayTradeNo      string `json:"pay_trade_no" xorm:"default '' comment('交易流水号') VARCHAR(120)"`
	RefundTime      int    `json:"refund_time" xorm:"default 0 comment('退款时间') INT(11)"`
	RefundErrorCode string `json:"refund_error_code" xorm:"default '' comment('退款错误码') VARCHAR(32)"`
	RefundErrorMsg  string `json:"refund_error_msg" xorm:"default '' comment('退款错误信息') VARCHAR(120)"`
	PayStatus       int    `json:"pay_status" xorm:"default 1 comment('支付状态 1未支付 2已支付 3已退款') TINYINT(1)"`
	AddTime         int    `json:"add_time" xorm:"default 0 comment('订单添加时间') INT(11)"`
	RefundMeta      string `json:"refund_meta" xorm:"comment('退款返回数据') TEXT"`
	RefundId        string `json:"refund_id" xorm:"default '' comment('退款id') VARCHAR(64)"`
	RefundStatus    int    `json:"refund_status" xorm:"default 1 comment('退款状态1 未退  2已退款 3退款失败') TINYINT(4)"`
}
