package dtos

type PaginateRequest struct {
	Filters FilterRequest `json:"filters"`
	Search  string        `json:"search"`
	Page    int           `json:"page"`
	PerPage int           `json:"perPage"`
}

type FilterRequest struct {
	IsActive *int   `json:"isActive" validate:"omitempty,oneOrZero"`
	Status   string `json:"status"`
}

func (p *PaginateRequest) SetDefaults() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PerPage == 0 {
		p.PerPage = 30
	}
}
