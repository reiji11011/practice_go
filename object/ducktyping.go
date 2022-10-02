package object

import "fmt"

type Animal interface {
	Cry()
}

type Duck struct{ Name string }
type Cat struct{ Name string }

func (d Duck) Cry() {
	fmt.Println(d.Name + "がガー！ガー！")
}

func (c Cat) Cry() {
	fmt.Println("にゃー！にゃー！")
}

func DuckTyping() {
	duck := Duck{Name: "アヒル"}
	cat := Cat{}
	duck.Cry()
	cat.Cry()

	// インターフェースを実装していれば、違う型（Duck型, Cat型）を代入することができる。
	var animal Animal
	animal = duck
	animal.Cry()
	animal = cat
	animal.Cry()
}
