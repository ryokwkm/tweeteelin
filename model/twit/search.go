package twit

import (
	"github.com/dghubble/go-twitter/twitter"
	"time"
)

/**
 * ツイートを検索。
 * 指定した言語で検索
 */
func GetTimeline(userID int64, twitterAuth TwitterAuth, sinceID int64) []twitter.Tweet {
	client := getClient(twitterAuth)
	params := &twitter.UserTimelineParams{
		UserID:          userID,
		TrimUser:        Bool(false),
		IncludeRetweets: Bool(true),
		//Count:200
	}
	if sinceID > 0 {
		params.SinceID = sinceID
	}
	tweets, _, _ := client.Timelines.UserTimeline(params)

	time.Sleep(2 * time.Second)

	return tweets
}

func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}
