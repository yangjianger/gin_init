package models

type UserInfo struct {
	Uid               int    `json:"uid" xorm:"not null pk INT(11)"`
	AddressUsername   string `json:"address_username" xorm:"default '' VARCHAR(120)"`
	Username          string `json:"username" xorm:"default '' comment('用户名') VARCHAR(120)"`
	Avatar            string `json:"avatar" xorm:"default '' comment('用户头像') VARCHAR(255)"`
	WebOpenid         string `json:"web_openid" xorm:"default '' comment('服务号openid') index VARCHAR(120)"`
	WxcodeOpenid      string `json:"wxcode_openid" xorm:"default '' comment('小程序openid') index VARCHAR(120)"`
	Tel               string `json:"tel" xorm:"default '' comment('手机号') VARCHAR(12)"`
	Sex               int    `json:"sex" xorm:"default 1 comment('性别 0未知 1男2女') TINYINT(1)"`
	Birthday          int    `json:"birthday" xorm:"default 0 comment('生日') INT(11)"`
	Age               int    `json:"age" xorm:"default 0 comment('年龄') INT(3)"`
	GroupId           int    `json:"group_id" xorm:"default 1 comment('分组.1.前台；2.后台') INT(1)"`
	RegDate           int    `json:"reg_date" xorm:"default 0 comment('注册时间') INT(11)"`
	RegIp             string `json:"reg_ip" xorm:"default '' comment('注册ip') VARCHAR(20)"`
	LastLoginDate     int    `json:"last_login_date" xorm:"default 0 comment('最后一次登录时间') INT(11)"`
	LastLoginIp       string `json:"last_login_ip" xorm:"default '' comment('最后登录ip') VARCHAR(20)"`
	Intro             string `json:"intro" xorm:"default '' comment('用户介绍') VARCHAR(255)"`
	IsComplete        int    `json:"is_complete" xorm:"default 0 comment('信息是否完善 0没有 1已完善') TINYINT(1)"`
	Status            int    `json:"status" xorm:"default 1 comment('用户状态 1正常状态 2 删除至回收站 3锁定') index TINYINT(1)"`
	Address           string `json:"address" xorm:"comment('用户地址') TEXT"`
	City              string `json:"city" xorm:"default '' comment('用户城市') VARCHAR(120)"`
	AddressRemark     string `json:"address_remark" xorm:"default '' comment('用户地址备注') VARCHAR(255)"`
	JoinLotteryNum    int    `json:"join_lottery_num" xorm:"default 0 comment('参与抽奖数') INT(11)"`
	InitLotteryNum    int    `json:"init_lottery_num" xorm:"default 0 comment('发起抽奖数') INT(11)"`
	SuccessLotteryNum int    `json:"success_lottery_num" xorm:"default 0 comment('中奖成功数') INT(11)"`
	IsBan             int    `json:"is_ban" xorm:"default 0 comment('是否被封禁 0没有 1封禁') TINYINT(1)"`
	BanTime           int    `json:"ban_time" xorm:"default 0 comment('封禁时间') INT(11)"`
	IsForeverBan      int    `json:"is_forever_ban" xorm:"default 0 comment('是否永久封禁 0不是 1是') TINYINT(1)"`
	IsPush            int    `json:"is_push" xorm:"default 1 comment('是否可以推送 0不能 1可以') index TINYINT(1)"`
	WxAvatar          string `json:"wx_avatar" xorm:"default '' comment('微信源头像') VARCHAR(255)"`
}
