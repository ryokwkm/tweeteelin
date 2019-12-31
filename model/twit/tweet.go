package twit

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

type TwitterAuth struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

var Log *logrus.Logger
var client *twitter.Client
var Debug bool //初期値false Tweetしない

func getClient(twitterAuth TwitterAuth) *twitter.Client {
	if twitterAuth.AccessToken == "" {
		logrus.Fatal("Application Access Token required")
	}

	config := oauth1.NewConfig(twitterAuth.ConsumerKey, twitterAuth.ConsumerSecret)
	token := oauth1.NewToken(twitterAuth.AccessToken, twitterAuth.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = twitter.NewClient(httpClient)
	return client
}

func Tweet(tweetWord string, twitterAuth TwitterAuth, params *twitter.StatusUpdateParams) {
	client := getClient(twitterAuth)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	client.Accounts.VerifyCredentials(verifyParams)
	//fmt.Printf("User's ACCOUNT:", user)

	// Home Timeline
	//homeTimelineParams := &twitter.HomeTimelineParams{
	//	Count:     2,
	//	TweetMode: "extended",
	//}
	//tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	//fmt.Printf("User's HOME TIMELINE:", tweets)

	if !Debug {
		time.Sleep(2 * time.Second)
		client.Statuses.Update(tweetWord, params) //t, _, e :=
	} else {
		Log.Print(tweetWord)
	}
}

func GetTweetById(tweetId int64, twitterAuth TwitterAuth) *twitter.Tweet {
	client := getClient(twitterAuth)
	time.Sleep(2 * time.Second)
	statusShowParams := &twitter.StatusShowParams{}
	tweet, _, _ := client.Statuses.Show(tweetId, statusShowParams)
	return tweet
}

func Tweets(keywords []string, twitterAuth TwitterAuth) {
	if len(keywords) > 0 {
		for _, keyword := range keywords {
			time.Sleep(1 * time.Second)
			Tweet(keyword, twitterAuth, nil)
		}
	}
}

func reply(keyword string, status twitter.Tweet, twitterAuth TwitterAuth) {
	params := &twitter.StatusUpdateParams{
		InReplyToStatusID: status.ID,
		PossiblySensitive: nil,
	}
	keyword = "@" + status.User.ScreenName + " " + keyword
	time.Sleep(30 * time.Second)
	Tweet(keyword, twitterAuth, params)
}

//いいね
func PutFavorites(status twitter.Tweet) {
	if !Debug {
		time.Sleep(2 * time.Second)
		params := &twitter.FavoriteCreateParams{ID: status.ID}
		client.Favorites.Create(params) //tweet, _, _ :=
	}
}

//フォロー。フォロバにのみ使用
func PutFollow(userIds []int64) {
	for _, userId := range userIds {
		if !Debug {
			time.Sleep(2 * time.Second)
			params := &twitter.FriendshipCreateParams{UserID: userId}
			client.Friendships.Create(params) //t, _, e :=
		} else {
			Log.Print(fmt.Sprintf("follow ! %d", userId))
		}
	}

}

//リツイート
func PutRetweet(status twitter.Tweet) {
	if !Debug {
		time.Sleep(2 * time.Second)
		//params = &twitter.StatusRetweetParams{ID: status.ID,}
		client.Statuses.Retweet(status.ID, nil)
	} else {
		Log.Print("retweet -->" + status.Text)
	}
}

func MakeFeelingLuckyURL(keyword string) string {
	return "http://www.google.com/search?tbs=qdr:d&btnI=I%27m+Feeling+Lucky&lr=lang_ja&ie=UTF-8&oe=UTF-8&q=" + url.QueryEscape(keyword)
}

//自分のフォロワー一覧取得。フォロバ機能で使用
func GetMyFollower(twitterAuth TwitterAuth, userId int64) []int64 {
	params := &twitter.FollowerIDParams{
		UserID: userId,
		//Count:  5,
		//Cursor: 1516933260114270762,
	}
	time.Sleep(2 * time.Second)
	client := getClient(twitterAuth)
	followerIDs, _, _ := client.Followers.IDs(params)

	var returnIds []int64
	if len(followerIDs.IDs) > 0 {
		for _, userId := range followerIDs.IDs {
			returnIds = append(returnIds, userId)
		}
	}
	return returnIds
}

//リツイートしてくれたユーザーのID一覧を取得
func GetMyRetweeter(twitterAuth TwitterAuth) []int64 {
	client := getClient(twitterAuth)
	retweetTimelineParams := &twitter.RetweetsOfMeTimelineParams{
		Count:     2,
		TweetMode: "extended",
	}
	time.Sleep(2 * time.Second)
	tweets, _, _ := client.Timelines.RetweetsOfMeTimeline(retweetTimelineParams)

	var reTweetIds []int64
	for _, tweet := range tweets {
		reTweetIds = append(reTweetIds, tweet.ID)
	}

	var retweetUserIds []int64
	for _, retweetId := range reTweetIds {
		time.Sleep(2 * time.Second)
		params := &twitter.StatusRetweetsParams{}
		retweets, _, _ := client.Statuses.Retweets(retweetId, params)
		for _, retweet := range retweets {
			retweetUserIds = append(retweetUserIds, retweet.User.ID)
		}

	}
	return retweetUserIds
}

func GetMyFrends(twitterAuth TwitterAuth, userId int64) []int64 {
	params := &twitter.FriendIDParams{
		UserID: userId,
	}
	time.Sleep(2 * time.Second)
	client := getClient(twitterAuth)
	friendIDs, _, _ := client.Friends.IDs(params)

	var returnIds []int64
	if len(friendIDs.IDs) > 0 {
		for _, userId := range friendIDs.IDs {
			returnIds = append(returnIds, userId)
		}
	}
	return returnIds
}
