package main

import (
	"../lib/dbutil"
	"./config"
	"./util"
	"github.com/ryokwkm/trends-model/twit"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	log = util.CreateMyLog()
	MODE := config.MODE_DEBUG

	//dbへ接続
	//deferは関数終了時に実行されるので、 main() 以外で実行するとすぐにcloseされてしまう
	db := dbutil.GormConnect(config.DbConfig(MODE))
	defer db.Close()

	//デバッグ
	if MODE == config.MODE_DEBUG {
		//twit.Debug = true
		users := twit.GetDebugUser(db)
		for _, user := range users {
			if twitterAuth, err := twit.GetTwitterUserAuth(db, user.Id); err == false {
				EelinController(db, twitterAuth, user, log)
			}
		}
	}
}
