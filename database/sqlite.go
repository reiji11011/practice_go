package database

type Sqlite struct {
	Host string
}

func (s Sqlite) ConnectDb() string {
	return s.Host + "に接続しました"
}
