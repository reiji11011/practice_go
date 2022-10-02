package _type

import "fmt"

func Test() {
	// インターフェース型
	var i interface{}
	fmt.Printf("%v, %T \n", i, i)

	// インターフェース型の変数には好きな値を入れることができる。
	i = true
	fmt.Printf("%v, %T \n", i, i)
	i = 1
	i = "123"
	fmt.Printf("%v, %T \n", i, i)

	// interface{}をキャストすると失敗する。
	//fmt.Println(string(i))
	// => cannot convert i (variable of type interface{}) to type string:
	// need type assertion

	// 型アサーション（抽象型から具象型への変換）
	fmt.Println(i.(string))
	// interface{}は今string型なので、uint32で型アサーションしようとするとpanicが発生する。
	//fmt.Println(i.(uint32))
	// => panic: interface conversion: interface {} is string, not uint32

	// 型アサーションは2つの変数で受け取る。
	i1 := interface{}(123)
	fmt.Printf("%v, %T \n", i1, i1)
	if v, ok := i1.(int); ok {
		fmt.Println(ok)
		fmt.Printf("%v, %T \n", v, v)
	}
}
