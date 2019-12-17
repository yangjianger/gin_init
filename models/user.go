package models

type User struct {
	Uid      int    `json:"uid" xorm:"not null pk autoincr INT(11)"`
	Username string `json:"username" xorm:"not null default '' comment('用户名') index VARCHAR(120)"`
	Tel      string `json:"tel" xorm:"not null default '' comment('手机号') index VARCHAR(12)"`
	Password string `json:"password" xorm:"not null default '' comment('密码') CHAR(32)"`
	Unionid  string `json:"unionid" xorm:"not null default '' comment('微信union_id') index VARCHAR(120)"`
	Avatar   string `json:"avatar" xorm:"not null default '' comment('头像') VARCHAR(255)"`
	Status   int    `json:"status" xorm:"default 0 comment('判断用户状态 0正常 -1 删除') TINYINT(1)"`
}
