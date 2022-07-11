package service

type MapStorage interface {
	Insert(key int64, value any)
}
