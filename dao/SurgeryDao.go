package dao

import (
	"MigrationSurgery/model"
	"context"
	"MigrationSurgery/db"
	log "github.com/sirupsen/logrus"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type SurgeryDao struct {
}

func (d SurgeryDao) Paginate(pagenumber int64, nperpage int64) ([]model.SurgeryMongo, error) {
	log.Println(pagenumber,nperpage)
	options := options.Find()
	options.SetLimit(nperpage)
	options.SetSort(bson.M{})
	options.SetSkip(pagenumber)

	db := db.GetMongoDB()
	cur, err := db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var jobs []model.SurgeryMongo
	for cur.Next(context.TODO()) {
		var job model.SurgeryMongo
		err := cur.Decode(&job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	log.Println(len(jobs))
	return jobs, nil
}
func (d SurgeryDao) GetCount() (int64, error) {
	db := db.GetMongoDB()
	return db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).CountDocuments(context.TODO(),bson.D{})
}
func (d SurgeryDao) BulkInsert(Entity []model.Surgery, nperpage int64) error {
	sqldb := db.GetMysqlDB()
	b := make([]interface{}, len(Entity))
	for i := range Entity {
		b[i] = Entity[i]
	}
	err := gormbulk.BulkInsert(sqldb, b, int(nperpage))
	if err != nil {
		log.Printf("error in saving surgery")
		return err
	}
	return nil

}
