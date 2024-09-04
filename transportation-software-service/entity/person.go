package entity

import "github.com/google/uuid"

type Person struct {
	ID          uuid.UUID `bson:"id"`           // Unique identifier for the person
	FirstName   string    `bson:"firstname"`    // First name of the person
	LastName    string    `bson:"lastname"`     // Last name of the person
	Phone       string    `bson:"phone_number"` // Phone number of the person
	AddressLine string    `bson:"address_line"` // Address line
	Locality    string    `bson:"locality"`     // Locality or neighborhood
	City        string    `bson:"city"`         // City
	State       string    `bson:"state"`        // State or province
	PostalCode  string    `bson:"postal_code"`  // Postal code
	Country     string    `bson:"country"`      // Country
	Latitude    float64   `bson:"latitude"`
	Longitude   float64   `bson:"longitude"`
}
