package object

import (
	"fmt"
)

type Country struct {
	Name string
	Tax  float64
}

func Polymorphism() {
	japan := Country{Name: "日本", Tax: 0.08}
	india := Country{Name: "インド", Tax: 1.18}
	china := Country{Name: "中国", Tax: 0.06}

	countries := []Country{japan, india, china}
	sumPrice(1000.00, countries)

	Polymorphism2()
}

func sumPrice(price float64, countries []Country) float64 {
	var sum float64
	for _, country := range countries {
		// 国によって合計方法が変わる。
		switch country.Name {
		case "日本":
			sum += price * country.Tax
		case "インド":
			sum += price
		case "中国":
			sum += price*country.Tax + 2000.00
		default:
			sum += price
		}
	}
	fmt.Println(sum)
	return sum
}

// 上記をポリモーフィズムを使って解決をする。
// インターフェースを実装する構造体によって上記をリファクタリングする。

type Calculate interface {
	Sum(price float64) float64
}

// 類似したクラスに対する振る舞いを共通化する仕組み

type Japan struct {
	Name string
	Tax  float64
}

type India struct {
	Name string
	Tax  float64
}

type China struct {
	Name string
	Tax  float64
}

// 条件分岐は減ったが構造体の定義は増えた。
// 同じインターフェースを実装していても、構造体によって挙動が違う。

func (j Japan) Sum(price float64) float64 {
	return price * j.Tax
}

func (i India) Sum(price float64) float64 {
	return price
}

func (c China) Sum(price float64) float64 {
	return price*c.Tax + 2000.00
}

func Polymorphism2() {
	japan2 := Japan{Name: "日本", Tax: 0.08}
	india2 := India{Name: "インド", Tax: 1.18}
	china2 := China{Name: "中国", Tax: 0.06}

	countries2 := []Calculate{japan2, india2, china2}
	sumPrice2(1000.00, countries2)
}

// ポリモーフィズムを使ったことにより、メソッドの呼び出し側がスッキリと記述できる。
// 新しい構造体　（例えばAmerica）が追加されてもsumPrice2()は変更しなくて良い。
// →その場合、既存のコードに手を加えず、新規のコードを追加するだけで対応できているので、オープンクローズドの原則を満たしている。
// ソフトウェアの構成要素（クラス、モジュール、関数など）は拡張のために開いていて、修正のために閉じていなければならない。

func sumPrice2(price float64, countries []Calculate) float64 {
	var sum float64
	for _, country := range countries {
		sum += country.Sum(price)
	}
	fmt.Println(sum)
	return sum
}

// 参考
// Go言語でポリモーフィズムとオープンクローズドの原則とそれらの便利さを理解する
// https://qiita.com/__m/items/cc45791c7cc2ba7fac50
