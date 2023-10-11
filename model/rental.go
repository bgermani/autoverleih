package model

import (
	"gorm.io/datatypes"
)

type Rental struct {
	ID             int            `json:"id" gorm:"primary_key"`
	AutoId         int            `json:"auto_id"`
	CustomerId     int            `json:"customer_id"`
	KilometerCount int            `json:"kilometer_count"`
	Start          datatypes.Date `json:"period_start"`
	End            datatypes.Date `json:"period_end" `
	CreatedAt      datatypes.Date `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	ModifiedAt     datatypes.Date `json:"modified_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
