package di

import (
	"testing"
	"time"
)

// 好きな時間をNowとして取得できるようにする。
type MockTimeManager struct {
	MockTime *time.Time
}

// 構造体のフィールドを返却する（= 自分の好きな時間を返却する）
func (t *MockTimeManager) Now() time.Time {
	if t.MockTime == nil {
		return time.Now()
	}
	return *t.MockTime
}

// 現在時刻に依存するため、Evenになったり、Oddになったりする。
func TestUnixTimeSample(t *testing.T) {
	want := "Odd"
	got := UnixTimeSample()
	if want != got {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func TestUnixTimeSample2(t *testing.T) {
	// テストケースの用意
	cases := []struct {
		t    time.Time
		want string
	}{
		{t: time.Unix(1664094749, 0), want: "Odd"},
	}

	// 構造体を初期化する
	mockTimeManager := &MockTimeManager{}

	for _, c := range cases {
		// 自分の好きな時間を設定する。
		mockTimeManager.MockTime = &c.t
		got := UnixTimeSample2(mockTimeManager)
		if got != c.want {
			t.Errorf("want %v, but got %v", c.want, got)
		}
	}
}
