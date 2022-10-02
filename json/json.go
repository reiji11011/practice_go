package json

import (
	"encoding/json"
	"fmt"
)

type GoStruct struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Test() {
	// jsonを読み込む
	var data GoStruct
	jsonData := `{"id":"1","name":"田中"}`
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", jsonData)
	// => {"id":"1","name":"田中"}

	// 構造体へマッピングする
	// ネストしたjsonを構造体へマッピングする
}
