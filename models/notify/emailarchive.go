package notify

import "time"

type EmailArchiveRequest struct {
	EmailAddress string `json:"email_address,omitempty"`
	Service      string `json:"service,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

type EmailArchive struct {
	ID           string    `json:"id" boil:"id"`
	EmailAddress string    `json:"email_address" boil:"email_address"`
	EmailMessage string    `json:"email_message" boil:"email_message"`
	WhichService string    `json:"which_service" boil:"which_service"`
	SentTime     time.Time `json:"sent_time" boil:"created_at"`
}

type EmailArchiveSlice []*EmailArchive
