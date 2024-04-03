package entity

type Message struct {
	ID          string `json:"messageId"`
	SenderID    string `json:"senderId,omitempty"`
	RecipientID string `json:"recipientId,omitempty"`
	Message     string `json:"message,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	ReceivedAt  string `json:"receivedAt,omitempty"`
	ReadAt      string `json:"readAt,omitempty"`
}
