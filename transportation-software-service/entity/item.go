package entity

type Item struct {
	ItemCode        string `bson:"item_code" json:"item_code"`               // Code for the item
	ItemName        string `bson:"item_name" json:"item_name"`               // Name of the item
	ItemDescription string `bson:"item_description" json:"item_description"` // Description of the item
	Unit            string `bson:"unit" json:"unit"`                         // Measurement unit (e.g., kg, pcs)
	Quantity        int    `bson:"quantity" json:"quantity"`                 // Quantity of the item
}
