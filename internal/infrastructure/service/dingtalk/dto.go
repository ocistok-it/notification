package dingtalk

import "regexp"

type (
	Message struct {
		AccessToken string      `json:"-"`
		RobotID     string      `json:"robot_id"`
		Type        MessageType `json:"type"`
		Content     string      `json:"content"`
		Mobile      []string    `json:"mobile"`
		IsAtAll     bool        `json:"is_at_all"`
	}

	MessageType string
)

var (
	regStr     = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	regPattern = regexp.MustCompile(regStr)
)

const (
	TypeText       MessageType = "text"
	TypeLink       MessageType = "link"
	TypeMarkdown   MessageType = "markdown"
	TypeActionCard MessageType = "actionCard"
	TypeFeedCard   MessageType = "feedCard"
)

func (m *Message) TextMessage() map[string]interface{} {
	request := map[string]interface{}{
		"msgtype": m.Type,
		"text": map[string]interface{}{
			"content": m.Content,
		},
		"at": map[string]interface{}{
			"atMobiles": m.Mobile,
			"isAtAll":   m.IsAtAll,
		},
	}

	return request
}
