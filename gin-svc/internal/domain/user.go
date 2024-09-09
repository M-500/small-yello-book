package domain

import (
	"gin-svc/internal/models"
	"time"
)

type DUser struct {
	GlobalNumber  string       `json:"globalNumber"`
	UserName      string       `json:"userName"`
	NickName      string       `json:"nickName"`
	IPAddr        string       `json:"IPAddr"`
	Avatar        string       `json:"avatar"`
	Password      string       `json:"password"`
	Email         string       `json:"email"`
	Signature     string       `json:"signature"`
	FansCount     int          `json:"fansCount"`
	FollowerCount int          `json:"followerCount"`
	LikeCntCount  int          `json:"likeCntCount"`
	Age           int          `json:"age"`
	Male          string       `json:"male"`
	BirthDay      time.Time    `json:"birthDay"`
	Addr          string       `json:"addr"`
	Profession    string       `json:"profession"`
	School        string       `json:"school"`
	Level         int          `json:"level"`
	RoleList      []Role       `json:"roleList"`
	PerList       []Permission `json:"perList"`
}

func (d *DUser) ToDomainPermission(res []models.SysPermissionModel) {
	for _, re := range res {
		d.PerList = append(d.PerList, Permission{
			Id:     int(re.ID),
			PerKey: re.PerKey,
			Title:  re.Title,
			Action: re.Action,
			Path:   re.Path,
			Mark:   re.Mark,
		})
	}
}
