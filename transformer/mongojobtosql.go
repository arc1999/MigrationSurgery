package transformer

import "migrationSurgery/model"

func Transform(msurgeries []model.SurgeryMongo) []model.Surgery{
	var surgeries []model.Surgery
	for _, msurgery := range msurgeries {
		var surgery model.Surgery
		surgery.DateUpdated = msurgery.DateUpdated
		surgery.Code = msurgery.Code
		surgery.CreatedByName = msurgery.CreatedByName
		surgery.DateCreated = msurgery.DateCreated
		surgery.UpdatedBy = msurgery.UpdatedBy
		surgery.UpdatedByName = msurgery.UpdatedByName
		surgery.CreatedBy = msurgery.CreatedBy
		surgery.Name = msurgery.Name
		surgeries = append(surgeries, surgery)
	}
	return surgeries
}
