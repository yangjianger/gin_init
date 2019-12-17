package models

type LotteryOrderList struct {
	Orderid         string `json:"orderid" xorm:"not null pk comment('订单id') VARCHAR(32)"`
	Lid             int    `json:"lid" xorm:"default 0 comment('抽奖id') index INT(11)"`
	Uid             int    `json:"uid" xorm:"default 0 comment('用户id') index INT(11)"`
	Username        string `json:"username" xorm:"default '' comment('用户名') VARCHAR(120)"`
	Title           string `json:"title" xorm:"default '' comment('抽奖标题') VARCHAR(120)"`
	ReceiveUsername string `json:"receive_username" xorm:"default '' comment('收件人名称') VARCHAR(120)"`
	SendOrder       string `json:"send_order" xorm:"default '' comment('快递单号') VARCHAR(32)"`
	SendCompany     string `json:"send_company" xorm:"default '' comment('快递公司') VARCHAR(20)"`
	IsNotice        int    `json:"is_notice" xorm:"default 0 comment('是否通知用户1.已通知；0.未通知') TINYINT(1)"`
	AddTime         int    `json:"add_time" xorm:"default 0 comment('中奖时间') INT(11)"`
	EndTime         int    `json:"end_time" xorm:"default 0 comment('订单关闭时间') INT(11)"`
	ReceiveTel      string `json:"receive_tel" xorm:"default '' comment('收件人手机号') VARCHAR(12)"`
	ReceiveAddress  string `json:"receive_address" xorm:"comment('收件人地址') TEXT"`
	ReceiveCity     string `json:"receive_city" xorm:"default '' comment('收件人城市') VARCHAR(32)"`
	ReceiveRemarks  string `json:"receive_remarks" xorm:"default '' comment('收件人备注') VARCHAR(120)"`
	Status          int    `json:"status" xorm:"default 0 comment('订单状态 0正常 -1超时 1发货') index TINYINT(1)"`
	Mid             int    `json:"mid" xorm:"default 0 comment('玩法id') INT(11)"`
	IsAdmin         int    `json:"is_admin" xorm:"default 0 comment('是否是后台订单 0不是 1是') TINYINT(1)"`
	ReceiveAvatar   string `json:"receive_avatar" xorm:"default '' comment('头像') VARCHAR(120)"`
	LotteryType     int    `json:"lottery_type" xorm:"default 1 comment('奖品类型 1实物  2小程序优惠券  3购买折扣卷') TINYINT(4)"`
	TicketCode      string `json:"ticket_code" xorm:"default '' comment('优惠券码') VARCHAR(255)"`
	WxcodeAppid     string `json:"wxcode_appid" xorm:"default '' comment('小程序appid') VARCHAR(120)"`
	WxcodePath      string `json:"wxcode_path" xorm:"default '' comment('小程序路径') VARCHAR(255)"`
	WxcodeTitle     string `json:"wxcode_title" xorm:"default '' comment('小程序标题') VARCHAR(120)"`
}
