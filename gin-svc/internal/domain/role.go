package domain

type Role struct {
	Id         int    `json:"id"`
	RoleName   string `json:"roleName"`
	RoleKey    string `json:"roleKey"`
	Sort       int    `json:"sort"`
	Status     bool   `json:"status"`
	CreateTime uint   `json:"createTime"`
}

type Permission struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Action string `json:"action"`
	PerKey string `json:"perKey"`
	Mark   string `json:"mark"`
	Path   string `json:"path"`
	Type   string `json:"-"`
	Status bool   `json:"status"`
}

type RoleDetail struct {
	RoleBase Role         `json:"roleBase"`
	PerList  []Permission `json:"perList"`
}
