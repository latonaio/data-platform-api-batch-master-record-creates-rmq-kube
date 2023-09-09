package dpfm_api_processing_formatter

type BatchUpdates struct {
	Product				string  `json:"Product"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Plant				string	`json:"Plant"`
	Batch				string	`json:"Batch"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityStartTime	string	`json:"ValidityStartTime"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	ValidityEndTime		string	`json:"ValidityEndTime"`
	ManufactureDate		*string	`json:"ManufactureDate"`
}
