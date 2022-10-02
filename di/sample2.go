package di

import "fmt"

// 参考：https://www.soudegesu.com/post/go/test-with-interface/

type Client interface {
	Get() int
}

type Sample2 struct {
	client Client
}

// doGet関数でGet()をラップする。
// 外部パッケージからGet()を直接呼び出せないようにする。
func (s *Sample2) doGet() int {
	return s.client.Get()
}

// インターフェースを実装した構造体を用意する。
type clientImpl struct{}

func (c clientImpl) Get() int {
	return 1
}

func Sample2Main() {
	fmt.Println(PrintSample2())
	fmt.Println(PrintSample3())
	var c clientImpl
	fmt.Println(c.PrintSample4())
	var i Client = clientImpl{}
	fmt.Println(PrintSample5(i))
}

func PrintSample2() int {
	sample := &Sample2{client: &clientImpl{}}
	return sample.doGet()
}

// 関数内で構造体を初期化している。
// 構造体をモック化すればGet()のテストはできる。
func PrintSample3() int {
	client := clientImpl{}
	return client.Get()
}

func (c clientImpl) PrintSample4() int {
	return c.Get()
}

func PrintSample5(c Client) int {
	return c.Get()
}
