package nginx

type Http struct {
	BasicContext `json:"http"`
}

func (h *Http) QueryAll(kw KeyWords) (parsers []Parser) {
	if h.filter(kw) {
		parsers = append(parsers, h)
	}
	return h.subQueryAll(parsers, kw)
}

func (h *Http) Query(kw KeyWords) (parser Parser) {
	if h.filter(kw) {
		parser = h
	}
	return h.subQuery(kw)
}

func (h *Http) BitSize(_ Order, _ int) byte {
	return 0
}

func (h *Http) BitLen(_ Order) int {
	return 0
}

func (h *Http) Size(_ Order) int {
	return 0
}

func NewHttp() *Http {
	return &Http{BasicContext{
		Name:     TypeHttp,
		Value:    "",
		Children: nil,
	}}
}