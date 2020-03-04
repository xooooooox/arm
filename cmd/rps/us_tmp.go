package main

// datetime 2020-03-04 38:42:20

// UsUserBindUsUserWxMp 会员绑定微信公众号会员表
type UsUserBindUsUserWxMp struct {
	UsId         int64 // id
	UsUserId     int64 // 会员id
	UsUserWxMpId int64 // 微信公众号会员id
	UsTime       int64 // 时间
}

// UsUser 会员表
type UsUser struct {
	UsId       int64   // id
	UsParent   int64   // 上级
	UsName     string  // 姓名
	UsMobile   string  // 手机号
	UsEmail    string  // 电子邮件
	UsAge      int     // 年龄
	UsBalance  float64 // 余额
	UsIntegral float64 // 积分
	UsTime     int64   // 时间
	UsStatus   int     // 状态
	UsNote     string  // 备注
}

// UsUserWxMp 微信公众号会员表
type UsUserWxMp struct {
	UsId         int64  // id
	UsParent     int64  // 上级
	UsMp         int64  // 微信公众号
	UsOpenid     string // openid
	UsNickname   string // 昵称
	UsSex        int    // 性别
	UsCountry    string // 国家
	UsProvince   string // 省
	UsCity       string // 城市
	UsHeadImgUrl string // 头像
	UsPrivilege  string // 特权
	UsUnionId    string // 联合id
	UsTime       int64  // 时间
	UsNote       string // 备注
}

// UsWxMp 微信公众号表
type UsWxMp struct {
	UsId        int64  // id
	UsName      string // 名称
	UsAppId     string // 微信公众号id
	UsAppSecret string // 微信公众号秘钥
	UsTime      int64  // 时间
	UsStatus    int    // 状态
	UsNote      string // 备注
}
