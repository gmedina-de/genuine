package repository

type Repository interface {
	Migrate(model interface{})

	Create(model interface{})
	RetrieveAll(model interface{})
	RetrieveFirst(model interface{}, conditions ...interface{})
	Update(model interface{}, column string, value interface{})
	Delete(model interface{})
}
