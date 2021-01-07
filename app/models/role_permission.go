package models

type RolePermission struct {
	Model
	RoleId       uint `json:"roleId"`
	PermissionId uint `json:"permissionId"`
}
