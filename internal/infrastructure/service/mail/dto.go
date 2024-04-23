package mail

type (
	Message struct {
		DefaultUser string   `json:"default_user"`
		To          []string `json:"to"`
		Cc          []string `json:"cc"`
		Subject     string   `json:"subject"`
		Message     string   `json:"message"`
	}
)
