package model

import (
	"hrforceAdmin/db"
)

// Companys B端企业用户表
type Companys struct {
	ID          int    `json:"id" xorm:"not null pk  INT(16) 'id'"`
	CompanyCode string `json:"company_code" xorm:"not null unique VARCHAR(20)" validate:"required"`
	CompanyName string `json:"company_name" xorm:"index VARCHAR(70)" validate:"required"`
	LogoUri     string `json:"logo_uri" xorm:"comment('logo的地址') VARCHAR(160)"`
	// LoginName   string `json:"login_name" xorm:"index VARCHAR(20)"`
	// LoginMail   string `json:"login_mail" xorm:"comment('公司登录名(使用邮箱)') index VARCHAR(50)"`
	Password string `json:"password" xorm:"VARCHAR(256)" validate:"required"`
	// Industry             int    `json:"industry" xorm:"comment('企业所属的行业代码') INT(7)"`
	// CompanyType          int    `json:"company_type" xorm:"comment('企业性质。    '1' => '国营企业',     '2' => '私营企业',     '3' => '外资企业',     '4' => '合资企业',     '5' => '事业单位',     '6' => '政府机关',     '7' => '其他',') TINYINT(3)"`
	// EmployeeCount        int    `json:"employee_count" xorm:"comment('企业员工数.1 = 50人以下，2=51-100, 3= 101-150, 4= 151-500, 5=500以上') TINYINT(3)"`
	ContactName string `json:"contact_name" xorm:"comment('联系人名字') VARCHAR(20)"`
	// Address         string `json:"address" xorm:"comment('公司地址') VARCHAR(120)"`
	CompanyPhonenum string `json:"company_phonenum" xorm:"comment('公司电话') VARCHAR(20)"`
	ContactPhonenum string `json:"contact_phonenum" xorm:"comment('联系人的电话') VARCHAR(20)"`
	// LicenceImgUri        string `json:"licence_img_uri" xorm:"comment('营业执照的照片uri') VARCHAR(200)"`
	// Description          string `json:"description" xorm:"VARCHAR(10000)"`
	LastLoginTimestamp int `json:"last_login_timestamp" xorm:"comment('最后一次登录的时间戳') INT(11)"`
	Balance            int `json:"balance" xorm:"not null comment('余额') INT(16)" validate:"gte=0"`
	Status             int `json:"status" xorm:"not null comment('0 = 等待激活，1= 正常 2 = 被限制,3=未完成填写资料 ') TINYINT(1)"`
	// PaperCount           int    `json:"paper_count" xorm:"not null default 0 comment('该企业已创建的问卷数，包括已完成和未完成的') INT(16)"`
	RegisterTimestamp int `json:"register_timestamp" xorm:"created INT(16)"`
	// ValidateTimestamp    int    `json:"validate_timestamp" xorm:"INT(16)"`
	// UseSmtpDefault int `json:"use_smtp_default" xorm:"default 1 TINYINT(1)"`
	// SmtpServer           string `json:"smtp_server" xorm:"VARCHAR(64)"`
	// SmtpPort             string `json:"smtp_port" xorm:"VARCHAR(6)"`
	// SmtpUsername         string `json:"smtp_username" xorm:"VARCHAR(64)"`
	// SmtpPassword         string `json:"smtp_password" xorm:"VARCHAR(64)"`
	// Background           string `json:"background" xorm:"VARCHAR(255)"`
	// Website              string `json:"website" xorm:"VARCHAR(255)"`
	// CompanyAbbreviations string `json:"company_abbreviations" xorm:"VARCHAR(20)"`
	// Financing            int    `json:"financing" xorm:"not null default 0 TINYINT(4)"`
	// Legalperson          string `json:"legalperson" xorm:"VARCHAR(40)"`
	// CompanyBussiness     string `json:"company_bussiness" xorm:"VARCHAR(100)"`
	// OrganizationCode     string `json:"organization_code" xorm:"VARCHAR(24)"`
	// Agent                string `json:"agent" xorm:"VARCHAR(16)"`
	// AgentPhone           string `json:"agent_phone" xorm:"VARCHAR(20)"`
	// Attorney             string `json:"attorney" xorm:"VARCHAR(64)"`
	// IsRealname           int    `json:"is_realname" xorm:"not null default 0 INT(4)"`
	// HaveReward           int    `json:"have_reward" xorm:"not null default 0 TINYINT(2)"`
	// AuthorizationImgUri  string `json:"authorization_img_uri" xorm:"VARCHAR(200)"`
	// TmpCompanyName       string `json:"tmp_company_name" xorm:"VARCHAR(70)"`
	// InviteCode           string `json:"invite_code" xorm:"not null VARCHAR(10)"`
	// InviteReward         int    `json:"invite_reward" xorm:"not null default 0 INT(2)"`
	// SubscibeMail int `json:"subscibe_mail" xorm:"default 1 TINYINT(4)"`
	// MajorLink    string `json:"major_link" xorm:"VARCHAR(255) default 'https://www.zhishidashi.com/fission/front/product_details/33?rowKey=79808869858497659_posterFission_ecd221dd4e5c4cdebc24e184a4423ca0&fission=751ef22a0257e1ebbc678363ada394cf'"`
}

// TCodeCompanysize ??
type TCodeCompanysize struct {
	ID     int    `json:"id" xorm:"not null pk default 0 INT(11)"`
	Name   string `json:"name" xorm:"comment('行业名称') VARCHAR(50)"`
	Index  int    `json:"index" xorm:"comment('序号') INT(11)"`
	Status int    `json:"status" xorm:"comment('状态') TINYINT(2)"`
}

// GetByComName 根据公司名模糊查询
func (c *Companys) GetByComName(page int) (users []*Companys, err error) {
	pi := 10 * page
	users = make([]*Companys, 0)
	err = db.MasterDB.Where("company_name like ?", "%"+c.CompanyName+"%").Limit(10, pi).Find(&users)

	if err != nil {
		return nil, err
	}
	return
}

// Get 根据前端提交的参数查询
func (c *Companys) Get() (users []*Companys, err error) {
	users = make([]*Companys, 0)
	err = db.MasterDB.Find(&users, c)
	if err != nil {
		return nil, err
	}
	return
}

// Count 企业用户统计
func (c *Companys) Count() (total int64, err error) {
	// com := new(Companys)
	total, err = db.MasterDB.Where("id >?", 1).Count(c)
	return
}

// UpdateStatus 禁用/激活账号
func (c *Companys) UpdateStatus() (affected int64, err error) {
	affected, err = db.MasterDB.ID(c.ID).Cols("status").Update(c)
	return
}

// UpdatebyStruct 根据传入的map数据多字段更新
func (c *Companys) UpdatebyStruct() (affected int64, err error) {
	affected, err = db.MasterDB.ID(c.ID).Update(c)
	return
}

// UpdateBalance 修改Balance为0
func (c *Companys) UpdateBalance() (affected int64, err error) {

	affected, err = db.MasterDB.ID(c.ID).Cols("balance").Update(c) //
	return
}

// DelByID 根据ID删除用户
func (c *Companys) DelByID() (affected int64, err error) {
	// com := new(Companys)
	affected, err = db.MasterDB.ID(c.ID).Delete(c)
	return
}

// Insert  单用户插入
func (c *Companys) Insert() (affected int64, err error) {
	affected, err = db.MasterDB.Insert(c)
	return
}
