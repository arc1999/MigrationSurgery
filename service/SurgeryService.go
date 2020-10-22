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
	log.Println(totaldoc)
	var i int64
	i=0
	var noofpages = totaldoc/nperpage
	log.Println(noofpages)
	for i < noofpages {
		msurgeries, err := d.Paginate(i*nperpage, nperpage)
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
	msurgeries, err := d.Paginate(i*nperpage, totaldoc - (nperpage*(i)))
	if err != nil {
		log.Fatal(err)
	}
	surgeries := transformer.Transform(msurgeries)
	err = d.BulkInsert(surgeries, nperpage)
	if err != nil {
		log.Fatal(err)
	}
}
