package model

import "time"

type Message struct {
    ID string `json:"id"`
    Type string `json:"type"`
    Content string `json:"content"`
    ImageURL string `json:"image_url,omitempty"`
    CreatedAt time.Time `json:"created_at"`
}
