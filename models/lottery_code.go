package models

type LotteryCode struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Lid        int    `json:"lid" xorm:"comment('活动id') index INT(11)"`
	Uid        int    `json:"uid" xorm:"comment('抽奖用户uid') index INT(11)"`
	Username   string `json:"username" xorm:"default '' comment('抽奖用户名') VARCHAR(52)"`
	UserAvatar string `json:"user_avatar" xorm:"default '' comment('抽奖用户头像') VARCHAR(255)"`
	TopId      int    `json:"top_id" xorm:"default 0 comment('拉人用户申请id') index INT(11)"`
	AddTime    int    `json:"add_time" xorm:"default 0 comment('添加时间') INT(11)"`
	Code       string `json:"code" xorm:"default '' comment('抽奖码') index VARCHAR(32)"`
	Status     int    `json:"status" xorm:"default 0 comment('是否中奖 0没有 1中奖') TINYINT(1)"`
	IsUserSelf int    `json:"is_user_self" xorm:"default 0 comment('是否是用户本人 0不是 1是') index TINYINT(1)"`
	IsNotice   int    `json:"is_notice" xorm:"default 0 comment('是否邀请通知0没有 1通知') TINYINT(1)"`
	IsReward   int    `json:"is_reward" xorm:"default 2 comment('是否是奖励获得 1是 2不是') TINYINT(1)"`
}
