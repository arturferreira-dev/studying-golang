package generics

type Paginate struct {
	Page       int
	TotalItems int
	TotalPages int
	Result     []interface{}
}

func (p Paginate) GetResult() Paginate {
	return p
}
