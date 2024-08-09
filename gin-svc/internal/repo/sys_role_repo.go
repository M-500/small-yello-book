package repo

import "gin-svc/internal/repo/dao"

type RoleRepo interface {
}

type roleRepoImpl struct {
	perDao dao.PermissionDAO
}
