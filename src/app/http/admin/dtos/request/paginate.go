package dtos

type PaginateRequest struct {
	Filters []interface{} `json:"filters"`
	Search  string        `json:"search"`
	Page    int           `json:"page"`
	PerPage int           `json:"perPage"`
}

func (p *PaginateRequest) SetDefaults() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PerPage == 0 {
		p.PerPage = 30
	}
}
