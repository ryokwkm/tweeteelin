package twit

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwitterUserFollower struct {
	UserId   int    `gorm:"primary_key;column:user_id"`
	Created  string `gorm:"primary_key;column:created"`
	Follower int    `gorm:"column:follower"`
}

func UserFollowerSave(db *gorm.DB, user_id int, follower int) {
	userFollower := TwitterUserFollower{
		UserId:   user_id,
		Created:  time.Now().Format("2006-01-02"),
		Follower: follower,
	}
	db.Delete(&userFollower)
	db.Create(&userFollower)
}
