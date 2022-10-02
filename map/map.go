package _map

import (
	"fmt"
)

type Children struct {
	Character map[string]string
}

func Test() string {
	fmt.Println("")

	// mapの初期化
	// var
	var m1 map[int32]string
	fmt.Println(m1)
	fmt.Println(m1[0])
	fmt.Println(m1[1])

	// := （初期値を設定しないとpanic）
	//m8 := map[int32]string
	//fmt.Println(m8)
	// =>  map[int32]string (type) is not an expression

	// make
	m3 := make(map[int]bool)
	fmt.Println(m3)

	// varでは初期値を設定できない。
	//var m2 map[uint8]uint8{
	//	1: 1,
	//	2: 2,
	//}
	//fmt.Println(m2)

	m4 := map[uint8]uint8{
		1: 1,
		2: 2,
	}
	fmt.Println(m4)
	fmt.Println(m4[1])
	// エラーは発生せず、ゼロ値が出力される。
	fmt.Println(m4[3])
	// keyの型が違う場合はpanic
	//fmt.Println(m4[true])
	// => cannot use true (untyped bool constant) as uint8 value in map index

	// スライスをキーにする
	// スライスは無効なキー
	//m5 := map[[]int32][]string{
	//	[1, 2]: ["1", "2"],
	//}
	//fmt.Println(m5)
	// => Invalid map key type: comparison operators == and != must be fully defined for the key type

	// マップをキーにする。
	// マップは無効のキー
	//var m6 map[int32]int32
	//var m7 map[m6]int32
	//fmt.Println(m7)
	// =>  m6 (variable of type map[int32]int32) is not a type

	// 配列をキーにする。
	var m9 map[[3]int]string
	fmt.Println(m9)

	// スライスはキーにできない。
	//var m10 map[[]int]string
	//fmt.Println(m10)

	type Number struct {
		init int8
	}

	// 構造体をキーにする
	m11 := map[Number]Number{}
	fmt.Println(m11)

	// 初期化
	_ = Children{
		//Character: {"主観": "論理的"},
		Character: map[string]string{"主観": "論理的"},
	}

	// ファイルへマップの吐き出し
	fmt.Println("mapのテスト")
	a := map[string]string{"a": "1"}
	//f, _ := os.Create("test.txt")
	println(a)
	var a1 map[string]string
	println(a1)

	m := make(map[string]int)

	fmt.Println(m)

	m["key1"] = 10
	m["key2"] = 20

	fmt.Println(m)

	//m1 := map[string]int{key1: 10, key2:20}

	// makeで初期化するときに容量（キャパシティ）を指定する
	// マップリテラルで初期化をする
	// 最適な初期化方法は？
	// 宣言と初期化
	// 初期化せずにnil（ゼロ値）で値をセットする。
	// キーの追加・変更・削除
	// マップ自体の削除
	// キーの個数・値の個数
	// キーの存在を確認
	// キーのみ取得
	// 値のみ取得
	// mapのループ
	// mapのソート
	// mapの比較

	return ""
}
