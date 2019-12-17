package models

type LotteryTicketCode struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Mid        int    `json:"mid" xorm:"default 0 comment('玩法id') INT(11)"`
	TicketCode string `json:"ticket_code" xorm:"default '' comment('兑换码') VARCHAR(255)"`
	Type       int    `json:"type" xorm:"default 1 comment('是否可用 1可用 2不可用过') TINYINT(1)"`
	Lid        int    `json:"lid" xorm:"default 0 comment('活动id') INT(11)"`
}
