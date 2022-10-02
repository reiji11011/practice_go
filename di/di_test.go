package di

import (
	"testing"
	"time"
)

// 現在時刻に依存している
// 外部環境に依存しているものをテストの関数内で使うとユニットテストが上手く動かせなくなる。
func TestIsMorning(t *testing.T) {
	// 実行する時間帯によってテストの結果が変わる。
	got := true
	want := isMorning()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// 午前中を返却する
func MockNowMorning() time.Time {
	// Time型のインスタンスを作成する。
	return time.Unix(1660094749, 0)
}

// 午後を返却する
func MockNowNoon() time.Time {
	// Time型のインスタンスを作成する。
	return time.Unix(1666094749, 0)
}

// モックを使ってテストを記述する。
// 現在時刻に依存しない形でテストを書ける。
func TestIsMorning2(t *testing.T) {
	// 午前の場合のテスト
	got := true
	now := MockNowMorning()
	want := isMorning2(now)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	// 午後の場合のテスト
	got = false
	now = MockNowNoon()
	want = isMorning2(now)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// 上記の記述をリファクタリングする。
func TestIsMorning2_2(t *testing.T) {
	// データを用意する。
	tests := []struct {
		time time.Time
		want bool
	}{
		{time: time.Unix(1660094749, 0), want: true},
		{time: time.Unix(1666094749, 0), want: false},
	}

	// 検証
	for _, test := range tests {
		got := isMorning2(test.time)
		if test.want != got {
			t.Errorf("エラーです")
		}
	}
}
