package controllers

import (
    D "docker/database"
    M "docker/models"
)

func getPermissions(user M.User) [] string {
    // متغیر role_id را بر اساس فیلد RoleID از struct M.User تعریف می‌کنیم
    role_id := user.RoleID

    // جستجو برای اطلاعات سطح دسترسی‌های نقش بر اساس role_id
    var rolePermissions []*M.RoleHasPermission
    if err := D.DB().Where("role_id = ?", role_id).Find(&rolePermissions).Error; err != nil {
        // در صورت وجود خطا، آن را به عنوان خروجی تابع برمی‌گردانیم
        // شما می‌توانید برنامه خود را برای مدیریت خطاهای مختلف بهتر کنید
        // این یک نمونه ساده است
        return nil
    }

    // ایجاد یک slice برای ذخیره id سطح دسترسی‌ها
    var permissionsID []uint
    for _, rp := range rolePermissions {
        permissionsID = append(permissionsID, rp.PermissionID)
    }

    // جستجو برای اطلاعات سطح دسترسی‌ها بر اساس permissionsID
    var permissions []*M.Permission
    if err := D.DB().Where("id IN ?", permissionsID).Find(&permissions).Error; err != nil {
        // در صورت وجود خطا، آن را به عنوان خروجی تابع برمی‌گردانیم
        return nil
    }
	var permissions_id [] string
	for i := 0; i < len(permissions); i++ {
		permissions_id = append(permissions_id, permissions[i].RandomID)
	}
    // ارسال لیست permissions به عنوان خروجی تابع
    return permissions_id
}
    