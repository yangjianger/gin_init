package models

type LotteryList struct {
	Id                 int    `json:"id" xorm:"not null pk autoincr comment('活动id') INT(11)"`
	Title              string `json:"title" xorm:"default '' comment('活动名称') VARCHAR(120)"`
	Cover              string `json:"cover" xorm:"default '' comment('活动封面图') VARCHAR(120)"`
	Num                int    `json:"num" xorm:"default 0 comment('奖品总数量') INT(11)"`
	ApplyNum           int    `json:"apply_num" xorm:"default 0 comment('抽奖人数') INT(11)"`
	Type               int    `json:"type" xorm:"default 1 comment('抽奖方式 1按时间自动开奖 2按人数自动开奖 3手动开奖') index TINYINT(1)"`
	IsLottery          int    `json:"is_lottery" xorm:"default 0 comment('是否开奖 0没有 1开奖') index TINYINT(1)"`
	LotteryTime        int    `json:"lottery_time" xorm:"default 0 comment('开奖时间') INT(11)"`
	LotteryPersonNum   int    `json:"lottery_person_num" xorm:"default 0 comment('开奖所需人数') INT(11)"`
	AddTime            int    `json:"add_time" xorm:"default 0 comment('添加时间') INT(11)"`
	UpdateTime         int    `json:"update_time" xorm:"default 0 comment('更新时间') INT(11)"`
	OnlineTime         int    `json:"online_time" xorm:"default 0 comment('上线时间') INT(11)"`
	EndTime            int    `json:"end_time" xorm:"default 0 comment('抽奖结束时间') INT(11)"`
	Uid                int    `json:"uid" xorm:"default 0 comment('发布用户id') index INT(11)"`
	Username           string `json:"username" xorm:"default '' comment('发布用户名') VARCHAR(32)"`
	UsernameAvatar     string `json:"username_avatar" xorm:"default '' comment('用户头像') VARCHAR(120)"`
	IsAdmin            int    `json:"is_admin" xorm:"default 0 comment('是否是后台发布 0前台 1后台') index TINYINT(1)"`
	Listorder          int    `json:"listorder" xorm:"default 0 comment('排序id') INT(7)"`
	IsHeight           int    `json:"is_height" xorm:"default 0 comment('是否是高级版 0不是 1是') TINYINT(1)"`
	Detail             string `json:"detail" xorm:"comment('抽奖说明') TEXT"`
	ApplyCondition     int    `json:"apply_condition" xorm:"default 0 comment('抽奖条件 0无 1.关注抽奖帝服务号 2.参与抽奖次数大于n次 3.发起抽奖次数大于n次') TINYINT(1)"`
	ApplyConditionNum  int    `json:"apply_condition_num" xorm:"default 0 comment('申请抽奖次数') INT(11)"`
	ApplyConditionTips string `json:"apply_condition_tips" xorm:"default '' comment('抽奖条件显示文案') VARCHAR(32)"`
	WriteIp            string `json:"write_ip" xorm:"default '' comment('发布ip') VARCHAR(20)"`
	IsSponsor          int    `json:"is_sponsor" xorm:"default 0 comment('是否有赞助商信息 0没有 1有') TINYINT(1)"`
	SponsorDetail      string `json:"sponsor_detail" xorm:"comment('赞助商介绍') TEXT"`
	SponsorName        string `json:"sponsor_name" xorm:"default '' comment('赞助商名称') VARCHAR(120)"`
	SponsorWxcodeUrl   string `json:"sponsor_wxcode_url" xorm:"default '' comment('赞助商小程序链接') VARCHAR(255)"`
	SponsorWxcodeName  string `json:"sponsor_wxcode_name" xorm:"default '' comment('商家小程序名称') VARCHAR(120)"`
	SponsorBtnName     string `json:"sponsor_btn_name" xorm:"default '' comment('赞助商按钮名称') VARCHAR(32)"`
	IsWxcodeUrl        int    `json:"is_wxcode_url" xorm:"default 0 comment('是否有小程序跳转 0没有 1跳转') TINYINT(1)"`
	Status             int    `json:"status" xorm:"default 0 comment('是否上线 0待审核 -1删除 -2下线 1上线  ') index TINYINT(1)"`
	IsTeam             int    `json:"is_team" xorm:"default 0 comment('是否可组队 0不能 1组队') TINYINT(1)"`
	TeamNum            int    `json:"team_num" xorm:"default 0 comment('组队所需人数') INT(5)"`
	UvNum              int    `json:"uv_num" xorm:"default 0 comment('活动uv') INT(11)"`
	SendStatus         int    `json:"send_status" xorm:"default 0 comment('发货状态 0进行中 1待发货 2已发货') TINYINT(1)"`
	IsSuccess          int    `json:"is_success" xorm:"default 0 comment('开奖是否成功 0不成功 1成功') TINYINT(1)"`
	SponsorWxcodeId    int    `json:"sponsor_wxcode_id" xorm:"default 0 comment('小程序对应表id') INT(11)"`
	IsWxcodeNotice     int    `json:"is_wxcode_notice" xorm:"default 0 comment('是否已经小程序通知填写收货信息') TINYINT(1)"`
	IsLotteryNotice    int    `json:"is_lottery_notice" xorm:"default 0 comment('是否发送开奖通知 0没有 1已通知') TINYINT(1)"`
	AdminType          int    `json:"admin_type" xorm:"default 1 comment('后台抽奖类型 1普通抽奖 2组队抽奖 3抽奖码') index TINYINT(4)"`
	Price              string `json:"price" xorm:"default 0.00 comment('商品总价') DECIMAL(11,2)"`
	PopularityAward    int    `json:"popularity_award" xorm:"default 2 comment('人气奖：1.有；2.没有；') TINYINT(1)"`
	UserTel            string `json:"user_tel" xorm:"default '' comment('用户手机号') VARCHAR(12)"`
	IsShare            int    `json:"is_share" xorm:"default 2 comment('是否开启分享  1开启分享 2没有') TINYINT(1)"`
	Imgs               string `json:"imgs" xorm:"comment('轮播图') TEXT"`
	IsHome             int    `json:"is_home" xorm:"default 0 comment('是否上首页 0不上 1审核中  2审核通过 3审核不通过') TINYINT(1)"`
	IsCopy             int    `json:"is_copy" xorm:"default 2 comment('开启一键复制 1开启 2不开启') TINYINT(1)"`
	IsDeposit          int    `json:"is_deposit" xorm:"default 1 comment('是否交押金1 没有 2已交 3已退') TINYINT(1)"`
	DepositType        int    `json:"deposit_type" xorm:"default 0 comment('首页申请押金类型 0、未知  1、99   2、499') TINYINT(1)"`
}
