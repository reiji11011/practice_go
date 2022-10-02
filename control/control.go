package control

import "fmt"

func add(a, b uint32) uint32 {
	return a + b
}

func Test() string {
	// if
	a := 2
	if a == 1 {
		fmt.Println(a)
	} else {
		fmt.Printf("%d は1じゃありません\n", a)
	}

	// mapのkeyがあるかを判別する
	// if 簡易文; 条件式 { 条件式がtrueだった場合の処理 }
	// 簡易文：変数の定義、代入など
	a1 := map[string]string{"1": "1"}
	fmt.Println(a1)
	fmt.Println(a1["1"])
	if value, ok := a1["1"]; ok {
		fmt.Printf("mapのkey（%s）は存在します\n", value)
	}

	// 関数の返り値を変数に代入し、条件式で使う。
	if value := add(1, 20); value > 20 {
		fmt.Println("20より大きいです。")
	}

	isBook := false
	if isBook == true {
		fmt.Println("trueです")
	}
	if isBook {
		fmt.Println("trueです")
	}
	if !isBook {
		fmt.Println("falseです")
	}

	// for
	fmt.Println("for")
	s := []int32{1, 2, 3, 4, 5}
	fmt.Println(len(s))

	for i := 0; i < 5; i++ {
		fmt.Println("a")
	}

	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// range
	for i, v := range s {
		fmt.Println(i+1, v)
	}

	// switch
	status := "300"

	switch status {
	case "100":
		fmt.Println("100です")
	case "200":
		fmt.Println("200です")
	default:
		fmt.Println("statusがありません")
	}
	return ""
}
