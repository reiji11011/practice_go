package http

import (
	"fmt"
	"net/http"
)

type Serve1 struct{}

func (s Serve1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serve1")
}

type str1 string

func (s str1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("str1")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func Test() {

	// サーバーを起動する。
	// 第一引数 addr
	// "host:port" 形式で記述する
	// 空（""）の場合は、":http" (port 80)

	// 第二引数 handler
	// Handlerインターフェース
	// ServeHTTP()を実装した型なら引数として使うことができる。
	// nilの場合は、DefaultServeMuxが使われ、アクセスには404が返却される

	//http.ListenAndServe("localhost:8081", nil)

	// 以下のように複数サーバーは同時に起動ができない。
	//http.ListenAndServe("localhost:8082", nil)

	// 以下は実行されない。
	//fmt.Println("疎通テスト")

	// Handler interfaceを満たす型を用意する。
	// 構造体Serve1型を用意して引数に渡して実行をするとターミナルに任意の文字列が出力される。
	//s1 := Serve1{}
	//http.ListenAndServe("localhost:8083", s1)

	// string型のstr1を引数として渡す。
	// インターフェースの型を満たすので処理は成功する。
	//var s2 str1
	//http.ListenAndServe("localhost:8083", s2)

	// 上記のやり方だとリクエストに対して同じ処理しか返すことができない。
	// 常にServeHTTPメソッドが実行され単一の処理しか返すことができない。
	// → だからマルチプレクサを使って複数のリクエストに対して複数の処理を返すようにする

	// マルチプレクサを利用して複数のリクエストを捌けるようにする。
	mux := http.NewServeMux()
	mux.HandleFunc("", hello)
}
