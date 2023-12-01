package repository

type Repository[T interface{}, U interface{}, V interface{}] interface {
	FindOne(id string) (T, error)
	FindAll() ([]T, error)
	InsertOne(data U) error
	UpdateOne(id string, data V) (T, error)
	DeleteOne(id string) (T, error)
}
