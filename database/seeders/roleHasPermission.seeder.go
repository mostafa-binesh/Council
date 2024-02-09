package seeders

import(
	D "docker/database"
	M "docker/models"
)
//! roles : 
// "admin",->1
// "super-user",->2
// 	"normally-user",->3

//! permissions :
// "view-admin",->1
// 	"view-user",->2
// 	"edit",->3
// 	"create",->4
// 	"delete",->5
func RoleHasPermissionSeeder(){
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 3,
		PermissionID: 1,
	})
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 2,
		PermissionID: 1,
	})
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 1,
		PermissionID: 1,
	})
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 1,
		PermissionID: 3,
	})
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 1,
		PermissionID: 4,
	})
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 1,
		PermissionID: 5,
	})
	D.DB().Create(&M.RoleHasPermission{
		RoleID: 1,
		PermissionID: 1,
	})
}