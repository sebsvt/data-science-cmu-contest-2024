package valueobject

type Address struct {
	AddressLine string   `bson:"address_line" json:"address_line"` // Address line
	Locality    string   `bson:"locality" json:"locality"`         // Locality or neighborhood
	City        string   `bson:"city" json:"city"`                 // City
	State       string   `bson:"state" json:"state"`               // State or province
	PostalCode  string   `bson:"postal_code" json:"postal_code"`   // Postal code
	Country     string   `bson:"country" json:"country"`           // Country
	Latitude    *float64 `bson:"latitude" json:"latitude"`         // latitude can be null
	Longitude   *float64 `bson:"longitude" json:"longitude"`       // logitude can be null
}
