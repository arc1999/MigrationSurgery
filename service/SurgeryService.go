package service

import (
	"MigrationSurgery/transformer"
	"log"
	"MigrationSurgery/dao"
	"os"
	"strconv"
)

var d dao.SurgeryDao

type SurgeryService struct {
}

func (is SurgeryService) Migrate() {
	totaldoc, err := d.GetCount()
	if err != nil {
		log.Fatal(err)
	}
	perpage := os.Getenv("N_PER_PAGE")
	nperpage, err := strconv.ParseInt(perpage, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var i int64
	for i*nperpage < totaldoc {
		msurgeries, err := d.Paginate(i, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		surgeries := transformer.Transform(msurgeries)
		err = d.BulkInsert(surgeries, nperpage)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
}
