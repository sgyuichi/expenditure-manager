package entity

// Expenditure describes expenditure struct
type Expenditure struct {
	ID          string   `json:"id" bson:"_id"`
	UserID      string   `json:"user_id" bson:"user_id"`
	Value       int      `json:"value" bson:"value"`
	Date        int64    `json:"date" bson:"date"`
	Recipient   string   `json:"recipient" bson:"recipient"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
}
