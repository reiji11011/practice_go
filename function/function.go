package function

import "fmt"

type number int

func (n number) PrintNumber() {
	fmt.Println(1)
}

// どんな型でも引数として渡すことができる関数
func PrintAny(any interface{}) {
	fmt.Println(any)
}

func ManyNum(num ...int) {
	fmt.Println(num)
}

func ManyNum2(num []int) {
	fmt.Println(num)
}

//func PrintStr() func() {
//	fmt.Println("関数")
//}

//func ReturnFunc() func() {
//	return PrintStr()
//}

func Test() string {
	// private関数の呼び出し
	private()

	// 型を作成し、紐づいた関数を実行する
	var num number
	num.PrintNumber()

	// 引数にstrでもintでも何でも受け取れて、その内容を出力する関数
	PrintAny(2)
	PrintAny("文字列")
	PrintAny(map[string]string{"1": "1"})
	PrintAny([]int{1, 2, 3})

	// 引数の数が1~N個に対応できる関数
	n1 := []int{1, 2, 3, 4, 5}
	ManyNum(n1...)
	ManyNum2(n1)

	// 関数名がない関数
	// 関数を値として使って、変数へ代入をする。
	f := func(text string) string { return text + "を受け取り実行" }
	fmt.Println(f("引数"))

	// 高層関数
	// 戻り値として関数を返す関数
	//ReturnFunc()

	// 関数を引数に受け取る関数

	// 変数の中身を書き換える関数

	// 変数の中身を書き換えない関数

	return ""
}

func private() {
	fmt.Println("private関数です。引数なし、返り値型なし")
}
