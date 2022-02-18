package base

type Model interface {
	Name() string
	GetDB() Database
}
