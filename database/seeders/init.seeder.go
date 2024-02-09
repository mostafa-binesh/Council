package seeders

func InitSeeder() {
	RoleSeeder()
	PermissionSeeder()
	RoleHasPermissionSeeder()
	UserSeeder()
	LawSeeder()
	LawCommentsSeeder()
	AdminSeeder()
}
