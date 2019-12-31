package dbutil

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConfig struct {
	USER     string
	PASS     string
	PROTOCOL string
	DBNAME   string
}

func GormConnect(dbConfig DbConfig) *gorm.DB {
	DBMS := "mysql"
	USER := dbConfig.USER         //"root"
	PASS := dbConfig.PASS         //"root"
	PROTOCOL := dbConfig.PROTOCOL //"tcp(localhost:8889)"
	DBNAME := dbConfig.DBNAME     //"vacation_tweets"
	const OPTION = "interpolateParams=true&charset=utf8mb4"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

//db接続テスト
func main() {
	db := GormConnect(DbConfig{
		USER:     "root",
		PASS:     "root",
		PROTOCOL: "tcp(localhost:8889)",
		DBNAME:   "vacation_tweets",
	})
	defer db.Close()

	//db.SingularTable(true)
	db.LogMode(true)

	//accounts := []Locations{}
	//db.Find(&accounts, "is_realtime = 1")

	//spew.Dump(accounts)
}
