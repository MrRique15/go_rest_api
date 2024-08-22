package models

type KafkaOrderEvent struct {
	Event        string              `json:"event"`
	Order        Order               `json:"order"`
	CreatedAt    string              `json:"created_at"`
	Source       string              `json:"source"`
	Status       string              `json:"status"`
	EventHistory []KafkaEventHistory `json:"event_history"`
}

type KafkaEventHistory struct {
	Source    string `json:"source"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	Event     string `json:"event"`
}
