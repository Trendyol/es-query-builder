package order

type Order string

const (
	Asc     Order = "asc"
	Desc    Order = "desc"
	Default Order = "_default"
)

func (order Order) String() string {
	return string(order)
}
