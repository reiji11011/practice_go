package main

import (
	"myapp/generics"
)

type Test struct {
	Title string
	Page  int
}

func main() {
	//author := Author.Author{Name: "田中太郎"}
	//result := author.ConnectAuthorDb()
	//fmt.Println(result)
	//result = author.ConnectAuthorDbDi()
	//fmt.Println(result)
	//test := NewTest("Go本", 10)
	//fmt.Println(test)
	//test2 := NewTest("Go本", 10)
	//fmt.Println(*test2)
	// 構造体を試す
	//structs := structs.Structs{}
	//structs.Test()
	//structs.Embedded()
	//_ = literal.Test()
	//_ = arrays.Test()
	//_ = slices.Test()
	//_ = databasesql.Test()
	//gorptest.Test()
	//_ = volatiletech.Test()
	//_ = loop.Test()
	//_ = sort.Test()
	//_ = _map.Test()
	//_ = gotest.Test()
	//_ = testify.Test()
	//_ = control.Test()
	//_ = linter.Test()
	//_ = os.Test()
	//io.Test()
	//function.Test()
	//bytes.Test()
	//err.Test()
	//_type.Test()
	//json.Test()
	//http.Test()
	//defergo.Test()
	//object.Isa()
	//object.Polymorphism()
	//object.DuckTyping()
	//object.Inheritance()
	//_interface.Test()
	//_interface.MockTest()
	//di.Main()
	// わからなくなったからサンプルコードに学ぶ
	//di.Sample()
	//di.Sample2Main()
	//di.TddDiMain()
	//di.TddMockMain()
	//swagger.Test()
	//env.Test()
	//strings.Test()
	//Pointers.Test()
	generics.Test()
}

//func (i int)Test1()  {
//	fmt.Println("組み込み型はメソッドを持てない")
//}

//type integer int
//func (i integer)Test1()  {
//	fmt.Println("組み込み型はメソッドを持てない")
//}
