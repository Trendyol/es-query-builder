package operator

type Operator string

const (
	Or  Operator = "or"
	And Operator = "and"
)

func (operator Operator) String() string {
	return string(operator)
}
