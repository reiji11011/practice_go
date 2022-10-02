package di

import (
	"fmt"
	"time"
)

// 参考：https://ken-aio.github.io/post/2019/10/17/go-test-interface/

// 現在時刻を取得するインターフェースを定義する。
type ITimeManager interface {
	Now() time.Time
}

type TimeManager struct{}

// 現在の時刻を返却する関数を定義する。
func (t *TimeManager) Now() time.Time {
	return time.Now()
}

func Sample() {
	UnixTimeSample()
	// TimeManagerをDIする。
	UnixTimeSample2(&TimeManager{})

	t := time.Unix(1664094749, 0)
	fmt.Println(t)
	fmt.Println(t.Hour())
}

// Unix時間が偶数ならEven、奇数ならOddを返却する関数
func UnixTimeSample() string {
	unixNow := time.Now().Unix()
	fmt.Println(unixNow)
	if unixNow%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// インターフェースを引数として渡すことによってユニットテストでのモック化ができるようになる。
func UnixTimeSample2(timeManager ITimeManager) string {
	unixNow := timeManager.Now().Unix()
	fmt.Println(unixNow)
	if unixNow%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
