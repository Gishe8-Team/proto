package profile

import (
	"github.com/gofrs/uuid"
	"github.com/volatiletech/null/v8"
)

type Profile struct {
	ID            uuid.UUID   `boil:"id" json:"id,omitempty"`
	WalletID      uuid.UUID   `boil:"wallet_id" json:"wallet_id"`
	ClubID        null.String `boil:"club_id" json:"club_id,omitempty"`
	IsClubAdmin   null.Bool   `boil:"is_club_admin" json:"is_club_admin,omitempty"`
	UserID        uuid.UUID   `boil:"user_id" json:"user_id"`
	FirstName     string      `boil:"first_name" json:"first_name"`
	LastName      string      `boil:"last_name" json:"last_name"`
	OstanID       int         `boil:"ostan_id" json:"ostan_id"`
	CityID        int         `boil:"city_id" json:"city_id"`
	NationalID    null.String `boil:"nid" json:"national_id,omitempty"`
	GenderID      int16       `boil:"gender_id" json:"gender_id"`
	GenderName    string      `boil:"gender_name" json:"gender_name,omitempty"`
	WantReminder  null.Bool   `boil:"want_reminder" json:"want_reminder,omitempty"`
	WantPromotion null.Bool   `boil:"want_promotion" json:"want_promotion,omitempty"`
	AvatarURL     null.String `boil:"avatar_url" json:"avatar_url,omitempty"`
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
