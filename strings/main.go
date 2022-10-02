package strings

import (
	"fmt"
	"strings"
)

func Test() {
	fmt.Println("strings")

	// Split
	// 文字列を指定した区切り文字でsliceに変換をする。
	s := "a,b,c,d"
	fmt.Println(s)
	r := strings.Split(s, ",")
	fmt.Println(r)
	for i, v := range r {
		fmt.Println(i, v)
	}
}
