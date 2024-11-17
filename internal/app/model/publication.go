package model

import "time"

type Publication struct {
	ID        int `json:"id"`
	Text      string `json:"text"`
	CreatedAt time.Time `json:"created"`
	FileIds     []int `json:"file_ids"`
}