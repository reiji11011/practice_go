package structs

import "fmt"

type Human interface {
	Walk()
}

// * Goでクラスという言葉は相応しくないが、理解しやすくするためにコメントを振る。

// 親
type Teen struct {
	Age uint32
}

// 子
// Teenを継承しているように動かしてみる。
type Child struct {
	Name string
	Teen // embedded field
}

func (t Teen) Walk() {
	fmt.Println("二足歩行で歩く")
}

func (t Teen) Run() {
	fmt.Println("タッタッタ")
}

func (c Child) Walk() {
	fmt.Println("四足歩行で歩く")
}
func (c Child) Cry() {
	fmt.Println("オギャー")
}

func Embedded() {
	teen := Teen{Age: 16}
	child := Child{Name: "太郎", Teen: teen}

	// 親の挙動
	// 子のフィールドにはアクセスできない。
	//fmt.Println(teen.Name)
	// 子のメソッドにはアクセスできない。
	//teen.Cry()

	// 子の挙動
	// 親のフィールドにアクセスができる。
	fmt.Println(child.Age)
	// 親のメソッドに明示的にアクセスができる。
	child.Teen.Run()
	// 以下のようにも書ける。
	child.Run()

	// インターフェースを試す。
	child.Walk()
	child.Teen.Walk()

	// 継承ではないので親の型に子の型を代入することはできない。
	//t := Teen{}
	//c := Child{}
	//t = c
	// => Cannot use 'c' (type Child) as the type Teen
}
