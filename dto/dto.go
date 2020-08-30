package dto

type (
	LoginDto struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	CaptchaDto struct {
		Phone string `json:"phone"`
	}
)
