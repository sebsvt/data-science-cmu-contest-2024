package event

type UpdatedEmail struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type CreatedAccount struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EventHandler interface {
	Sender(name string, event_bytes []byte) error
}
