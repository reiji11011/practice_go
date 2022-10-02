package _interface

import "fmt"

type Walker interface {
	Walk()
}

type Baby struct {
	Name string
}
type Cat struct {
	Name string
}

type Duck struct {
}

// 構造体にインターフェースを埋め込む
type Dog struct {
	Name string
	Walker
}

// 以下のインターフェースを満たすものはWalkerとRunnerを同時に満たすことになる
type Runner interface {
	Run()
}

type Human interface {
	Walker
	Runner
}

func (b Baby) Walk() {
	fmt.Println("赤ちゃんが歩きます。")
}

func (c Cat) Walk() {
	fmt.Println("赤ちゃんが歩きます。")
}

func Test() {
	// インターフェースを型として与えた変数には、どの型のインスタンスかに関係なく、そのインターフェースを満たす型を格納することができる。

	// Baby型の構造体
	baby := Baby{Name: "太郎"}
	// Cat型の構造体
	cat := Cat{Name: "猫太郎"}
	// Walkerインターフェース型のスライスとして上記を変数に代入する。
	walker := []Walker{baby, cat}
	fmt.Println(walker)
	// => [{太郎} {猫太郎}]

	// インターフェース値
	// 動的な型情報と値情報を持った構造体
	// 実装クラスインスタンス化したものを入れることができる構造体
	var i interface{}
	i = 123
	fmt.Printf("%T, %v \n", i, i)
	// type=int, value=123

	// インターフェースの実体には構造体のポインタが入る。

	// 構造体をインターフェースに代入した場合は、内部でコピーされてその参照であるポインタをインターフェース値が保持している。
	cat2 := Cat{Name: "猫田中"}
	var i2 interface{}
	i2 = cat2
	fmt.Printf("%T: %+v\n", i2, i2)

	// 元の構造体を書き換える
	cat2.Name = "猫小林"
	// インターフェース値には構造体のコピーが入っているので、変わらない。
	fmt.Println(cat2)
	fmt.Printf("%T: %+v\n", i2, i2)

	// ポインタのアドレスを代入すると書き換えることができる。
	cat3 := Cat{Name: "猫3"}
	var i3 interface{} = &cat3
	fmt.Printf("%T: %+v\n", &cat3, &cat3)
	fmt.Printf("%T: %+v\n", i3, i3)

	// 同じ構造体のポインタアドレスを保持しているため、インターフェース値でも値が変更されている。
	// 参考：https://tyablog.net/2020/01/05/difference-between-interface-and-pointer-in-golang/
	cat3.Name = "猫3を書き換える"
	fmt.Println(cat3)
	fmt.Printf("%T: %+v\n", i3, i3)

	// ポインタ
	n := 123
	fmt.Println(n)
	fmt.Println(&n)

	// インターフェースを変数に代入する。
	var walker2 Walker
	walker2 = Baby{Name: "次郎"}
	fmt.Println(walker2)
	walker2.Walk()

	walker3 := &Baby{Name: "三郎"}
	fmt.Println(walker3)
	walker3.Walk()

	// 構造体がインターフェースを満たしているかを確認する。
	// ブランク識別子（_）によるダミー変数宣言を使って構造体がインターフェースを満たしていることを確認する。
	var _ Walker = &Cat{}
	var _ Walker = &Dog{}
	// 以下はインターフェースを満たしていないのでエラーが表示される。
	// Cannot use '&Duck{}' (type *Duck) as the type Walker Type does not implement 'Walker' as some methods are missing: Walk()
	//var _ Walker = &Duck{}

	// 構造体のフィールドをインターフェースにすることによって、モックで差し替えるようにする。

}
