package requests

type Header struct {
	Product				string  `json:"Product"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Plant				string	`json:"Plant"`
	Batch				string	`json:"Batch"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityStartTime	string	`json:"ValidityStartTime"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	ValidityEndTime		string	`json:"ValidityEndTime"`
	ManufactureDate		*string	`json:"ManufactureDate"`
	CreationDate         string  `json:"CreationDate"`
	CreationTime         string  `json:"CreationTime"`
	LastChangeDate       string  `json:"LastChangeDate"`
	LastChangeTime       string  `json:"LastChangeTime"`	
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
