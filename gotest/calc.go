package gotest

import (
	"fmt"
)

func Test() string {
	Calc(1, 2)
	Greet("Tanaka")
	return ""
}

// 最小でgoのテストを実行してみる。
func Calc(a, b uint32) uint32 {
	fmt.Println("go testingのテスト")
	return a + b
}

func Greet(name string) {
	fmt.Printf("Hello! %s", name)
}
