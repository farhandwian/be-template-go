package model

// MapAccess represents a single permission.
type MapAccessKeto struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Group       string `json:"group,omitempty"` // e.g., the role name ("viewer", "owner")
	Type        string `json:"type,omitempty"`  // e.g., "READ", "EDIT"
	Namespace   string `json:"namespace"`       // e.g., "app"
	Object      string `json:"object"`          // e.g., "dashboard", "user-management"
	Relation    string `json:"relation"`        // e.g., "read", "edit"
	Enabled     bool   `json:"enabled"`         // computed flag (true if the role grants this permission)
}

// RolePermissions groups a role with its assigned permissions.
type RolePermissions struct {
	Role        string          `json:"role"`
	Permissions []MapAccessKeto `json:"permissions"`
}

func GetMapAccessKeto() []MapAccessKeto {
	return []MapAccessKeto{
		{
			ID:          "1",
			Description: "Data Perangkat",
			Group:       "Dashboard",
			Type:        "READ",
			Namespace:   "app",
			Object:      "dashboard:data-perangkat",
			Relation:    "read",
			Enabled:     false,
		},
		{
			ID:          "2",
			Description: "Data Si JagaCai",
			Group:       "Dashboard",
			Type:        "READ",
			Namespace:   "app",
			Object:      "dashboard:jaga-cai",
			Relation:    "read",
			Enabled:     false,
		},
	}
}
