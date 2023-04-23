package sso

import "time"

type Claims struct {
	Issuer         string    `json:"iss"`
	Subject        string    `json:"sub"`
	Audience       []string  `json:"aud"`
	ExpirationTime time.Time `json:"exp"`
	NotBefore      time.Time `json:"nbf"`
	IssuedAt       time.Time `json:"iat"`
}
type LoginToken struct {
	Claims
	UserID string `json:"user_id"`
	ACL    string `json:"acl"`
}

type CodeToken struct {
	Claims
	UserID string `json:"user_id"`
	Mobile string `json:"mobile,omitempty"`
	Email  string `json:"email,omitempty"`
}
