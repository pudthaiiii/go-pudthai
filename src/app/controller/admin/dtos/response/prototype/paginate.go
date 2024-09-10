package dtos

type UsageItemResponse struct {
	ID          int    `json:"id"`
	OrgCode     string `json:"orgCode"`
	ProductCode string `json:"productCode"`
}
