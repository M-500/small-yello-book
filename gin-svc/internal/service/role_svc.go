package service

import (
	"context"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/types"
)

type RoleSvc interface {
	GetDetailByID(ctx context.Context, id int) (domain.RoleDetail, error)
	CreateRole(ctx context.Context, role types.CreateRoleReq) error
	DeleteRole(ctx context.Context, id int) error
	UpdateRole(ctx context.Context, role types.UpdateRoleReq) error
}

func NewRoleService(roleRepo repo.RoleRepo) RoleSvc {
	return &roleSvcImpl{
		roleRepo: roleRepo,
	}
}

type roleSvcImpl struct {
	roleRepo repo.RoleRepo
}

func (r *roleSvcImpl) toDMRole(data models.SysRoleModel) domain.Role {
	return domain.Role{
		Id:       int(data.ID),
		RoleName: data.RoleName,
		RoleKey:  data.RoleKey,
		//Sort:       int(data.Sort),
		Status:     data.Status,
		CreateTime: &data.CreatedAt,
	}
}
func (r *roleSvcImpl) toDMPermission(data models.SysPermissionModel) domain.Permission {
	return domain.Permission{
		Id:     int(data.ID),
		Title:  data.Title,
		Action: data.Action,
		PerKey: data.PerKey,
		Mark:   data.Mark,
		Path:   data.Path,
		Type:   data.Type,
		Status: data.Status,
	}
}

func (r *roleSvcImpl) toRoleModel(data types.CreateRoleReq) models.SysRoleModel {
	return models.SysRoleModel{
		//Model:     gorm.Model{},
		//CreatorId: 0,
		RoleName: data.RoleName,
		RoleKey:  data.RoleKey,
		Status:   data.Status,
		RoleSort: data.Sort,
		//Remark:    data.,
	}
}

func (r *roleSvcImpl) GetDetailByID(ctx context.Context, id int) (domain.RoleDetail, error) {
	var res domain.RoleDetail
	byID, err := r.roleRepo.GetRoleByID(ctx, id)
	if err != nil {
		return domain.RoleDetail{}, err
	}
	res.RoleBase = r.toDMRole(*byID)
	pers, err := r.roleRepo.FindPermissionListByRoleId(ctx, id)
	if err != nil {
		return res, err
	}
	for _, item := range pers {
		res.PerList = append(res.PerList, r.toDMPermission(item))
	}
	return res, nil
}

func (r *roleSvcImpl) CreateRole(ctx context.Context, req types.CreateRoleReq) error {

	temp := r.toRoleModel(req)
	return r.roleRepo.CreateRole(ctx, &temp)
}

func (r *roleSvcImpl) DeleteRole(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *roleSvcImpl) UpdateRole(ctx context.Context, role types.UpdateRoleReq) error {
	//TODO implement me
	panic("implement me")
}
