package entity

import "time"

type Post struct {
	PublicationId   string    `json:"publicationId"`
	Description     string    `json:"description"`
	PublicationDate time.Time `json:"publicationDate"`
	Receiver        string    `json:"receiver"`
}
