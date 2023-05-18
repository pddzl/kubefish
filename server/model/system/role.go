package system

import "github.com/pddzl/kubefish/server/global"

type RoleModel struct {
	global.KF_MODEL
	RoleName string `json:"roleName" gorm:"unique"`
	//Users    []*UserModel `json:"users"`
	Menus []*MenuModel `json:"menus" gorm:"many2many:role_menus;"`
}

func (RoleModel) TableName() string {
	return "sys_role"
}
