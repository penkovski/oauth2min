package mysql

type Storage struct{}

func New() *Storage {
	return &Storage{}
}
