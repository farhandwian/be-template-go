package model

import "time"

// AuthenticationMethod struct untuk authentication_methods
type AuthenticationMethod struct {
	Method      string    `json:"method"`
	AAL         string    `json:"aal"`
	CompletedAt time.Time `json:"completed_at"`
}

// Traits struct untuk menyimpan informasi pengguna
type Traits struct {
	Username string `json:"username"`
}

// Identity struct untuk menyimpan informasi identity pengguna
type Identity struct {
	ID             string    `json:"id"`
	SchemaID       string    `json:"schema_id"`
	SchemaURL      string    `json:"schema_url"`
	State          string    `json:"state"`
	StateChangedAt time.Time `json:"state_changed_at"`
	Traits         Traits    `json:"traits"`
	MetadataPublic *string   `json:"metadata_public"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	OrganizationID *string   `json:"organization_id"`
}

// Device struct untuk menyimpan informasi perangkat pengguna
type Device struct {
	ID        string `json:"id"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	Location  string `json:"location"`
}

// Session struct untuk menyimpan seluruh informasi JSON
type Session struct {
	ID                          string                 `json:"id"`
	Active                      bool                   `json:"active"`
	ExpiresAt                   time.Time              `json:"expires_at"`
	AuthenticatedAt             time.Time              `json:"authenticated_at"`
	AuthenticatorAssuranceLevel string                 `json:"authenticator_assurance_level"`
	AuthenticationMethods       []AuthenticationMethod `json:"authentication_methods"`
	IssuedAt                    time.Time              `json:"issued_at"`
	Identity                    Identity               `json:"identity"`
	Devices                     []Device               `json:"devices"`
}
