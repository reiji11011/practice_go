package gorptest

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"

	// go get gopkg.in/gorp.v1
	"gopkg.in/gorp.v1"
)

func Test() {
	dbmap := initDb()
	fmt.Println(dbmap)
	// *sql.DBのCloseメソッドで処理終了後にコネクションを閉じるようにする。
	defer dbmap.Db.Close()
	var err error

	// TRUNCATE TABLEを実行して、既存の行を削除する。
	err = dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// 行を新規追加する。
	// DBに保存するデータを構造体で用意する。
	post1 := Post{Title: "Go入門", Body: "Goの入門本です", Created: time.Now().UnixNano()}
	post2 := Post{Title: "Java入門", Body: "Javaの入門本です", Created: time.Now().UnixNano()}

	// gorpの機能を使って新規保存する。（INSERT文を実行する）
	// 引数はポインタで用意をする。
	// 構造体のインスタンスにPKがなかった場合は自動でIDが付与される。
	err = dbmap.Insert(&post1)
	err = dbmap.Insert(&post2)
	checkErr(err, "Insert failed")

	// 1行のintデータを取得する（SELECT文）
	count, err := dbmap.SelectInt("select count(*) from posts")
	if err != nil {
		log.Fatalln(err, "SelectInt failed")
	}
	log.Println("Rows after inserting:", count)

	// プレースホルダーを渡して変数の中身によって取得する情報を変更する。
}

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`               // Column size set to 50
	Body    string `db:"article_body,size:1024"` // Set both column name and size
}

func initDb() *gorp.DbMap {
	// database/sql APIを使ってDBへのコネクションを開始する
	// 第一引数に使いたいDBのdriverを指定する。
	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
	checkErr(err, "sql.Open failed")

	// 接続を確認する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		fmt.Println("接続失敗")
	} else {
		fmt.Println("接続成功")
	}

	// DbMapはgorpのルート構造体
	// DBのdriverを第一引数に指定する。
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
