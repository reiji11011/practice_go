package linter

import "fmt"

func Test() string {
	// 未達
	a := 1
	if a > 0 {
		return ""
		fmt.Println("到達")
	}
	// => unreachable code

	// fmt.Sprintf等の不正な使用
	fmt.Sprintf("%s", 1)
	// => fmt.Sprintf format %s has arg 1 of wrong type int
	// =>  result of fmt.Sprintf call not used

	return ""
}
