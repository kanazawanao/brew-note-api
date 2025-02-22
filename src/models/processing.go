package models

type Processing struct {
	ID          int    `json:"id"`
	Method      string `json:"method"`
	Description string `json:"description"`
}
