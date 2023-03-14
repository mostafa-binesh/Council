package seeders

import (
	D "docker/database"
	M "docker/models"
	"time"
)

func LawSeeder() {
	// TODO add notifiation date later
	// ! statutes : اساس نامه ها
	x := `<p dir="RTL" style="text-align:right; margin-bottom:11px"><strong>ماده1 -تعریف</strong><br />كانون فرهنگی و تربیتی دانشآموزان به مركزی اطلاق میشود كه برای اجرای برنامـه هـا و<br />فعالیتهای فرهنگی، ادبی، هنری، ورزشی، اجتماعی، علمی و مهارتی در جهت رشد و شكوفایی<br />استعدادهای نوجوانان و جوانان دانشآموز به ویژه دورههای راهنمـا یی تحصـیلی و متوسـطه در<br />سراسر كشور تشكیل میگردد<br /><strong>ماده 2 -هدف</strong><br />الف- بارور كردن اوقات فراغت دانشآموزان به منظور رشد و شكوفایی استعدادهای معنوی،<br />فكری، جسمی و عاطفی آنان از طریق اجرای برنامهها و فعالیـت هـای فرهنگـ ی، ادبـ ی، هنـر ی،<br />ورزشی، اجتماعی، علمی و مهارتی<br />ب- پرورش تقویت روحیه تعاون و اخوت دینـ ی و ایجـاد همبسـتگیی بـ ین دانـش آمـوزان و<br />آموختن روش زندگی اجتماعی به آنان بر اساس موازین اسلامی<br /><strong>فصل دوم- خط مشی، اصول برنامهها و شیوه تأسیس كانون فرهنگی و تربیتی<br />ماده 3</strong> -خط مشی و اصول برنامهریزی كانونهای فرهنگی و تربیتـ ی توسـط معـاون پرورشـ ی<br />پیشنهاد میشود و پس از تأیید وزیر آموزشوپرورش اجرا میگردد<br /><strong>ماده 4 </strong>-تأسیسكانون فرهنگی و تربیتی در هر شهر، شهرستان، منطقه و ناحیه بنا به پیشنهاد<br />مدیركل استان مربوط و موافقـت معـاون پرورشـی وزارت آمـوزش وپـرورش در قالـب<br />برنامهها و اعتبارات مصوب خواهد بود</p>`
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساسنامه كانونهای فرهنگی و تربیتی دانشآموزان",
		SessionNumber:      503,
		SessionDate:        time.Date(1990, 2, 19, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/12035",
		NotificationDate:   time.Date(1990, 3, 14, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساسنامه مدارس ایثارگران (رزمندگان ، جانبازان و آزادگان)",
		SessionNumber:      514,
		SessionDate:        time.Date(1990, 9, 24, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/9000",
		NotificationDate:   time.Date(1990, 9, 30, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اساسنامه مجتمع آموزشی تطبیقی",
		SessionNumber:      536,
		SessionDate:        time.Date(1992, 1, 30, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/11700",
		NotificationDate:   time.Date(1990, 2, 25, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               2,
		Title:              "اسانامه دبیرستان تربیت بدنی و علوم ورزشی(وابسته به آموزش و پرورش)",
		SessionNumber:      561,
		SessionDate:        time.Date(1993, 10, 7, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/11700",
		NotificationDate:   time.Date(1993, 11, 10, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	// ! regulations : آیین نامه ها
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه کانون زبان ایران",
		SessionNumber:      11,
		SessionDate:        time.Date(1989, 4, 12, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/4720",
		NotificationDate:   time.Date(1989, 8, 3, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه ارزشیابی درس تربیت بدنی  ",
		SessionNumber:      553,
		SessionDate:        time.Date(1993, 2, 25, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/11083",
		NotificationDate:   time.Date(1993, 3, 14, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه اعطای دیپلم افتخاری به جانبازان  ",
		SessionNumber:      114,
		SessionDate:        time.Date(1995, 2, 19, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/230",
		NotificationDate:   time.Date(1995, 4, 10, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               1,
		Title:              "آیین نامه  ارزشیابی مدارک تحصیلی مدارش کشور های خارجی   ",
		SessionNumber:      143,
		SessionDate:        time.Date(1997, 12, 3, 0, 0, 0, 0, time.UTC),
		NotificationNumber: "120/8686",
		NotificationDate:   time.Date(1997, 12, 28, 0, 0, 0, 0, time.UTC),
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})

	// ! enactments : مصوبه ها
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "ماده واحده تشکیل ستاد هماهنگی و پشتیبانی اجرای سند تحول",
		SessionNumber:      4515874,
		NotificationNumber: "581/5841",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "نظام جامع راهنمایی و مشاوره در آموزش و پرورش",
		SessionNumber:      4514174,
		NotificationNumber: "5522/9542",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "مجموعه مصوبات شورای عالی آموزش و پرورش",
		SessionNumber:      13265,
		NotificationNumber: "581/1238",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "آیین نامه تاسیس مدارس اتباع خارجی",
		SessionNumber:      13265,
		NotificationNumber: "581/1238",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "نحوه جبران عقب ماندگی تحصیلی دانش آموزان در دوران کرونا",
		SessionNumber:      13265,
		NotificationNumber: "581/1238",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "نواختن زنگ زبان و ادبیات فارسی در مدارس",
		SessionNumber:      13265,
		NotificationNumber: "581/1238",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})
	D.DB().Create(&M.Law{
		Type:               3,
		Title:              "ماده واحده تشکیل ستاد هماهنگی و پشتیبانی اجرای سند تحول",
		SessionNumber:      13265,
		NotificationNumber: "581/1238",
		Body:               x,
		Image:              "https://s2.uupload.ir/files/placeholder-image_ux76.png",
	})

}
