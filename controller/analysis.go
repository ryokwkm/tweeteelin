package main

import (
	"strings"

	html_replace "github.com/ryokwkm/htmlreplace"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/jinzhu/gorm"
	"github.com/ryokwkm/trends-model/twit"
	"github.com/sirupsen/logrus"
)

type AnalysisUsecase struct {
	Db          *gorm.DB
	TwitterAuth twit.TwitterAuth
	User        twit.TwitterUsers
	log         *logrus.Logger
}

func AnalysisController(db *gorm.DB, twitterAuth twit.TwitterAuth, user twit.TwitterUsers, logs *logrus.Logger) {
	usecase := AnalysisUsecase{
		Db:          db,
		TwitterAuth: twitterAuth,
		User:        user,
		log:         logs,
	}
	usecase.Execute()
}

//ツイート一覧取得、DBに書き出し、エクスポートまで
// t.Analyze("寿司が食べたい。", tokenizer.Normal)
func (u AnalysisUsecase) Execute() {
	t := tokenizer.New()
	words := getWords()

	for _, word := range words {
		word = html_replace.GetOnlyText(word)
		tokens := t.Tokenize(word)
		u.PrintAnalysis(tokens)
		//u.GetMeishiWord(tokens)
	}
}

func (u AnalysisUsecase) GetMeishiWord(tokens []tokenizer.Token) []string {
	ret := []string{}

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		feature := token.Features()
		if feature[0] == "名詞" && feature[6] == "*" {
			u.log.Printf("%s  %v\n", token.Surface, strings.Join(token.Features(), ","))
			ret = append(ret, token.Surface)
		}
		//features := strings.Join(token.Features(), ",")
		//u.log.Printf("%s\t%v\n", token.Surface, features)
	}
	return ret
}

func (u AnalysisUsecase) PrintAnalysis(tokens []tokenizer.Token) {
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			//fmt.Printf("continue --> %s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		u.log.Printf("%s   %v\n", token.Surface, features)
	}
	u.log.Printf("-----------------------")
}

