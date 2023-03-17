package profile

import "github.com/gofrs/uuid"

type Profile struct {
	ID            uuid.UUID `json:"id,omitempty"`
	WalletID      uuid.UUID `json:"wallet_id"`
	ClubID        uuid.UUID `json:"club_id,omitempty"`
	IsClubAdmin   bool      `json:"is_club_admin,omitempty"`
	UserID        uuid.UUID `json:"user_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	OstanID       int       `json:"ostan_id"`
	CityID        int       `json:"city_id"`
	NationalID    string    `json:"national_id,omitempty"`
	GenderID      int16     `json:"gender_id"`
	WantReminder  bool      `json:"want_reminder,omitempty"`
	WantPromotion bool      `json:"want_promotion,omitempty"`
	AvatarURL     string    `json:"avatar_url,omitempty"`
}

type QueryProfile struct {
	UserId        uuid.UUID `json:"user_id,omitempty"`
	ClubID        uuid.UUID `json:"club_id,omitempty"`
	IsAdmin       bool      `json:"is_admin,omitempty"`
	FirstName     string    `json:"first_name,omitempty"`
	LastName      string    `json:"last_name,omitempty"`
	OstanID       int       `json:"ostan_id,omitempty"`
	CityID        string    `json:"city_id,omitempty"`
	GenderID      int16     `json:"gender_id,omitempty"`
	WantReminder  bool      `json:"want_reminder,omitempty"`
	WantPromotion bool      `json:"want_promotion,omitempty"`
	Limit         int       `json:"limit,omitempty"`
	Offset        int       `json:"offset,omitempty"`
}
