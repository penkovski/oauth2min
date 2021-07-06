package storage

type Mysql struct{}

func NewMysql() *Mysql {
	return &Mysql{}
}
