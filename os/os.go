package os

import (
	"fmt"
	"os"
)

func Test() string {
	// 実行場所からのパスを指定する。
	// f: *File（os.File型のファイルオブジェクト）
	// 構造体も型の1つであり、型に対してメソッドがある。
	// そのため、os.File型に定義されているメソッドを利用してファイル操作ができる
	// 読み取り専用でファイルを開く
	f, err := os.Open("./os/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	// 書き込み権限でファイルを開く
	f2, err := os.Create("./os/test2.txt")
	if err != nil {
		fmt.Println(f2, err)
	}
	fmt.Println(f2)

	// ファイルを読み込むためのbyteのスライスを用意する
	// byteはuint8のエイリアス
	// bytesの大きさをintで返却する
	data := make([]byte, 1024)
	fmt.Printf("data: %d \n", data)

	// []byteにファイルの内容を読み込む
	// 返り値はbyteの数, 何byteが読みこかれたか
	// io.Readerインターフェースを実装している
	// os.File型はio.Reader型を満たしており、io.Reader型として扱うことができる
	r, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("errが発生しました")
	}
	// スライスの一部分だけを取得する
	fmt.Println("スライスの一部分だけを取得する")
	fmt.Println(data[r])
	fmt.Println(string(data[r]))

	// 文字列にエンコードする前のバイト列を出力
	// Slice[:] 先頭から最後尾まで
	fmt.Println("文字列にエンコードする前のバイト列を出力")
	fmt.Println(r)                // => 18
	fmt.Println(data)             // => [227 131 134 227 130 185 227 131 136 227 129 167 227 129 153 227 128 130 0 0 0 0 0 0 0 ...]
	fmt.Println(data[:r])         // => [227 131 134 227 130 185 227 131 136 227 129 167 227 129 153 227 128 130]
	fmt.Println(string(data[:r])) // => テストです。

	// 書き込みの下準備
	s := "こんにちは"
	b := []byte(s)
	fmt.Println(s)
	fmt.Println(b)

	f3, err := os.Open("./os/test2.txt")
	f4, err := os.Create("./os/test3.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 読み込み権限のファイルに書き込むため以下は失敗する。
	_, err = f3.Write(b)
	if err != nil {
		// エラーは書き出されるが、処理は継続できる。
		fmt.Println(err)
	}

	_, err = f4.Write(b)
	if err != nil {
		fmt.Println(err)
	}

	return ""
}
