package io

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

type Text struct{}

// io.ReaderインターフェースをText構造体が実装をする
func (t Text) Read(p []byte) (int, error) {
	// バイト文字を読み込む
	// バイト文字を返却する
	return 1, nil
}

// 入出力処理の抽象化
func Test() {
	// 入力処理のインターフェース
	//io.Reader

	// 以下は全てio.Reader()インターフェースを実装している。
	// ファイル、メモリ、文字列、ネットワークの「読み込み」を入力処理を共通化している。

	// ファイルを読み込む
	file := os.File{}
	fmt.Printf("%v は %T 型 \n", file, file)
	// メモリへ書き込まれたデータを読み込む
	b := bytes.Buffer{}
	fmt.Printf("%v は %T 型 \n", b, b)
	// 文字列から読み込みを行う
	s := strings.Reader{}
	fmt.Printf("%v は %T 型 \n", s, s)
	// ネットワーク情報を読み込む
	n := net.TCPConn{}
	fmt.Printf("%v は %T 型 \n", n, n)

	// bytes.Buffer{}
	// バイトを扱うのに便利（らしい）
	// 文字列からバイト列にしたい
	// バイト列から文字列にしたい
	// io.Readerとして使いたい
	// io.Writerとして使いたい

	// ゼロ値
	var buf bytes.Buffer
	fmt.Printf("%v は %T 型 \n", buf, buf)

	i := []byte("a")
	r, err := buf.Write(i)
	if err != nil {
		fmt.Println(r, err)
	}
	fmt.Println(r)
}
