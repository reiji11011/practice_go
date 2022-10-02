package database

type Database interface {
	ConnectDb()
}

type Mysql struct {
	Host string
}

func (m Mysql) ConnectDb() string {
	return m.Host + "に接続しました"
}
