package models

type LotteryListInfo struct {
	Lid             int    `json:"lid" xorm:"not null pk comment('抽奖id') INT(11)"`
	ContactTitle    string `json:"contact_title" xorm:"default '' comment('联系标题') VARCHAR(255)"`
	ContactDesc     string `json:"contact_desc" xorm:"default '' comment('联系描述') VARCHAR(512)"`
	RecommendWxcode string `json:"recommend_wxcode" xorm:"comment('推荐小程序') TEXT"`
}
