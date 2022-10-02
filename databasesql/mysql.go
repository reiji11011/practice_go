package databasesql

import (
	"database/sql"
	"fmt"

	// 接続するDBのドライバーを用意する
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Test() string {
	// 第一引数：ドライバー名
	// 第二引数：ドライバーがデータへアクセスする方法
	//  local MySQL server instanceの“hello” databaseへ接続するには以下のように記述する
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	db, err := sql.Open("mysql",
		"user:pass@tcp(127.0.0.1:7306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	// 接続が終了したらCloseする
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
		return ""
	} else {
		fmt.Println("データベース接続成功")
	}

	var (
		id   int
		name string
	)

	// Fetching Data from the Database
	rows, err := db.Query("select id, name from users where id = ?", 2)
	fmt.Println(rows)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// クエリで取得したカラムを変数に割り当てる。
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return ""
}
