package proto

type InternalMessage struct {
	ConnectionID string `json:"connection_id"`
	ServiceID    string `json:"service_id"`
	Message      string `json:"message"`
	RoutingKey   string `json:"routing_key"`

	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"status_code"`
	Body       []byte            `json:"body"`
}
