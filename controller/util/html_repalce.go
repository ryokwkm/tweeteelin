package util

import "strings"

type htmlArray struct {
	Old string
	New string
}

func HtmlReplace(body string) string {
	replaces := []htmlArray{
		//{Old: "&bsol;", New: "\"}, // ちょっと危ないからバックスラッシュ
		{Old: "&lt;", New: "<"},   // 小なり
		{Old: "&gt;", New: ">"},   // 大なり
		{Old: "&amp;", New: "&"},  // アンパサンド
		{Old: "&quot;", New: "“"}, // クォーテーション
		{Old: "&apos;", New: "'"}, // アポストロフィ
		{Old: "&#39;", New: "'"},  // アポストロフィ
		{Old: "&#34;", New: "“"},  //
		{Old: "&#92;", New: "\\"}, //

		{Old: "&nbsp;", New: ""},     // 空白
		{Old: "&iquest;", New: "¿"},  // 逆疑問符
		{Old: "&iexcl;", New: "¡"},   // 逆感嘆符
		{Old: "&dagger;", New: "†"},  // ダガー
		{Old: "&Dagger;", New: "‡"},  // ダブルダガー
		{Old: "&para;", New: "¶"},    // 段落記号
		{Old: "&sect;", New: "§"},    // セクションマーク
		{Old: "&hellip;", New: "…"},  // 省略記号
		{Old: "&numero;", New: "№"},  // ナンバー
		{Old: "&oline;", New: "‾"},   // オーバーライン
		{Old: "&blank;", New: "␣"},   // 空白記号
		{Old: "&lsquo;", New: "‘"},   // シングルクォーテーションマーク(開始)
		{Old: "&rsquo;", New: "’"},   // シングルクォーテーションマーク(終了)
		{Old: "&ldquo;", New: "“"},   // ダブルクォーテーションマーク(開始)
		{Old: "&rdquo;", New: "”"},   // ダブルクォーテーションマーク(終了)
		{Old: "&phone;", New: "☎"},   // 電話マーク
		{Old: "&male;", New: "♂"},    // 雄記号
		{Old: "&female;", New: "♀"},  // 雌記号
		{Old: "&copy;", New: "©"},    // 著作権マーク
		{Old: "&reg;", New: "®"},     // 登録商標マーク
		{Old: "&trade;", New: "™"},   // トレードマーク
		{Old: "&sung;", New: "♪"},    // 八分音符
		{Old: "&flat;", New: "♭"},    // フラット
		{Old: "&sharp;", New: "♯"},   // シャープ
		{Old: "&natur;", New: "♮"},   // ナチュラル
		{Old: "&spades;", New: "♠"},  // スペード
		{Old: "&clubs;", New: "♣"},   // クラブ
		{Old: "&hearts;", New: "♥"},  // ハート
		{Old: "&diams;", New: "♦"},   // ダイヤ
		{Old: "&star;", New: "☆"},    // 白星
		{Old: "&starf;", New: "★"},   // 黒星
		{Old: "&sext;", New: "✶"},    // 六角星
		{Old: "&xutri;", New: "△"},   // 上向き正三角形
		{Old: "&xdtri;", New: "▽"},   // 下向き正三角形
		{Old: "&utrif;", New: "▴"},   // 上向き小正三角形(黒)
		{Old: "&utri;", New: "▵"},    // 上向き小正三角形(白)
		{Old: "&rtrif;", New: "▸"},   // 右向き小正三角形(黒)
		{Old: "&rtri;", New: "▹"},    // 右向き小正三角形(白)
		{Old: "&dtrif;", New: "▾"},   // 下向き小正三角形(黒)
		{Old: "&dtri;", New: "▿"},    // 下向き小正三角形(白)
		{Old: "&ltrif;", New: "◂"},   // 左向き小正三角形(黒)
		{Old: "&ltri;", New: "◃"},    // 左向き小正三角形(白)
		{Old: "&ultri;", New: "◸"},   // 直角三角形(左上)
		{Old: "&urtri;", New: "◹"},   // 直角三角形(右上)
		{Old: "&lltri;", New: "◺"},   // 直角三角形(左下)
		{Old: "&lrtri;", New: "⊿"},   // 直角三角形(右下)
		{Old: "&squ;", New: "□"},     // 正方形(白)
		{Old: "&squf;", New: "▪"},    // 正方形(黒)
		{Old: "&rect;", New: "▭"},    // 長方形
		{Old: "&marker;", New: "▮"},  // 縦長方形
		{Old: "&fltns;", New: "▱"},   // 平行四辺形
		{Old: "&loz;", New: "◊"},     // ひし形(白)
		{Old: "&lozf;", New: "⧫"},    // ひし形(黒)
		{Old: "&boxbox;", New: "⧉"},  // ２つ結合した正方形
		{Old: "&block;", New: "█"},   // ブロック
		{Old: "&uhblk;", New: "▀"},   // ブロック(上半分)
		{Old: "&lhblk;", New: "▄"},   // ブロック(下半分)
		{Old: "&blk14;", New: "░"},   // シェード(明るい)
		{Old: "&blk12;", New: "▒"},   // シェード(中間)
		{Old: "&blk34;", New: "▓"},   // シェード(暗い)
		{Old: "通貨・単位記号", New: ""},    //
		{Old: "文字実体参照", New: "表示"},   // 備考・説明
		{Old: "&yen;", New: "¥"},     // 円マーク
		{Old: "&dollar;", New: "$"},  // ドル
		{Old: "&cent;", New: "¢"},    // セント
		{Old: "&euro;", New: "€"},    // ユーロ
		{Old: "&pound;", New: "£"},   // ポンド
		{Old: "&curren;", New: "¤"},  // 国際通貨記号
		{Old: "&micro;", New: "µ"},   // マイクロ
		{Old: "&deg;", New: "°"},     // 度記号
		{Old: "&prime;", New: "′"},   // プライム(分)
		{Old: "&Prime;", New: "″"},   // ダブルプライム(秒)
		{Old: "&ell;", New: "ℓ"},     // リットル
		{Old: "&angst;", New: "Å"},   // オングストローム
		{Old: "&ohm;", New: "Ω"},     // オーム
		{Old: "&mho;", New: "℧"},     // モー
		{Old: "&plus;", New: "0"},    // プラス
		{Old: "&minus;", New: "−"},   // マイナス
		{Old: "&times;", New: "×"},   // 乗算記号
		{Old: "&divide;", New: "÷"},  // 除算記号
		{Old: "&equals;", New: "="},  // イコール
		{Old: "&plusmn;", New: "±"},  // プラスマイナス
		{Old: "&mnplus;", New: "∓"},  // マイナスプラス
		{Old: "&ne;", New: "≠"},      // 等しくない
		{Old: "&efDot;", New: "≒"},   // ほぼ等しい
		{Old: "&equiv;", New: "≡"},   // 合同
		{Old: "&le;", New: "≤"},      // 小なりイコール
		{Old: "&ge;", New: "≥"},      // 大なりイコール
		{Old: "&permil;", New: "‰"},  // パーミル（千分率）
		{Old: "&radic;", New: "√"},   // 根号（ルート）
		{Old: "&or;", New: "∨"},      // 論理和
		{Old: "&and;", New: "∧"},     // 論理積
		{Old: "&cup;", New: "∪"},     // 和集合
		{Old: "&cap;", New: "∩"},     // 共通部分
		{Old: "&sup;", New: "⊃"},     // 含む
		{Old: "&sub;", New: "⊂"},     // 含まれる
		{Old: "&sum;", New: "∑"},     // 総和
		{Old: "&prod;", New: "∏"},    // 総乗
		{Old: "&prop;", New: "∝"},    // 比例
		{Old: "&infin;", New: "∞"},   // 無限大
		{Old: "&there4;", New: "∴"},  // ゆえに
		{Old: "&becaus;", New: "∵"},  // なぜならば
		{Old: "&ang;", New: "∠"},     // 角
		{Old: "&angrt;", New: "∟"},   // 直角
		{Old: "&int;", New: "∫"},     // 積分
		{Old: "&Int;", New: "∬"},     // 二重積分
		{Old: "&conint;", New: "∮"},  // 周回積分
		{Old: "&nabla;", New: "∇"},   // ナブラ
		{Old: "&frac12;", New: "½"},  // ２分の１
		{Old: "&frac13;", New: "⅓"},  // ３分の１
		{Old: "&frac23;", New: "⅔"},  // ３分の２
		{Old: "&frac14;", New: "¼"},  // ４分の１
		{Old: "&frac34;", New: "¾"},  // ４分の３
		{Old: "&uarr;", New: "↑"},    // 上向き矢印
		{Old: "&darr;", New: "↓"},    // 下向き矢印
		{Old: "&larr;", New: "←"},    // 左向き矢印
		{Old: "&rarr;", New: "→"},    // 右向き矢印
		{Old: "&varr;", New: "↕"},    // 上下向き矢印
		{Old: "&harr;", New: "↔"},    // 左右向き矢印
		{Old: "&nwarr;", New: "↖"},   // 左上向き矢印
		{Old: "&nearr;", New: "↗"},   // 右上向き矢印
		{Old: "&searr;", New: "↘"},   // 右下向き矢印
		{Old: "&swarr;", New: "↙"},   // 左下向き矢印
		{Old: "&rlarr;", New: "⇄"},   // 右向き矢印＋左向き矢印
		{Old: "&lrarr;", New: "⇆"},   // 左向き矢印＋右向き矢印
		{Old: "&udarr;", New: "⇅"},   // 上向き矢印＋下向き矢印
		{Old: "&duarr;", New: "⇵"},   // 下向き矢印＋上向き矢印
		{Old: "&uuarr;", New: "⇈"},   // ペアの上向き矢印
		{Old: "&ddarr;", New: "⇊"},   // ペアの下向き矢印
		{Old: "&llarr;", New: "⇇"},   // ペアの左向き矢印
		{Old: "&rrarr;", New: "⇉"},   // ペアの右向き矢印
		{Old: "&uArr;", New: "⇑"},    // 上向き二重矢印
		{Old: "&dArr;", New: "⇓"},    // 下向き二重矢印
		{Old: "&lArr;", New: "⇐"},    // 左向き二重矢印
		{Old: "&rArr;", New: "⇒"},    // 右向き二重矢印
		{Old: "&vArr;", New: "⇕"},    // 上下向き二重矢印
		{Old: "&hArr;", New: "⇔"},    // 左右向き二重矢印
		{Old: "&nwArr;", New: "⇖"},   // 左上向き二重矢印
		{Old: "&neArr;", New: "⇗"},   // 右上向き二重矢印
		{Old: "&seArr;", New: "⇘"},   // 右下向き二重矢印
		{Old: "&swArr;", New: "⇙"},   // 左下向き二重矢印
		{Old: "&lAarr;", New: "⇚"},   // 左向き三重矢印
		{Old: "&rAarr;", New: "⇛"},   // 右向き三重矢印
		{Old: "&lsh;", New: "↰"},     // 上から左向き矢印
		{Old: "&rsh;", New: "↱"},     // 上から右向き矢印
		{Old: "&ldsh;", New: "↲"},    // 下から左向き矢印
		{Old: "&rdsh;", New: "↳"},    // 下から右向き矢印
		{Old: "&crarr;", New: "↵"},   // Enterキー・改行
		{Old: "&larrb;", New: "⇤"},   // 左タブ
		{Old: "&rarrb;", New: "⇥"},   // 右タブ
		{Old: "&zigrarr;", New: "⇝"}, // ジグザグ右向き矢印
		{Old: "&cularr;", New: "↶"},  // 上半円の反時計回り矢印
		{Old: "&curarr;", New: "↷"},  // 上半円の時計回り矢印
		{Old: "&olarr;", New: "↺"},   // 開いた円の反時計回り矢印
		{Old: "&orarr;", New: "↻"},   // 開いた円の時計回り矢印
		{Old: "&boxh;", New: "─"},    // 罫線 水平線
		{Old: "&boxv;", New: "│"},    // 罫線 垂直線
		{Old: "&boxvh;", New: "┼"},   // 罫線 水平線＋垂直線
		{Old: "&boxdr;", New: "┌"},   // 罫線 左上
		{Old: "&boxhd;", New: "┬"},   // 罫線 上
		{Old: "&boxdl;", New: "┐"},   // 罫線 右上
		{Old: "&boxvr;", New: "├"},   // 罫線 左
		{Old: "&boxvl;", New: "┤"},   // 罫線 右
		{Old: "&boxur;", New: "└"},   // 罫線 左下
		{Old: "&boxhu;", New: "┴"},   // 罫線 下
		{Old: "&boxul;", New: "┘"},   // 罫線 右下
		{Old: "&boxH;", New: "═"},    // 二重罫線 水平線
		{Old: "&boxV;", New: "║"},    // 二重罫線 垂直線
		{Old: "&boxVH;", New: "╬"},   // 二重罫線 水平線＋垂直線
		{Old: "&boxDR;", New: "╔"},   // 二重罫線 左上
		{Old: "&boxHD;", New: "╦"},   // 二重罫線 上
		{Old: "&boxDL;", New: "╗"},   // 二重罫線 右上
		{Old: "&boxVR;", New: "╠"},   // 二重罫線 左
		{Old: "&boxVL;", New: "╣"},   // 二重罫線 右
		{Old: "&boxUR;", New: "╚"},   // 二重罫線 左下
		{Old: "&boxHU;", New: "╩"},   // 二重罫線 下
		{Old: "&boxUL;", New: "╝"},   // 二重罫線 右下
		{Old: "&check;", New: "✓"},   // チェックマーク
		{Old: "&cross;", New: "✗"},   // エックスマーク
		{Old: "&copysr;", New: "℗"},  // マルPマーク
		{Old: "&target;", New: "⌖"},  // 位置表示
		{Old: "&telrec;", New: "⌕"},  // 電話録音記号
		{Old: "&vzigzag;", New: "⦚"}, // 縦ジグザグ線
		{Old: "&cirmid;", New: "⫯"},  // 垂直線上円
		{Old: "&midcir;", New: "⫰"},  // 垂直線下円
		{Old: "&malt;", New: "✠"},    // マルタ十字
		{Old: "&incare;", New: "℅"},  // care of…（…様方）
	}

	for _, replace := range replaces {
		body = strings.Replace(body, replace.Old, replace.New, -1)
	}

	return body
}

