package models

type Block struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Image       string `json:"image"`
}
