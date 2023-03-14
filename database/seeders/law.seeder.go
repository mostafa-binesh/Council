package seeders

import (
	D "docker/database"
	M "docker/models"
	"time"
)

func LawSeeder() {
	// TODO add notifiation date later
	// ! statutes : اساس نامه ها
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساسنامه كانونهای فرهنگی و تربیتی دانشآموزان",
		SessionNumber:      503,
		SessionDate:        time.Date(1990, 2, 19, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/12035",
		NotificationDate:   time.Date(1990, 3, 14, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساسنامه مدارس ایثارگران (رزمندگان ، جانبازان و آزادگان)",
		SessionNumber:      514,
		SessionDate:        time.Date(1990, 9, 24, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/9000",
		NotificationDate:   time.Date(1990, 9, 30, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساسنامه مجتمع آموزشی تطبیقی",
		SessionNumber:      536,
		SessionDate:        time.Date(1992, 1, 30, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/11700",
		NotificationDate:   time.Date(1990, 2, 25, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اسانامه دبیرستان تربیت بدنی و علوم ورزشی(وابسته به آموزش و پرورش)",
		SessionNumber:      561,
		SessionDate:        time.Date(1993, 10, 7, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/11700",
		NotificationDate:   time.Date(1993, 11, 10, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	// ! regulations : آیین نامه ها
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه کانون زبان ایران",
		SessionNumber:      11,
		SessionDate:        time.Date(1989, 4, 12, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/4720",
		NotificationDate:   time.Date(1989, 8, 3, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه ارزشیابی درس تربیت بدنی  ",
		SessionNumber:      553,
		SessionDate:        time.Date(1993, 2, 25, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/11083",
		NotificationDate:   time.Date(1993, 3, 14, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه اعطای دیپلم افتخاری به جانبازان  ",
		SessionNumber:      114,
		SessionDate:        time.Date(1995, 2, 19, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/230",
		NotificationDate:   time.Date(1995, 4, 10, 0, 0, 0, 0, time.UTC),
		Body:               "<p>این مصوبه ی اول است</p><p>خط دوم</p>",
		Image:              "https://irimc.org/Portals/0/Images/News/621.jpg",
	})
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه  ارزشیابی مدارک تحصیلی مدارش کشور های خارجی   ",
		SessionNumber:      143,
		SessionDate:        time.Date(1997, 12, 3, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/8686",
		NotificationDate:   time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
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
