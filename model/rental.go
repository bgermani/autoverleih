package model

import (
	"time"
)

type Rental struct {
	ID             int       `json:"id" gorm:"primary_key"`
	AutoId         int       `json:"auto_id"`
	CustomerId     int       `json:"customer_id"`
	KilometerCount int       `json:"kilometer_count"`
	Start          time.Time `json:"period_start" sql:"type:timestamp without time zone"`
	End            time.Time `json:"period_end" sql:"type:timestamp without time zone"`
	CreatedAt      time.Time `json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	ModifiedAt     time.Time `json:"modified_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
