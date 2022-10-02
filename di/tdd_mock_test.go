package di

import (
	"bytes"
	"testing"
)

// スパイでSleepが呼び出された回数を記録する。
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// 旧
//func TestCountDown(t *testing.T) {
//	// テストでは、bytes.Bufferに送信して、生成されるデータをテストでキャプチャできるようにする。
//	buffer := &bytes.Buffer{}
//	Countdown(buffer)
//
//	got := buffer.String()
//	want := `3
//2
//1
//Go!`
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}

func TestCountDown(t *testing.T) {
	// テストでは、bytes.Bufferに送信して、生成されるデータをテストでキャプチャできるようにする。
	buffer := &bytes.Buffer{}
	//
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
