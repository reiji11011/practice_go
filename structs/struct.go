package structs

import "fmt"

// 構造体に構造体を埋め込む
type Machine struct {
	//car     Car
	//bicycle Bicycle
	vehicle Vehicle
}

type Vehicle interface {
	Warning()
}

type Car struct{}

type Bicycle struct{}

type String string

type Animal struct {
	Name
	Weight uint32
}

type Animal2 struct {
	*Name
}

type Name struct {
	NickName string
}

type Person struct {
	Study
}

type Study struct {
	level int64
}

type Cat struct {
	Animal
}

func (s Study) Read() {
	fmt.Println("何かを読む")
}

func (a Animal) PrintName() {
	fmt.Println("犬")
}

func (a2 Animal2) PrintName() {
	fmt.Println("犬2")
}

func (a Animal) Cry(voice string) {
	fmt.Println(voice)
}

func (c Car) Warning() {
	fmt.Println("プップー！")
}

func (b Bicycle) Warning() {
	fmt.Println("チリリン！")
}

func (s String) Warning() {
	fmt.Println("String!")
}

func Test() {
	//m := Machine{}

	// 構造体を経由しないと呼び出すことはできない。
	//m.Warning()
	//  m.Warning undefined (type Machine has no field or method Warning)
	// 埋め込んだ構造体経由でメソッドにアクセスすることができる。
	//m.car.Warning()
	//m.bicycle.Warning()

	// インターフェースを埋め込んでみる
	car := Car{}
	// インターフェースを実装した構造体（？）なら埋め込んで使うことができる。
	m1 := Machine{vehicle: car}
	m1.vehicle.Warning()

	// インターフェースを実装している型なら何でも渡すことができる。
	// メソッドを実装した型を渡すことができるのでテストがしやすくなる？
	bike := Bicycle{}
	m2 := Machine{vehicle: bike}
	m2.vehicle.Warning()

	// 構造体ではなく型であればOK
	var s String
	m3 := Machine{vehicle: s}
	m3.vehicle.Warning()

	// 埋めこんだ構造体のメソッドを使える。
	a := Animal{}
	a.PrintName()

	a2 := Animal2{}
	a2.PrintName()

	// 埋め込まれている構造体のメソッドを使えるわけではない。
	// 継承の場合、親のフィールドやメソッドは使えるが、構造体の場合はそうではない。
	//name := Name{}
	//name.PrintName()

	// 継承っぽいことをする。
	cat := Cat{}
	cat.Cry("にゃー")

}
