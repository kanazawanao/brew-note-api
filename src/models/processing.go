package models

type Processing struct {
	ID          string `json:"id"`
	Method      string `json:"method"`
	Description string `json:"description"`
}
