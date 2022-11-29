package dpfm_api_input_reader

import (
	"data-platform-api-batch-master-record-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBatch() *requests.Batch {
	data := sdc.BatchMasterRecord
	return &requests.Batch{
		Product:             data.Product,
		BusinessPartner:     data.BusinessPartner,
		Plant:               data.Plant,
		Batch:               data.Batch,
		ValidityEndDate:     data.ValidityEndDate,
		CountryOfOrigin:     data.CountryOfOrigin,
		ValidityStartDate:   data.ValidityStartDate,
		ManufactureDate:     data.ManufactureDate,
		CreationDateTime:    data.CreationDateTime,
		LastChangeDateTime:  data.LastChangeDateTime,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
