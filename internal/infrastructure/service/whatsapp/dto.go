package whatsapp

type (
	Message struct {
		PhoneNumber string `json:"to"`
		Message     string `json:"body"`
	}
)
