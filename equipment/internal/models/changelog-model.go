package models

type Changelog struct {
	IndexNum          int
	ID                int
	Timestamp         string
	Name              string
	Type              string
	Model             string
	Serial            string
	Description       string
	Brand             string
	Price             string
	Manufacturer      string
	Expiration        string
	PurchaseDate      string
	CalibrationDate   string
	CalibrationMethod string
	NextCalibration   string
	Location          string
	IssuedBy          string
	IssuedTo          string
	Remarks           string
	Status            string
	ModifiedBy        string
}
