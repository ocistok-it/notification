package mail

type (
	Message struct {
		To      []string `json:"to"`
		Cc      []string `json:"cc"`
		Subject string   `json:"subject"`
		Message string   `json:"message"`
	}
)
