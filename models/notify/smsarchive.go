package notify

import "time"

type SMSArchiveRequest struct {
	MobileNumber string `json:"mobile_number,omitempty"`
	Service      string `json:"service,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

type SMSArchive struct {
	ID           int64     `json:"id,omitempty" boil:"id"`
	MobileNum    string    `json:"mobile_num" boil:"mobile_num"`
	Message      string    `json:"message" boil:"message"`
	WhichService string    `json:"which_service" boil:"which_service"`
	SentTime     time.Time `json:"sent_time" boil:"created_at"`
}

type SMSArchiveSlice []*SMSArchive
