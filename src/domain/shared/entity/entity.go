package shared

type Entity interface {
	GetID() string
	IsValid() error
}
