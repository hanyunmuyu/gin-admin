package models

type Role struct {
	Model
	RoleName       string       `json:"roleName"`
	PermissionList []Permission `json:"permissionList" gorm:"many2many:role_permissions"`
}
