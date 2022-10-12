package shared

type Entity interface {
	GetId() string
	IsValid() (bool, error)
}
