package json

// エディタの自動変換を試す。
// jsonを貼り付けると、Golandの機能で構造体へと変換をしてくれる。
// The inserted text seems to be JSON. Do you want to generate a Go type from it?
// Generate Go Type from JSON
// befor: {"id":"1","name":"田中"}
type T struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
