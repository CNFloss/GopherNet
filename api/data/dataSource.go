package data

type Entity interface {
	GetID() int
	Convert(map[string]interface{}) (Entity, error)
}

type DataSource interface {
	AddEntity(Entity) error
	GetByID(id int) (Entity, error)
	GetAll() []Entity
	Init(string, Entity) error
}
