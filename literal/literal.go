package literal

import (
	"fmt"
	"reflect"
)

func Test() string {
	// 文字列
	s := "str"
	fmt.Println(s)

	// 文字列を取り出す
	// byte情報を取り出す
	fmt.Println(s[0])
	for _, ru := range s {
		fmt.Println(ru)
	}

	// 文字列の長さを数える
	fmt.Println("文字列の長さを数える")
	t := "やa😄"
	// 1byteを単位とした長さを取得する
	// a: 1byte, や: 3byte, 😄: 4byte
	fmt.Println(len(t))
	fmt.Println(len([]rune(t)))

	// 文字列の特定の場所を取り出す。
	t = "やa😄"
	fmt.Println("文字列の特定の場所を取り出す")
	// 文字列を[]rune型にキャストする（byte単位ではなく、コードポイント単位にする）
	// さらにstring型にキャストする（文字の単位をコードポイントにし、string型にキャストする）
	fmt.Println(string([]rune(t)[0]))
	fmt.Println(string([]rune(t)[0:2]))

	// 文字列を結合する。
	fmt.Println("文字列を結合する")

	// +=で結合するやり方はパフォーマンスが悪い
	fmt.Println("+=で結合")
	text := "a"
	text += "b"
	fmt.Println(text)

	text2 := "+で結合"
	fmt.Println(text + text2)
	
	// 文字
	r := 'a'
	fmt.Println(r)
	// => 97

	// 数値
	fmt.Println(1)
	fmt.Println(1 + 1)
	fmt.Println(1 * 2)
	fmt.Println(1 * 2.1)
	fmt.Println(4 / 2.1)

	// 型を与えた後の演算（加減）
	num := 1
	fnum := 0.1
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(fnum))
	fmt.Println(num + int(fnum))
	fmt.Println(float64(num) + fnum)
	fmt.Println(int(fnum))

	// 掛け算
	a := 1
	b := 2.1
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(a * int(b))
	fmt.Println(float64(a) * b)

	// 割り算
	a = 5
	b = 1.5
	c := 2
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(a / c)
	fmt.Println(float64(a) / b)

	// int型同士の割り算の戻り値はint
	fmt.Println("int型同士の割り算の戻り値はint")
	int1 := 5.0
	int2 := 2.0
	fmt.Println(int1 / int2)
	fmt.Println(reflect.TypeOf(int1 / int2))

	// 型
	d := 0.1
	fmt.Println(reflect.TypeOf(d))
	d = 1
	fmt.Println(reflect.TypeOf(d))

	// 四捨五入
	// 切り上げ
	// 切り捨て
	// 浮動小数
	// null許容（null.int型）

	return ""
}
