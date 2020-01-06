package main

import (
	"fmt"

	html_replace "github.com/ryokwkm/htmlreplace"

	"github.com/ryokwkm/markov"

	"github.com/jinzhu/gorm"
	"github.com/ryokwkm/trends-model/twit"
	"github.com/sirupsen/logrus"
)

type MarkovUsecase struct {
	Db          *gorm.DB
	TwitterAuth twit.TwitterAuth
	User        twit.TwitterUsers
	log         *logrus.Logger
}

func MarkovController(db *gorm.DB, twitterAuth twit.TwitterAuth, user twit.TwitterUsers, logs *logrus.Logger) {
	usecase := MarkovUsecase{
		Db:          db,
		TwitterAuth: twitterAuth,
		User:        user,
		log:         logs,
	}
	usecase.Execute()
}

//ツイート一覧取得、DBに書き出し、エクスポートまで
// t.Analyze("寿司が食べたい。", tokenizer.Normal)
func (u MarkovUsecase) Execute() {
	u.MarkovText()
}

func (u MarkovUsecase) MarkovText() {
	tweets := getTestWords()
	tweets = html_replace.GetOnlyTexts(tweets)

	count := 30
	h := markov.NewMarkov(tweets, u.log)
	for i := 0; i < count; i++ {
		fmt.Println(h.MakeWord())
	}
}

func getTestWords() []string {
	l := []string{

		`ハスター様とチェイス始めたワイクラーク「へへっ今度は私と遊んでくれんの？」
ハスター様が途中で近くにいた心眼に殴りかかっていくのを見たワイクラーク「え……まじかよ、王浮気すんの……？おい、ちょっ……クソッこっち見ろよ！浮気すんな！」`,
		`最近チェイスとかやってる時に出る悲鳴
「スケベ！！」「はああ～～地雷！私が地雷！」「当たらんのですよ！そんなものは！(調子乗ってる時)」「(被弾)下手くそかよ」「(風船)死んだ方がマシ」「(椅子)もうほっといてください…」

優鬼… `,
	}

	return l
}
