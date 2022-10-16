package shared_entity

type Entity interface {
	GetID() string
	IsValid() error
}
