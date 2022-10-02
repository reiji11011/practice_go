package slices

import "fmt"

func Test() string {
	// sliceのappend
	n1 := []int32{1, 2, 3}
	n2 := []int32{4, 5, 6}
	fmt.Println(n1, n2)
	n3 := append(n1, 1)
	fmt.Println(n3)

	// []int32 スライスの中身は int32であるため、[]int32（スライス）は追加できない。
	// 追加するためにはスライスの中身を展開する必要がある。
	// => Cannot use 'n2' (type []int32) as the type int32
	//n4 := append(n1, n2)
	//fmt.Println(n4)

	// スライスを展開して可変長引数として渡す。
	// unpack演算子を使った展開をする。
	n5 := append(n1, n2...)
	fmt.Println(n5)

	// 可変長引数の挙動を確認する
	// 以下2つは同じ挙動
	TestPackSlice(n1)
	TestPackLiteral(n1...)

	return ""
}

func TestPackSlice(n []int32) {
	fmt.Println(n)
	fmt.Println(n[0])
}

func TestPackLiteral(n ...int32) {
	fmt.Println(n)
	fmt.Println(n[0])
}
