package domain

import "time"

type DUser struct {
	GlobalNumber  string
	UserName      string
	NickName      string
	IPAddr        string
	Avatar        string
	Password      string
	Phone         string
	Signature     string
	FansCount     int
	FollowerCount int
	LikeCntCount  int
	Age           int
	Male          string
	BirthDay      time.Time
	Addr          string
	Profession    string
	School        string
	Level         int
	RoleList      []Role
	PerList       []Permission
}
