package mode

type Mode string

const (
	Min     Mode = "min"
	Max     Mode = "max"
	Sum     Mode = "sum"
	Avg     Mode = "avg"
	Median  Mode = "median"
	Default Mode = "_default"
)

func (mode Mode) String() string {
	return string(mode)
}
