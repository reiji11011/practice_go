package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer{}はWriterインターフェースを実装している。
	// Write()をメソッドとして持っている。
	buffer := bytes.Buffer{}

	// Writerとして送信し、書き込まれた内容を記憶する。
	// バッファーに書き込む。
	Greet(&buffer, "太郎")

	got := buffer.String()
	want := "Hello, 太郎"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
