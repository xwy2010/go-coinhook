package templates

import (
	"time"
)

type Admin struct {
	Uid      int       `json:"uid" xorm:"not null pk autoincr unique INT(11)"`
	Username string    `json:"username" xorm:"not null unique VARCHAR(45)"`
	Passwd   string    `json:"passwd" xorm:"not null VARCHAR(45)"`
	Mobile   int       `json:"mobile" xorm:"not null unique INT(11)"`
	Avatar   string    `json:"avatar" xorm:"VARCHAR(45)"`
	Ctime    time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Ltime    time.Time `json:"ltime" xorm:"not null default CURRENT_TIMESTAMP comment('最后登录时间') TIMESTAMP"`
}

type Admins struct {
	Id                 int    `json:"id" xorm:"not null pk autoincr INT(16)"`
	Username           string `json:"username" xorm:"not null comment('登录名') VARCHAR(20)"`
	Nickname           string `json:"nickname" xorm:"comment('显示的名字') VARCHAR(20)"`
	Password           string `json:"password" xorm:"not null VARCHAR(256)"`
	LastLoginTimestamp int    `json:"last_login_timestamp" xorm:"INT(16)"`
}
