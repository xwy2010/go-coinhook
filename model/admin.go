package model

import (
	"hrforceAdmin/db"
	"time"
)

// Admin 移动端后台管理员表
type Admin struct {
	UID      int       `json:"uid" xorm:"not null pk autoincr unique INT(11) 'uid'"`
	Username string    `json:"username" xorm:"not null unique VARCHAR(45)"`
	Password string    `json:"password" xorm:"not null VARCHAR(45)"`
	Mobile   string    `json:"mobile" xorm:"not null unique VARCHAR(20)"`
	Avatar   string    `json:"avatar" xorm:"VARCHAR(45)"`
	Ctime    time.Time `json:"ctime" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Ltime    time.Time `json:"ltime" xorm:"not null default CURRENT_TIMESTAMP comment('最后登录时间') TIMESTAMP"`
}

// GetByName 根据用户名查询
func (a *Admin) GetByName(username string) (user *Admin, has bool, err error) {
	user = &Admin{Username: username}
	has, err = db.MasterDB.Get(user)

	if !has {
		return nil, has, err
	}
	return user, has, err
}

// GetByMobile 根据手机号查询
func (a *Admin) GetByMobile(mobile string) (admin *Admin, err error) {
	admin = &Admin{Mobile: mobile}
	has, err := db.MasterDB.Get(admin)

	if !has {
		return nil, err
	}
	return admin, err
}
