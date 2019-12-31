package config

import "github.com/ryokwkm/trends-model/twit"
import "../../lib/dbutil"

const (
	MODE_PROD = 1 + iota //1
	MODE_MAINTE
	MODE_FATE

	MODE_DEBUG = 99
)

const (
	FUNC_TYPE_SPORTS = 1 + iota
	FUNC_TYPE_FATE   = 2
	FUNC_TYPE_ENT    = 3
	FUNC_TYPE_SEARCH = 9
	FUNC_TYPE_MAINTE = 99
)

const (
	MODE_TYPE_TWEET       = 0 + iota
	MODE_TYPE_ONLY_SEARCH = 1
)
const (
	NEWS_API_KEY = "be4b0579075d40a48719a1c7f6db0b58" //kwkmlight
)

var TwitterId int64 //フォロバで使用

/**
 *	98	:リモート開発
 *	99	:ローカル開発
 */
func Config(mode int) twit.TwitterAuth {

	switch mode {

	case MODE_DEBUG:
		//develop : kwkmlight
		TwitterId = 709547666158395392 //kwkmlight
		return twit.TwitterAuth{
			ConsumerKey:    "0cOKasSPL84vRnGvDVkZk8khD",
			ConsumerSecret: "GuSPd1JKS8sj9BTvQYFs0l0ugyX5VV55axmqsCJXRb7cVSEcMv",
			AccessToken:    "709547666158395392-0cQJ4jlLiydmgoA9w4p5NjVkhtgnSz3",
			AccessSecret:   "aKNopqCTVSDXj4C1eHNEzu113mVd8zAIuHgIoY6sL4cfg",
		}

	default:
		//develop : kwkmlight
		TwitterId = 709547666158395392 //kwkmlight
		return twit.TwitterAuth{
			ConsumerKey:    "0cOKasSPL84vRnGvDVkZk8khD",
			ConsumerSecret: "GuSPd1JKS8sj9BTvQYFs0l0ugyX5VV55axmqsCJXRb7cVSEcMv",
			AccessToken:    "709547666158395392-0cQJ4jlLiydmgoA9w4p5NjVkhtgnSz3",
			AccessSecret:   "aKNopqCTVSDXj4C1eHNEzu113mVd8zAIuHgIoY6sL4cfg",
		}
	}
}

func DbConfig(mode int) dbutil.DbConfig {
	if mode <= MODE_FATE {
		return dbutil.DbConfig{
			USER:     "vacation_tweets",
			PASS:     "barenai00",
			PROTOCOL: "tcp(localhost:3306)",
			DBNAME:   "vacation_tweets",
		}
	} else if mode == 999 {
		//develop :
		return dbutil.DbConfig{
			USER:     "vacation_tweets",
			PASS:     "barenai00",
			PROTOCOL: "tcp(localhost:3306)",
			DBNAME:   "vacation_tweets",
		}
	}

	//develop : mampのローカル
	return dbutil.DbConfig{
		USER:     "root",
		PASS:     "root",
		PROTOCOL: "tcp(localhost:8889)",
		DBNAME:   "vacation_tweets",
	}
}
