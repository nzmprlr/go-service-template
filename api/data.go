package api

type FooData interface {
	Incr() (int, error)
}
