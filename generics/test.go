package generics

import "fmt"

func Test() {
	fmt.Println(AddOne(1))
	fmt.Println(AddTwo(1))

	r := AddThree(1)
	fmt.Printf("%T, %v", r, r)
}

// 引数に1を加える関数を作成する
// 無効な演算: v + 1 (interface{} と untyped int の型が一致しません)
//func AddOne(v interface{}) interface{} {
//	return v + 1
//}

// 引数に1を加える関数
// インターフェースをキャストして利用する
// int以外でも使えるようになりたいな...。
func AddOne(v interface{}) interface{} {
	return v.(int) + 1
}

// 引数に2を加える関数を作成する
// switch文と.(type)を利用して複数の型で使える汎用的な関数を作成する
func AddTwo(v interface{}) interface{} {
	switch t := v.(type) {
	// 以下caseに演算したい型を全て書き込む
	case int:
		return t + 2
	}
	// 演算が不可能な型が渡ってきた場合はpanicを発生させる
	panic("unreachable")
}

func AddThree(v interface{}) interface{} {
	switch t := v.(type) {
	// 問題点①：型の列挙が大変
	case int, int16, int32, int64, float32, float64:
		//return t + 3
		return t
	}
	// 問題点②：未対応の型に気がつけない。
	// 例えば、float16型を渡すとpanicが発生してしまう。
	// コンパイルは通るが、実行時に失敗する。
	panic("unreachable")
}

// ジェネリクスを使って上記を書き換える。
func AddOneByG[T any](v T) {

}

// Type Parameters（型パラメーター）
// Type sets
// Constraints package
