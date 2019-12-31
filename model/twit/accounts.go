package twit

import "github.com/jinzhu/gorm"

type TwitterApp struct {
	Id             int    `gorm:"column:id"`
	AccountName    string `gorm:"column:account_name"`
	ConsumerKey    string `gorm:"column:consumerkey"`
	ConsumerSecret string `gorm:"column:consumersecret"`
	IsLogin        string `gorm:"column:is_login"`
	IsDeleted      string `gorm:"column:is_deleted"`
	created        string `gorm:"column:created"`
}

type TwitterUsers struct {
	Id            int    `gorm:"column:id"`
	AppId         int    `gorm:"column:app_id"`
	AccountName   string `gorm:"column:account_name"`
	UserId        string `gorm:"column:user_id"`
	AccessToken   string `gorm:"column:access_token"`
	AccessSecret  string `gorm:"column:access_secret"`
	LanguageId    int    `gorm:"column:language_id"`
	LocationId    int    `gorm:"column:location_id"`
	FunctionId    int    `gorm:"column:function_id"`
	Mode          int    `gorm:"column:mode"`
	FollowBack    int    `gorm:"column:followback"`
	FireLv        int    `gorm:"column:fire_lv"`
	SearchKeyword string `gorm:"column:search_keyword"`
	SearchOption  string `gorm:"column:search_option"`
	IsDeleted     int    `gorm:"column:is_deleted"`
	created       string `gorm:"column:created"`
}

//対象FUNCのユーザーリストを取得
func GetTwitterUserByFunc(db *gorm.DB, functionId int) (users []TwitterUsers) {
	if e := db.Find(&users, "function_id = ? and is_deleted = 0 and mode = 0", functionId).Error; e != nil {
		panic(e)
	}
	return
}

//対象modeのユーザーリストを取得
func GetTwitterUserByMode(db *gorm.DB, mode int) (users []TwitterUsers) {
	if e := db.Find(&users, " is_deleted = 0 and mode = ?", mode).Error; e != nil {
		panic(e)
	}
	return
}

//現在アクティブな全ユーザーを取得
func GetAllUser(db *gorm.DB) (users []TwitterUsers) {
	if e := db.Find(&users, "is_deleted = 0 and mode = 0").Error; e != nil {
		panic(e)
	}
	return
}

//デバッグユーザーを取得
func GetDebugUser(db *gorm.DB) (users []TwitterUsers) {
	if e := db.Find(&users, "is_deleted = 2").Error; e != nil {
		panic(e)
	}
	return
}

//対象ユーザーを取得
func GetTwitterUser(db *gorm.DB, appId int) (user TwitterUsers, error bool) {
	if e := db.Find(&user, "id = ? ", appId).Error; e != nil {
		error = false
	}
	return
}

func GetTwitterApps(db *gorm.DB, appId int) (app TwitterApp, error bool) {
	error = false
	if e := db.Find(&app, "id = ?", appId).Error; e != nil {
		error = true
	}
	return
}

//対象ユーザーのツイッター認証データを取得
func GetTwitterUserAuth(db *gorm.DB, userId int) (twitterAuth TwitterAuth, error bool) {
	user, error := GetTwitterUser(db, userId)
	if error {
		return
	}
	app, error := GetTwitterApps(db, user.AppId)
	if error {
		return
	}
	twitterAuth = TwitterAuth{
		ConsumerKey:    app.ConsumerKey,
		ConsumerSecret: app.ConsumerSecret,
		AccessToken:    user.AccessToken,
		AccessSecret:   user.AccessSecret,
	}
	return
}
