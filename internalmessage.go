package proto

import (
	"fmt"
	"github.com/Gishe8-Team/proto/models/event"
)

type InternalMessage struct {
	ConnectionID string `json:"connection_id"`
	ServiceID    string `json:"service_id"`
	RoutingKey   string `json:"routing_key"`

	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"status_code"`
	Body       []byte            `json:"body"`
	Request    ServiceRequests   `json:"request"`
}

func (im *InternalMessage) String() string {
	return fmt.Sprintf("[ConnectionId: %s, ServiceID: %s, RoutingKey: %s, Headers: %+v, StatusCode: %d, Body: %s, Req: %+v]",
		im.ConnectionID,
		im.ServiceID,
		im.RoutingKey,
		im.Headers,
		im.StatusCode,
		string(im.Body),
		im.Request)
}

type ServiceRequests struct {
	event.Services
}
