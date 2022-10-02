package Pointers

import "fmt"

func Test() {
	var a int
	// aのポインタ型を定義する
	fmt.Println(&a)
	// => 0xc0000b4008
	// ポインタ変数（ポインタ型が代入された変数）
	a1 := &a
	fmt.Println(a1)
	// => 0xc0000b4008
	fmt.Println(*a1)

	// ポインタ型を定義する
	// a2はint型のポイント型
	var a2 *int
	fmt.Println(a2)

	// 値渡し
	fmt.Println("値渡し")
	n := 1
	addValue(n)
	fmt.Println(n)

	// 参照渡し
	fmt.Println("参照渡し")
	n1 := 1
	addValue2(&n1)
	fmt.Println(n1)

	// 参照型の変数であるスライスやマップは、ポインタ型で渡さなくても元から参照渡しになっている
	s := []int{0}
	addSlice(s)
	fmt.Println(s)
}

func addValue(num int) int {
	return num + 1
}

func addValue2(num *int) {
	*num += 1
}

func addSlice(s []int) {
	s[0] = 1
}
