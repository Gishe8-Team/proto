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
