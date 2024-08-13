package service

import (
	"context"
	"errors"
	"gin-svc/internal/domain"
	"gin-svc/internal/models"
	"gin-svc/internal/repo"
	"gin-svc/internal/types"
	"gorm.io/gorm"
)

type RoleSvc interface {
	// 获取角色详情，通过角色ID 后台管理用  {RoleBase, PerList}
	GetDetailByID(ctx context.Context, id int) (domain.RoleDetail, error)
	// 创建角色信息
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
	err := r.roleRepo.DeleteRole(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *roleSvcImpl) UpdateRole(ctx context.Context, role types.UpdateRoleReq) error {
	perList := role.PerIds
	err := r.roleRepo.CheckPermissionIds(ctx, perList)
	if err != nil {
		return errors.New("权限列表不合法！")
	}
	base := models.SysRoleModel{
		Model: gorm.Model{
			ID: uint(role.ID),
		},
		//CreatorId: 0,
		RoleName: role.RoleName,
		Status:   role.Status,
		RoleSort: role.Sort,
		//Remark:    role.,
	}
	return r.roleRepo.UpdateRole(ctx, &base, role.PerIds)
}
