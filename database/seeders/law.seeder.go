package seeders

import (
	D "docker/database"
	M "docker/models"
)

func LawSeeder() {

	// ! regulations : آیین نامه ها
	// TODO add notifiation date later
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه ی اول",
		SessionNumber:      5642,
		NotificationNumber: "1251/21621",
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	// ! statutes : اساس نامه ها
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساس نامه ی اول",
		SessionNumber:      54612,
		NotificationNumber: "41/531",
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	// ! enactments : مصوبه ها
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "مصوبه ی اول",
		SessionNumber:      4515874,
		NotificationNumber: "581/5841",
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})

}