func getWords() []string {
	ret := []string{
		//
		//		`好きドラス…… https://t.co/qRq7xnVu3y`,
		//		`【  #新春メギドくじ  】
		//ラッキーメギドは『アンドラス』
		//あなたの勝算は・・・解剖、分解、観察で人助け！
		//
		//…つづきを見る
		// https://t.co/be43y6qW4R
		//アアッッッ/////アッ/////////ありがとう////////`,
		//		`あけましておめでとうございます！第五人格新年みくじで、2020年の運勢を占おう！ https://t.co/Spod58dOiX
		//リッペャアさん https://t.co/D3sIzOip1e`,
		//		`アモンかわ`,
		//		`@__hinoto はい……すいませんでした……`,
		//		`うっ……年越した直後に回せば出るかも☆と思って回したけど……ねずみ年のやつがきた……触媒的には合ってるんやけど違うんや……`,
		//		`………………あ、ありが……と、う………………🔨 https://t.co/TREKzYiAzM`,
		//		`2019年に「あけましておめでとうございま」
		//2020年に「す」で始めようと思ったのに普通に年越しの挨拶になってしまった……`,
		//		`す`,
		//		`あけましておめでとうございま`,
		//		`食育`,
		//		`@xEGPB4xmuN3SKpr 見たがるなよwwwwwwとりあえずただでは帰さないと思いますwww`,
		//		`2019年もありがとうございました！いい人生でした！`,
		//		`は？サイコーか？`,
		//		`SP:悟空猿鬼
		//SP:ハニー
		//N:四桜犬士郎
		//N:六力大仙
		//N:七色虹升
		//N:三蔵法子
		//SR:双六一
		//N:ロック
		//N:一声三鶴
		//R:七夕星太郎
		//	#shindanmaker
		//https://t.co/dp7N2uGMjh
		//	ハニーくんと六力さん出たんで優勝です`,
		//		`福さんwwwwwwwww`,
		//		`@xEGPB4xmuN3SKpr めっちゃ挑発してくるwwwwwwwww実際そうなったらどうするかな……ｳﾌﾌ`,
		//		`かっけえええええ！！！！！！！！`,
		//		`@xEGPB4xmuN3SKpr ほんとそういうの好きねwwwwwwwめちゃめちゃSAN値減りそうwwwwww`,

		`text`,
		`MOTHER警察(風月)に捕まった`,
		`リュカめちゃくちゃ使いやすいな……スマッシュ後の隙が大きいけど`,
		`弟と地獄みたいにスマブラやった`,
		`@ryokwkm まじでやんのwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwいいけどwwwwwwwwwwwwwwwwwwまじで誰得なんwwwwwwwww全体的にフリー素材なんで大丈夫です……wwwwwwww`,
		`ボンボンちゃんかわいい`,
		`YEEEEEEEEEEEEEEEEEEE

私と一緒に「identityV」で遊ぼう！ https://t.co/ONjfA4aVDX`,
		`すごい幸せそうな顔してる`,
		`好きドラス…… https://t.co/qRq7xnVu3y`,
		`【  #新春メギドくじ  】
ラッキーメギドは『アンドラス』
あなたの勝算は・・・解剖、分解、観察で人助け！

…つづきを見る
 https://t.co/be43y6qW4R
アアッッッ/////アッ/////////ありがとう////////`,
		`あけましておめでとうございます！第五人格新年みくじで、2020年の運勢を占おう！ https://t.co/Spod58dOiX
リッペャアさん https://t.co/D3sIzOip1e`,
		`アモンかわ`,
		`@__hinoto はい……すいませんでした……`,
		`うっ……年越した直後に回せば出るかも☆と思って回したけど……ねずみ年のやつがきた……触媒的には合ってるんやけど違うんや……`,
		`………………あ、ありが……と、う………………🔨 https://t.co/TREKzYiAzM`,
		`2019年に「あけましておめでとうございま」
2020年に「す」で始めようと思ったのに普通に年越しの挨拶になってしまった……`,
		`す`,
		`あけましておめでとうございま`,
		`@xEGPB4xmuN3SKpr 創造主の親心的にはあんま対峙してほしくないけどな((( 妙なことしないで幸せになってほしい((`,
		`@xEGPB4xmuN3SKpr ワクワクすんね！！！！！！！！`,
		`@Momir_Aluren マジの方の金欠だからガチャピン先生の力借りれなきゃ無理ですね……`,
		`@xEGPB4xmuN3SKpr 邪神かーなるほど、生物やな`,
		`@Momir_Aluren 私にはノアくん出なきゃ意味ないのだ……`,
		`食育`,
		`@xEGPB4xmuN3SKpr 見たがるなよwwwwwwとりあえずただでは帰さないと思いますwww`,
		`2019年もありがとうございました！いい人生でした！`,
		`は？サイコーか？`,
		`SP:悟空猿鬼
		SP:ハニー
		N:四桜犬士郎
		N:六力大仙
		N:七色虹升
		N:三蔵法子
		SR:双六一
		N:ロック
		N:一声三鶴
		R:七夕星太郎
		#shindanmaker
		https://t.co/dp7N2uGMjh
		ハニーくんと六力さん出たんで優勝です`,
		`福さんwwwwwwwww`,
		`@xEGPB4xmuN3SKpr めっちゃ挑発してくるwwwwwwwww実際そうなったらどうするかな……ｳﾌﾌ`,
		`かっけえええええ！！！！！！！！`,
		`@xEGPB4xmuN3SKpr ほんとそういうの好きねwwwwwwwめちゃめちゃSAN値減りそうwwwwww`,
		`開幕でシバいて離れないサバイバーちょっと焦るよね。え？まじか？命知らずか？優鬼していいの？みたいな((((`,
		`どう見ても200連昨日じゃないやろ……`,
		`ノアくん出ず https://t.co/cqEGjvvOLG`,
		`@xEGPB4xmuN3SKpr ママみ出ちゃうじゃん？「全力出してたのは見てたらわかるよ、頑張ったね。晩ごはん食べてくでしょ？帰ろ〜」って撫でてあげよう……`,
		`地元……物価が高い……(震え)`,
		`ぜったいほしい(血涙)`,
		`ビクターくんかわいい`,
		`動いたら体の痛みそうでもないわﾊﾊﾊﾊ`,
		`@xEGPB4xmuN3SKpr 平和か？wwwwww
権利逆転するタイミングあんのかな……🤔`,
		`今起きたんだけどバスのダメージが今きて体バキバキでケツ痛くてやばい。もう寝たい(？)`,
		`あの野郎(弟)TRPG始めたら「お？国家権力に逆らうんか？」ってすぐNPCとかKPCですら絞め技掛けようとしてくる💢「あ？こっちには確定魅了あんねんぶち抜くぞオラ」って返してる(どっこいどっこい)`,
		`@xEGPB4xmuN3SKpr 「………………もぉ～～、好き～～～～～♡許す～～♡♡♡」って抱っこしてぐるぐる回り始めるんで本当にチョロいし彼女に甘いですね() え？マジで彼氏と彼女の立場逆で草wwwww`,
		`村人かネスみたいな似たフォルムのキャラしか使ってないの草`,
		`スマブラのソフト自体は半年以上前からあったと言うのに`,
		`【速報】実家がSwitchを買いました【今スマブラやってる】`,
		`@Rurirena とか言ってる間に完食～そして完飲～砂漠そうそう～`,
		`@Rurirena ウマになったわね……🐁`,
		`ロックブーケちゃんサイコーカワイイよーーー！！！！！！！！`,
		`@Rurirena ヤ゛ダァ～～美味しそう～～～♡♡♡`,
		`大蛇丸のオの字も知らない母親に潜影邪手を覚えさせた`,
		`@xEGPB4xmuN3SKpr これはあれだな、「本格的な職質になる前に恋人ですって言ってほしかったなー…ちょっとショックー…」とかってあとで郁が拗ねるから頑張ってカバーしてね(丸投げ)`,
		`@xEGPB4xmuN3SKpr 可哀想だろ！！！！！！！！(手叩いて喜びながら)`,
		`謎の有償石1個持ってたから若干軽減される……なんの余りだろう……`,
		`は？かわいいこれ踊り食いできる`,
		`美人……`,
		`@tPoEGG 親めっちゃ報告受けたらしいけど全然起きませんでした☆☆☆ただいま帰りました！！！！！！！！`,
		`再度階段を駆け登り親に報告に行く猫「ﾄﾝﾄﾝﾄﾝﾁﾘﾝﾁﾘﾝﾁﾘﾝ」`,
		`いち早く起きて階段を降りてくる猫「॑⸜(* ॑꒳ ॑*  )⸝⋆*ﾄﾝﾄﾝﾄﾝ」
		様子をうかがう猫「(´⊙ω| 」`,
		`エデン:1
テュルソス:2
イクサバ:1
グリダーヴォル:1
ココ&amp;ミミ:1

ミカエル:1
ゴリラ:1

マルキアレス、ユーリ、ブリジール&amp;コーデリアをお迎え https://t.co/CKS9smSv2F`,
		`ぐ……ぐやぶゆがデレた…… https://t.co/ZAIcyk6JtU`,
		`帰宅時点で家族全員寝てんの面白すぎんか(リビングでぼっち)`,
		`ケツの痛みが腰痛に進化したんだけど`,
		`ていうか雨か 残念`,
		`お尻がめっちゃ痛い`,
		`@irisee113 ってお前は寝んのかーい！💥`,
		`@irisee113 寝るのやめるわ🥳🥳🥳`,
		`眠くないたすけて`,
		`いま海老名SAらしい`,
		`@chikichikibon 病院あってよかったよかった！
やばそう；；；；；大病じゃないことを祈る……`,
		`バス乗った！酔うまでチキンレースやな！！`,
		`まだバス来るまで時間あるし貯めとこう……`,
		`バス搭乗前になってワイヤレスイヤンホホの充電がピンチになるかなしみ`,
		`今流行りのマンキンじゃないですかぁ！ https://t.co/BJbAwIDrHh`,
		`@chikichikibon えっ何があったの；；；；大丈夫？ちゃんと病院行くんやで｡ﾟ( ﾟஇωஇﾟ)ﾟ｡`,
		`ｷﾞｬｰｰｰｰｰｰｰｰｰ`,
		`おぽこバーサーカーモード`,
		`@opokodayo wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwごwwwwwwwwwwwwwwwwwwごめんてwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww… https://t.co/DbthH8xggq`,
		`@opokodayo (๑&gt;؂•̀๑)
ん？？？何がついてないのかな？？？みーさんに教えて？？？？？？(くそ)`,
		`ぽこちゃんはド天然……の男の娘…… https://t.co/wGfdlnL6ij`,
		`@Tommystwa おぽこは男の娘……！おぼえた`,
		`おぽこは……おちんちん……？(錯乱)`,
		`@__hinoto アッッッッこのっ……散らかっ……か、かわっ……ウウッ……猫ッ……散らっ……か、かわいい……(即落ち)`,
		`家出る私完璧すぎたから聞いて。シーツ枕カバー毛布全て洗濯し、きちんとベッドメイキングして邪魔にならない程度にお香焚いて、加湿器の水を捨てて電源を切り、部屋の戸締りと消灯を全て確認し、ゴミを全て捨て、トイレ掃除と洗面台のパイプも軽く掃除しておいた`,
		`@chicaya_fgo ありがとう・:*三ᕕ( ᐛ )ᕗいってきやす！！`,
		`じゃ でっぱつ`,
		`あった(秒速発見)`,
		`朝あったじゃん！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！どこ！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！`,
		`あ！まだやることある！よかったー時間があって！

イヤンホホ紛失(死活問題)`,
		`何時に出よう、あんまり早く出ても暇すぎるんだよね`,
		`精神やばい時になるやつで自覚あるのは「お腹はすいてるのに口にもの入れたくない」です。まじでお腹はグーグー鳴ってんのに口に何か入ってる状態がめっちゃイヤ`,
		`@ElnathIV 2021年ぐらいとかリアリティがやばいwww
主犯がピアソンじゃなぁ……懐中電灯しか持ってないしな……向いてないな(真顔)`,
		`ゲームなんかやりたい時やればいいんだよ！！！やりたくない時は違うことしようぜ！！！！！`,
		`@ElnathIV ──2019/12/25、公開──(納期割れ)

わたしコラ技術皆無wwwwwwwwwこれも仲間入れて https://t.co/9hVmvcW8cb`,
		`@ElnathIV そんな……貴方は知らないうちに大人になってしまったのね……いいわ……サンタが幸せだけを届けると思ったら大間違いよ……絶望をデリバリーしてやるわ……ピアソンと共に！！！！！！！！`,
		`@ElnathIV え……？泥棒は遊びだったっていうの……！？ひどい、ひどいわ……！！！ https://t.co/1F7VIIrYuF`,
		`ママとバブちゃんが同時に来た https://t.co/7OuIt5qD8h`,
		`やっぱ私ガウェイン教だわ。日光浴びながらガチャ回すと出るわ。`,
		`ハスター様とチェイス始めたワイクラーク「へへっ今度は私と遊んでくれんの？」
ハスター様が途中で近くにいた心眼に殴りかかっていくのを見たワイクラーク「え……まじかよ、王浮気すんの……？おい、ちょっ……クソッこっち見ろよ！浮気すんな！」`,
		`最近チェイスとかやってる時に出る悲鳴
「スケベ！！」「はああ～～地雷！私が地雷！」「当たらんのですよ！そんなものは！(調子乗ってる時)」「(被弾)下手くそかよ」「(風船)死んだ方がマシ」「(椅子)もうほっといてください…」

優鬼… https://t.co/lAEVMSorjP`,
		`画像系フォルダが驚きの53G食ってたからスクショ整理始めたんだけどどうしてもアヴェンジャーのスクショは置いといてしまう((((((((`,
	}

	return ret
}
