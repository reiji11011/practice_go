package bytes

import (
	"bytes"
	"fmt"
)

func Test() {
	fmt.Println("bytes")

	// 可変サイズのバイトバッファ
	// 入出力を一時的に保存する保存領域
	// ReadとWriteメソッドを持つ（io.Reader()インターフェースを満たす）
	var b bytes.Buffer
	fmt.Printf("%v %T", b, b)

	// バッファに書き込む
	b1 := []byte{0, 1}
	b2 := bytes.NewBuffer(b1)
	fmt.Println(b2)

	s := []byte("aaa")
	r1, _ := b2.Read(s)
	fmt.Println(r1)
	fmt.Println(string(s[:r1]))

	// バッファから読み出す
}
