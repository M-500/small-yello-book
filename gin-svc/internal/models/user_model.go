package models

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	gorm.Model
	GlobalNumber  string     `json:"global_number" gorm:"column:global_number;type:varchar(64);comment:全局编号"`
	UserName      string     `json:"user_name" gorm:"column:user_name;type:varchar(128);comment:用户名,可以修改,默认生成"`
	Email         string     `json:"email" gorm:"unique;column:email;type:varchar(128);comment:邮箱"`
	NickName      string     `json:"nick_name" gorm:"column:nick_name;type:varchar(128);comment:昵称"`
	IPAddr        string     `json:"ip_addr" gorm:"column:ip_addr;type:varchar(100);comment:IP归属地"`
	Avatar        string     `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像"`
	Password      string     `json:"password" gorm:"column:password;type:varchar(128);comment:密码"`
	Signature     string     `json:"signature" gorm:"column:signature;type:varchar(255);comment:个性签名"`
	FansCount     int        `json:"fans_count" gorm:"column:fans_count;type:int;comment:粉丝数"`
	FollowerCount int        `json:"follower_count" gorm:"column:follower_count;type:int;comment:关注数"`
	LikeCntCount  int        `json:"like_cnt_count" gorm:"column:like_cnt_count;type:int;comment:点赞数"`
	Age           int        `json:"age" gorm:"column:age;type:int;comment:年龄"`
	Male          string     `json:"male" gorm:"column:male;type:varchar(18);comment:性别,男/女"`
	BirthDay      *time.Time `json:"birth_day" gorm:"column:birth_day;type:datetime;comment:生日"`
	Addr          string     `json:"addr" gorm:"column:addr;type:varchar(255);comment:地址"`
	Profession    string     `json:"profession" gorm:"column:profession;type:varchar(128);comment:职业"`
	School        string     `json:"school" gorm:"column:school;type:varchar(200);comment:学校"`
	Level         int        `json:"level" gorm:"column:level;type:int;comment:等级"`
}

func (UserModel) TableName() string {
	return "sys_user"
}
