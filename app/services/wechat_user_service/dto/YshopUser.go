package dto

type YshopUser struct {
	Id       int64  `json:"id"`
	RealName string `json:"real_name"`
	Mark     string `json:"mark"`
	Phone    string `json:"phone"`
	Integral int    `json:"integral"`
}
