package arrays

import "fmt"

func Test() string {
	// 宣言
	fmt.Println("配列の宣言")
	// ① var 変数名 [長さ]型
	var arr [1]string
	fmt.Println(arr)
	// ② 変数名 := [長さ]型{初期値, ...}
	ar := [3]string{"1"}
	fmt.Println(ar)

	// 初期を設定しない場合はゼロ値が設定される
	fmt.Println("初期を設定しない場合はゼロ値が設定される")
	var num = [3]int{}
	fmt.Println(num)

	fmt.Println("配列の上書き")
	ar[0] = "上書きした1"
	fmt.Println(ar[0])

	fmt.Println("配列の要素の上書き")
	ar[1] = "追加したindex1番目"
	fmt.Println(ar[1])

	// 配列の追加
	fmt.Println("配列の要素の追加")
	// 以下は上手くいかない
	//append(ar, "test")
	//fmt.Println(ar)

	// 配列の削除
	// 配列の削除は標準で用意されていないので自作が必要
	// ジェネリクスで作ってみたら面白そう

	// 新しいサイズで再代入をすることはできない。
	// => Cannot use '[4]string{}' (type [4]string) as the type [3]string
	//ar = [4]string{}
	//fmt.Println(ar)

	fmt.Println("空のスライスを定義する場合は:=ではなく、varを使うようにレコメンドされる")
	// Empty slice declaration using a literal
	//arr0 := [1]string{}
	//fmt.Println(arr0)

	// 配列から要素を取り出す
	var arr1 = [2]string{"a", "b"}
	fmt.Println(arr1)
	fmt.Println(arr1[0], arr1[1])
	fmt.Println(arr1[0:2])

	// 配列の中身の個数を調べる
	fmt.Println("配列の中身の個数を調べる")
	fmt.Println(len(arr1))

	fmt.Println("配列の範囲外はエラーが発生する")
	// invalid argument: array index 2 out of bounds [0:2]
	//fmt.Println(arr1[2])

	fmt.Println("配列の容量を超えて要素を定義するとエラーが発生する")
	// index 2 is out of bounds (>= 2)
	//var arr2 = [2]string{"a", "b", "c"}
	//fmt.Println(arr2)

	return ""
}
