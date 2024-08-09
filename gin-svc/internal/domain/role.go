package domain

import "time"

type Role struct {
	Id         int        `json:"id"`
	RoleName   string     `json:"roleName"`
	RoleKey    string     `json:"roleKey"`
	Sort       int        `json:"sort"`
	Status     bool       `json:"status"`
	CreateTime *time.Time `json:"createTime"`
}
