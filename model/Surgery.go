package model

import "time"

type Surgery struct {
	ID            int64     `json:"id"`
	DateCreated   time.Time `json:"dateCreated"`
	DateUpdated   time.Time `json:"dateUpdated"`
	CreatedBy     int64     `json:"createdBy"`
	UpdatedBy     int64     `json:"updatedBy"`
	Name          string    `json:"name"`
	Code          string    `json:"code"`
	CreatedByName string    `json:"createdByName"`
	UpdatedByName string    `json:"updatedByName"`
}
