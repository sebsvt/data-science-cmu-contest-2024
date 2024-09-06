package entity

import (
	"github.com/sebsvt/software-prototype/transportation-software-services/valueobject"
)

type Person struct {
	ID        string              `bson:"id" json:"id"`               // Unique identifier for the person
	FirstName string              `bson:"firstname" json:"firstname"` // First name of the person
	LastName  string              `bson:"lastname" json:"lastname"`   // Last name of the person
	Phone     string              `bson:"phone_number" json:"phone"`  // Phone number of the person
	Address   valueobject.Address `bson:"address" json:"address"`
}
