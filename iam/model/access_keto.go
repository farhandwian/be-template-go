package model

type AccessKetoStruct struct {
	Namespace string `json:"namespace"`
	Object    string `json:"object"`
	Relation  string `json:"relation"`
}

// Create an alias for clarity
type AccessKeto = AccessKetoStruct

func NewAccessKeto(access AccessKeto) AccessKeto {
	return access
}
