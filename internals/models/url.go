package models

import "time"

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}

type URLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
}

type ShortURL struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	Clicks      int64     `json:"clicks"`
}
