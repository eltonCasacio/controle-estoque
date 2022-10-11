package shared

type Entidade interface {
	GetId() string
	IsValid() (bool, error)
}
