package model

import (
	"errors"
	"time"
)

// UserInfo holds the parameters returned by the backends.
// This information will be serialized to build the JWT token contents.
type UserInfo struct {
	Sub       string   `json:"sub"`
	Picture   string   `json:"picture,omitempty"`
	Name      string   `json:"name,omitempty"`
	Email     string   `json:"email,omitempty"`
	Origin    string   `json:"origin,omitempty"`
	Expiry    int64    `json:"exp,omitempty"`
	Refreshes int      `json:"refs,omitempty"`
	Domain    string   `json:"domain,omitempty"`
	Groups    []string `json:"groups,omitempty"`
}

// Valid lets us use the user info as Claim for jwt-go.
// It checks the token expiry.
func (u UserInfo) Valid() error {
	if u.Expiry < time.Now().Unix() {
		return errors.New("token expired")
	}
	return nil
}

func (u UserInfo) AsMap() map[string]interface{} {
	m := map[string]interface{}{
		"sub": u.Sub,
	}
	if u.Picture != "" {
		m["picture"] = u.Picture
	}
	if u.Name != "" {
		m["name"] = u.Name
	}
	if u.Email != "" {
		m["email"] = u.Email
	}
	if u.Origin != "" {
		m["origin"] = u.Origin
	}
	if u.Expiry != 0 {
		m["exp"] = u.Expiry
	}
	if u.Refreshes != 0 {
		m["refs"] = u.Refreshes
	}
	if u.Domain != "" {
		m["domain"] = u.Domain
	}
	if len(u.Groups) > 0 {
		m["groups"] = u.Groups
	}
	return m
}
