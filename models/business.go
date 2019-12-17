package models

type Business struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Company      string `json:"company" xorm:"default '' comment('公司名称') VARCHAR(120)"`
	Name         string `json:"name" xorm:"default '' comment('姓名') VARCHAR(120)"`
	Weixin       string `json:"weixin" xorm:"default '' comment('微信号') VARCHAR(120)"`
	Tel          string `json:"tel" xorm:"default '' comment('手机号') CHAR(12)"`
	Introduction string `json:"introduction" xorm:"comment('产品简介') TEXT"`
	Uid          int    `json:"uid" xorm:"default 0 comment('提交用户uid') INT(11)"`
	AddTime      int    `json:"add_time" xorm:"default 0 comment('提交时间') INT(11)"`
	Status       int    `json:"status" xorm:"default 0 comment('是否处理 0没有 1处理 -1删除') TINYINT(1)"`
	Ip           string `json:"ip" xorm:"default '' comment('ip') CHAR(20)"`
	Post         string `json:"post" xorm:"default '' comment('职务') VARCHAR(120)"`
	CompanyType  int    `json:"company_type" xorm:"default 1 comment('公司类型  1.直客 2.代理商 3.平台商 4.公关公司 5.广告代理 6.其它类型') TINYINT(1)"`
	ProductName  string `json:"product_name" xorm:"default '' comment('产品名称') VARCHAR(120)"`
}
