package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-batch-master-record-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToBatchUpdates(header dpfm_api_input_reader.Batch) *BatchUpdates {
	data := batch

	return &BatchUpdates{
		Product:				data.Product,
		BusinessPartner:		data.BusinessPartner,
		Plant:					data.Plant,
		Batch:					data.Batch,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityStartTime:		data.ValidityStartTime,
		ValidityEndDate:		data.ValidityEndDate,
		ValidityEndTime:		data.ValidityEndTime,
		ManufactureDate:		data.ManufactureDate,
	}
}
