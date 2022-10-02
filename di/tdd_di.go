package di

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// 参考：スタブ・モック - テスト駆動開発でGO言語を学びましょう
// https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

func TddDiMain() {
	// 標準出力に書き込む。
	Greet(os.Stdout, "山田")

	// バッファに書き込まれるだけで、標準出力には出ない。
	buffer := bytes.Buffer{}
	Greet(&buffer, "太郎")

	// HTTPサーバーを起動して、ハンドラーを利用してサーバーに書き込む。
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

// 以下のテストを書く。
// fmt.Printf()を実行するとstdoutに出力される。
// 依存関係を注入できるようにする。

// 以下はテストで使いたい
//func Greet(writer *bytes.Buffer, name string) {
//	fmt.Fprintf(writer, "Hello, %s", name)
//}

// 以下はメインアプリケーションで使いたい。
//func Greet(writer *os.File, name string) {
//	fmt.Fprintf(writer, "Hello, %s", name)
//}

// インターフェースを渡せばアプリでもテストでも使える汎用的な関数になる。
// インターフェースを引数に渡す = 依存性を注入することによってテスト可能な関数にする。
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// HTTPハンドラーを用意する。
// ハンドラー内でGreet関数を再利用できる。
// つまり、標準出力だけではなく、ファイルやサーバーにも挨拶ができる汎用的な関数になった。
// これが引数にインターフェースを渡すことのメリットである。
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// 依存性注入の利点
// ① 再利用性が高まる。今回の例だと標準出力だけではなく、サーバー、ファイル、バッファにも挨拶ができるようになった。
// ② テストが容易になる。Write()先を標準出力ではなくバッファを指定することによってテストを実現した。これはどちらも同じインターフェースを実装しているから可能になっている。
