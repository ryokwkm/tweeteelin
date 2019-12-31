package main

import (
	"../model/twit"
	"github.com/jinzhu/gorm"
	"github.com/kr/pretty"
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
	sinceID := int64(0)
	tweets := twit.GetTimeline(targetUserID, u.TwitterAuth, sinceID)
	u.log.Info("tweets: %v", len(tweets))
	tweetLogs := twit.StructATweetLogs(tweets)
	tweetLogs.InsertByStruct(u.Db)

	//u.log.SetFormatter(&logrus.TextFormatter{
	//	DisableColors: true,
	//	FullTimestamp: true})

	//u.log.SetFormatter(&logrus.JSONFormatter{})

	u.log.Info(len(tweets))
	u.log.Infof("tweets: %# v", pretty.Formatter(tweets))
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
