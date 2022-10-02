package di

import (
	"fmt"
	"io"
	"os"
	"time"
)

func TddMockMain() {
	// mainでos.Stdoutに送信して、ターミナルに出力されたカウントダウンをユーザーに表示する。
	//Countdown(os.Stdout)

	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

// 旧
// 1秒ごとにカウントダウンをする。
//func Countdown(out io.Writer) {
//	for i := countdownStart; i > 0; i-- {
//		time.Sleep(1 * time.Second)
//		fmt.Fprintln(out, i)
//	}
//	time.Sleep(1 * time.Second)
//	fmt.Fprint(out, finalWord)
//}


// 「1秒ごとに」という条件をどのようにテストをするか？

// 課題
// Countdownはtime.Sleep()に依存している。そのためSleep()の挙動を制御できないので、真面目に書こうとするとテストの実行に4秒かかる。

// 解決策
// DIを行い、time.Sleep()の代わりをつくる。
// DIをする（依存関係をインターフェースに置き換える。）

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// 新
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		// インターフェースを使ってSleep()を表現することによって、テストにてスパイを使ってSleep()が呼び出された回数を記録できるようになった。
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
