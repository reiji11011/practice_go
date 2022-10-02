package Author

import (
	"myapp/database"
)

type Author struct {
	Name string
}

func (a Author) ConnectAuthorDb() string {
	// AuthorパッケージからDatabaseパッケージを呼び出している（依存している）
	// ConnectDB()はMysql{}構造体に依存しているため、Mysqlの変更やMysqlから他DBへの変更が難しい。
	db := database.Mysql{Host: "127.0.0.1"}
	return db.Host + "ok"
}
