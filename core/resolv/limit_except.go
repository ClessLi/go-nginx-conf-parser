package resolv

type LimitExcept struct {
	BasicContext
}

func NewLimitExcept(value string) *LimitExcept {
	return &LimitExcept{BasicContext{
		Name:     "limit_except",
		Value:    value,
		depth:    0,
		Children: nil,
	}}
}
