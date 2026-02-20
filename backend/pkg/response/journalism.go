package response

import "time"

type JournalismDetailResp struct {
	ID          uint      `json:"id"`
	SeriesID    uint      `json:"series_id"`
	SeriesName  string    `json:"series_name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	ViewCount   uint      `json:"view_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
