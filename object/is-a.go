package object

import "fmt"

type Clock struct {
	Minutes uint32
	Seconds uint32
}

// todo データを保持するようにしたい。
func (c Clock) AddSeconds() {
	c.Seconds += 1
	fmt.Println(c.Seconds)
}

// 以下構造体にClockの機能を持たせたい。
type DigitalClock struct {
	Clock
	View uint32
}

func Isa() {
	fmt.Println("is-a 関係のテスト")

	d := DigitalClock{}
	// 構造体を埋め込むことによって埋め込み先のメソッドやデータにアクセスできる。
	fmt.Println(d.Minutes)
	fmt.Println(d.Seconds)
	d.AddSeconds()
}
