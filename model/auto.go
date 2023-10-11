package model

type Auto struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}
