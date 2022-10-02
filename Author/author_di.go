package Author

import (
	"fmt"
	"myapp/database"
)

func (a Author) ConnectAuthorDbDi(mysql database.Mysql) string {
	fmt.Println(mysql)
	return "ok"
}
