package interfaces

type IHeep interface {
	Add(name string, body string) (err error)
	Get(name string) (body string, result interface{}, err error)
}
