package v1

import (
	"time"

	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/req"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

func toModelBook(d *aphv1req.CreateBook) *apv1models.Book {
	now := time.Now().UTC().Unix()
	return &apv1models.Book{
		Title:       d.Title,
		Original:    d.Original,
		PublishedAt: d.PublishedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
		CategoryID:  d.CategoryID,
	}
}
