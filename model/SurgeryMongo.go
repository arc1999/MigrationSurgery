package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SurgeryMongo struct {
	ID            int64     `json:"_id" bson:"_id"`
	DateCreated   time.Time `json:"dateCreated"`
	DateUpdated   time.Time `json:"dateUpdated"`
	Slug 		primitive.Binary	`json:"slug"`
	CreatedBy     int64     `json:"createdBy"`
	UpdatedBy     int64     `json:"updatedBy"`
	Name          string    `json:"name"`
	Code          string    `json:"code"`
	CreatedByName string    `json:"createdByName"`
	UpdatedByName string    `json:"updatedByName"`
}
