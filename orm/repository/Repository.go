package repository

type Repository interface {
	RetrieveAll(model interface{})
}
