package _interface

type Text struct{}

func (t Text) GetText() string {
	// 外部APIからテキストを取得する。
	// その戻り値を変数に詰めて返却をする。
	return ""
}

func MockTest() string {
	var s string
	var t Text
	// 外部APIからテキストを取得する。
	s = t.GetText()
	// そのテキストに処理を加えて返却をする。
	s = s[0:1]
	return s
}

// stringを渡したら適当な文字列を返却してくれる。
func MockGetText() string {
	return "aiueo"
}
