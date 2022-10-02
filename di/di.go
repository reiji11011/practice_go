package di

import (
	"fmt"
	"time"
)

func Main() {
	fmt.Println(isMorning())

	now := time.Now()
	// 例えば任意の時間で実行してみる。
	//now := time.Unix(1660094749, 0)
	fmt.Println(isMorning2(now))

	// インターフェースを実装した構造体を引数として受け渡す。
	t := Time{}
	fmt.Println(isMorning3(t))
}

// 現在時刻を取得し、0~11なら午前中、12~23なら午後、その他ならエラーを返す関数を作成する。
func isMorning() bool {
	hour := time.Now().Hour()
	return 0 <= hour && hour <= 11
}

// 現状の問題点
// 外部のtimeパッケージを利用している（= 依存している）
// time.Now()の戻り値にisMorning()が依存しているため、テストが書きづらい。
// 午前中であることのテストと午後であることのテストを同時に書くことができない。

// 外部依存を引数で渡すようにする。
// その結果、Time型を自由に操れるようになる。
func isMorning2(now time.Time) bool {
	hour := now.Hour()
	return 0 <= hour && hour <= 11
}

// 関数の引数にインターフェースを渡してみる。
// 関数の引数をインターフェース化するメリットはよくわかっていない。
// 別にインターフェースを使わなくても実装をモックで書き換えることができている。
// todo インターフェースを使う理由を調査する。

type Timer interface {
	Now() time.Time
}

type Time struct{}

func (t Time) Now() time.Time {
	return time.Now()
}

func isMorning3(time Timer) bool {
	hour := time.Now().Hour()
	return 0 <= hour && hour <= 11
}
