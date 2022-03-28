package stack

type Stack interface {
	Push(value interface{})
	Pop() (interface{}, error)
}