func Hiragana() []string {
	return []string{
		"ぁ",
		"あ",
		"ぃ",
		"い",
		"ぅ",
		"う",
		"ぇ",
		"え",
		"ぉ",
		"お",
		"か",
		"が",
		"き",
		"ぎ",
		"く",
		"ぐ",
		"け",
		"げ",
		"こ",
		"ご",
		"さ",
		"ざ",
		"し",
		"じ",
		"す",
		"ず",
		"せ",
		"ぜ",
		"そ",
		"ぞ",
		"た",
		"だ",
		"ち",
		"ぢ",
		"っ",
		"つ",
		"づ",
		"て",
		"で",
		"と",
		"ど",
		"な",
		"に",
		"ぬ",
		"ね",
		"の",
		"は",
		"ば",
		"ぱ",
		"ひ",
		"び",
		"ぴ",
		"ふ",
		"ぶ",
		"ぷ",
		"へ",
		"べ",
		"ぺ",
		"ほ",
		"ぼ",
		"ぽ",
		"ま",
		"み",
		"む",
		"め",
		"も",
		"ゃ",
		"や",
		"ゅ",
		"ゆ",
		"ょ",
		"よ",
		"ら",
		"り",
		"る",
		"れ",
		"ろ",
		"ゎ",
		"わ",
		"ゐ",
		"ゑ",
		"を",
		"ん",
		"ゔ",
		"ゕ",
		"ゖ",
	}
}
