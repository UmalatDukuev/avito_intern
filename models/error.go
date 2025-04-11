package models

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type" db:"type"`
}
