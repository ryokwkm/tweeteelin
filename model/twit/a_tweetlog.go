package twit

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ATweetLog struct {
	Id                   int    `gorm:"primary_key;column:id"`
	TweetId              string `gorm:"column:tweet_id"`
	Text                 string `gorm:"column:text"`
	OgJson               string `gorm:"column:og_json"`
	Retweet              int    `gorm:"column:retweet"`
	Favorite             int    `gorm:"column:favorite"`
	IsMedia              int    `gorm:"column:is_media"`
	IsRetweet            int    `gorm:"column:is_retweet"`
	EntityUserid         string `gorm:"column:entity_userid"`
	EntityUserscreenname string `gorm:"column:entity_userscreenname"`
	EntityUsername       string `gorm:"column:entity_username"`
	EntityRetweet        int    `gorm:"column:entity_retweet"`
	EntityFavorite       int    `gorm:"column:entity_favorite"`
	Created              string `gorm:"column:created" sql:"DEFAULT:current_timestamp"` //default:CURRENT_TIMESTAMP
}

type ATweetLogs []ATweetLog

func (t ATweetLog) TableName() string {
	tableName := "eelin"
	return "a_" + tableName + "_tweet_logs"
}

func StructATweetLogs(searchs []twitter.Tweet) ATweetLogs {
	ret := ATweetLogs{}
	for _, search := range searchs {
		tweetlog := ATweetLog{
			TweetId:  strconv.Itoa(int(search.ID)),
			Text:     search.Text,
			Retweet:  search.RetweetCount,
			Favorite: search.FavoriteCount,
		}
		if search.RetweetedStatus != nil {
			tweetlog.IsRetweet = 1
			tweetlog.EntityUserid = strconv.FormatInt(search.RetweetedStatus.User.ID, 10)
			tweetlog.EntityUsername = search.RetweetedStatus.User.Name
			tweetlog.EntityUserscreenname = search.RetweetedStatus.User.ScreenName
			tweetlog.EntityRetweet = search.RetweetedStatus.RetweetCount
			tweetlog.EntityFavorite = search.RetweetedStatus.FavoriteCount
		}
		ret = append(ret, tweetlog)
	}
	return ret
}

func (l ATweetLogs) InsertByStruct(db *gorm.DB) {
	for _, e := range l {
		db.Table(ATweetLog{}.TableName()).Create(&e)
	}
}
