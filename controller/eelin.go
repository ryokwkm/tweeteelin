package main

import (
	"github.com/jinzhu/gorm"
	"github.com/ryokwkm/trends-model/twit"
	"github.com/sirupsen/logrus"
)

type EelinUsecase struct {
	Db          *gorm.DB
	TwitterAuth twit.TwitterAuth
	User        twit.TwitterUsers
	log         *logrus.Logger
}

func EelinController(db *gorm.DB, twitterAuth twit.TwitterAuth, user twit.TwitterUsers, logs *logrus.Logger) {
	usecase := EelinUsecase{
		Db:          db,
		TwitterAuth: twitterAuth,
		User:        user,
		log:         logs,
	}
	usecase.Execute()
}

//ツイート一覧取得、DBに書き出し、エクスポートまで
func (u EelinUsecase) Execute() {

	targetUserID := int64(69249063) //えーりんのID
	//tweets := getTweets(targetUserID)
	maxID := int64(0) //int64(1187241519784181760) - 1

	for x := 0; x < 10; x++ {
		tweets := twit.GetTimeline(targetUserID, u.TwitterAuth, maxID, 20)
		if len(tweets) == 0 {
			break
		}
		tweetLogs := twit.StructATweetLogs(tweets)
		tweetLogs.InsertByStruct(u.Db)

		maxID = tweets.GetFirstID() - 1
		//u.log.Printf("%# v", pretty.Formatter(tweets))
	}

	//u.log.SetFormatter(&logrus.TextFormatter{
	//	DisableColors: true,
	//	FullTimestamp: true})

	//u.log.SetFormatter(&logrus.JSONFormatter{})

	//fmt.Printf(pretty.Formatter(tweets))

	//
	//t := twit.TweetLog{
	//	Body:     article.Snippet,
	//	Country:  u.LocIndex,
	//	ParentId: parentId,
	//}
	//t.InsertByStruct(u.Db, u.User.AccountName)
	//twit.Debug = false
}
