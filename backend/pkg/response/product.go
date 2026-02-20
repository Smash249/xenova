package response

import (
	"time"

	"github.com/Smash249/xenova/backend/internal/models"
)

type ProductDetailResp struct {
	ID          uint              `json:"id"`
	SeriesID    uint              `json:"series_id"`
	SeriesName  string            `json:"series_name"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Cover       string            `json:"cover"`
	Previews    models.StringList `json:"previews"`
	Price       float64           `json:"price"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
