package sort

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  uint32
}

func Test() string {
	// 比較する

	// 文字列（比較可能）
	s1 := "12"
	s2 := "123"
	fmt.Println(s1 < s2)

	// 数値（同じ型）
	var n1 int = 1
	var n2 int = 2
	fmt.Println(n1 < n2)

	// 数値（違う型）
	// var n3 uint8 = 3
	// fmt.Println(n1 < n3)
	// => Invalid operation: n1 < n3 (mismatched types int and uint8)

	// map
	//m1 := map[int32]{string}
	//m1[1] = "1"
	//fmt.Println(m1)

	// 構造体を比較する
	// 構造体を並び替える
	// 安定ソートを保証しない
	people := []Person{
		{Name: "A", Age: 1},
		{Name: "B", Age: 2},
		{Name: "C", Age: 3},
		{Name: "D", Age: 4},
		{Name: "E", Age: 1},
		{Name: "F", Age: 2},
		{Name: "G", Age: 3},
		{Name: "H", Age: 4},
		{Name: "I", Age: 1},
		{Name: "J", Age: 2},
		{Name: "K", Age: 3},
		{Name: "L", Age: 4},
	}
	// Slice
	//sort.Slice(people, func(i, j int) bool {
	//	return people[i].Age < people[j].Age
	//})
	//fmt.Println("sort.Slice；安定ソートを保証しない")
	//fmt.Println(people)

	// SliceStable
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("sort.SliceStable；安定ソートを保証する")
	fmt.Println(people)

	return ""
}
