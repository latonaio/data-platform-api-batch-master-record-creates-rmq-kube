package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-batch-master-record-creates-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToBatchCreates(sdc *dpfm_api_input_reader.SDC) (*Batch, error) {
	data := sdc.Batch

	batch, err := TypeConverter[*Batch](data)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func ConvertToBatchUpdates(batchData dpfm_api_input_reader.Batch) (*Batch, error) {
	data := batchData

	batch, err := TypeConverter[*Batch](data)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
