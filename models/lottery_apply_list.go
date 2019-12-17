package models

type LotteryApplyList struct {
	Id                int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Lid               int    `json:"lid" xorm:"default 0 comment('活动id') index unique(lid,uid) INT(11)"`
	Uid               int    `json:"uid" xorm:"default 0 comment('用户id') unique(lid,uid) index INT(11)"`
	Username          string `json:"username" xorm:"default '' comment('申请用户名') VARCHAR(120)"`
	UserAvatar        string `json:"user_avatar" xorm:"default '' comment('申请用户头像') VARCHAR(255)"`
	AddTime           int    `json:"add_time" xorm:"default 0 comment('申请时间') INT(11)"`
	ApplyIp           string `json:"apply_ip" xorm:"default '' comment('申请ip') VARCHAR(20)"`
	TopId             int    `json:"top_id" xorm:"default 0 comment('发起拼团ip') index INT(11)"`
	TopNum            int    `json:"top_num" xorm:"default 1 comment('当前该队伍的人数') INT(11)"`
	Status            int    `json:"status" xorm:"default 0 comment('是否中奖 0没有 1中奖 -1删除') index TINYINT(1)"`
	Mid               int    `json:"mid" xorm:"default 0 comment('中奖玩法id') INT(11)"`
	MetaTitle         string `json:"meta_title" xorm:"default '' comment('中奖产品名') VARCHAR(32)"`
	IsLeader          int    `json:"is_leader" xorm:"default 0 comment('是否是队长 0不是 1是') TINYINT(1)"`
	Orderid           string `json:"orderid" xorm:"default '' comment('订单id') VARCHAR(32)"`
	IsTeam            int    `json:"is_team" xorm:"default 0 comment('是否开启组队 0不开启 1开启') TINYINT(1)"`
	IsWxcodeNotice    int    `json:"is_wxcode_notice" xorm:"default 0 comment('是否已经通知') index TINYINT(1)"`
	IsPushNews        int    `json:"is_push_news" xorm:"default 0 comment('是否已经通知 0默认 1待发送  2已发送未推送  3推送成功 -1推送失败') TINYINT(1)"`
	InviteNum         int    `json:"invite_num" xorm:"default 1 comment('抽奖码获得数量') index INT(11)"`
	IsAdmin           int    `json:"is_admin" xorm:"default 0 comment('是否是后台活动') TINYINT(1)"`
	IsPopularityAward int    `json:"is_popularity_award" xorm:"default 2 comment('是否是人气奖 1是 2不是') TINYINT(1)"`
	LotteryType       int    `json:"lottery_type" xorm:"default 1 comment('奖品类型 1实物  2小程序优惠券  3购买折扣卷') TINYINT(1)"`
}
