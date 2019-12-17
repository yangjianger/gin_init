package models

type LotteryListMeta struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Lid            int    `json:"lid" xorm:"default 0 comment('抽奖id') index INT(11)"`
	Title          string `json:"title" xorm:"default '' comment('奖品名称') VARCHAR(120)"`
	Num            int    `json:"num" xorm:"default 0 comment('奖品数量') INT(11)"`
	Price          string `json:"price" xorm:"comment('奖品单价') DECIMAL(10,2)"`
	MetaType       int    `json:"meta_type" xorm:"default 1 comment('奖项类别：1.正常奖项；2.人气奖；3.阳光普照奖') index TINYINT(1)"`
	LotteryType    int    `json:"lottery_type" xorm:"default 1 comment('奖品类型 1实物  2小程序优惠券  3购买折扣卷') TINYINT(1)"`
	WxcodeTitle    string `json:"wxcode_title" xorm:"default '' comment('小程序标题') VARCHAR(120)"`
	WxcodeAppid    string `json:"wxcode_appid" xorm:"default '' comment('小程序appid') VARCHAR(120)"`
	WxcodePath     string `json:"wxcode_path" xorm:"default '' comment('小程序路径') VARCHAR(255)"`
	WxcodeId       int    `json:"wxcode_id" xorm:"default 0 comment('对应小程序表id') INT(11)"`
	TicketCodeDesc string `json:"ticket_code_desc" xorm:"default '' comment('兑换码描述') VARCHAR(512)"`
	Cover          string `json:"cover" xorm:"default '' comment('奖品封面') VARCHAR(120)"`
}
