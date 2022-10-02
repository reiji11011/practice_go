package loop

import "fmt"

func Test() string {
	// 空の配列に対してrangeを実行する。
	var empty []string
	fmt.Println(empty)
	//empty = []string{"a", "b"}
	//fmt.Println(empty)

	for _, v := range empty {
		fmt.Println(v)
		fmt.Println("表示")
	}

	// 反復処理をスキップする（loop）
	num := []int{1, 2, 3, 4, 5}
	for i, n := range num {
		// 条件に一致する場合は、それ以降の処理をスキップして次のループに進む。
		if i == 2 {
			continue
		}
		fmt.Println(i, n-1)
	}
	return ""
}
