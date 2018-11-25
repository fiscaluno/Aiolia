package institution

import (
	"time"

	"github.com/fiscaluno/pandorabox/db"
)

// Institution is a Institution
type Institution struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	Name          string     `json:"name"`
	ImageURL      string     `json:"image_url"`
	AverageRating float64    `json:"average_rating"`
	RatedByCount  int        `json:"rated_by_count"`
	Website       string     `json:"website"`
	Cnpj          string     `json:"cnpj"`
	Address       string     `json:"address"`
	City          string     `json:"city"`
	Province      string     `json:"province"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

// Emails        []string `json:"emails"`
// Phones        []string `json:"phones"`

// TableName for review
func (Institution) TableName() string {
	return "institution"
}

// GetAll []Institution
func GetAll() []Institution {
	db := db.Conn()
	defer db.Close()
	var entities []Institution
	db.Find(&entities)
	return entities
}

// Save a Institution
func (entity Institution) Save() (Institution, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Institution
func GetByID(id int) Institution {
	db := db.Conn()
	defer db.Close()

	var entity Institution

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Institution
func GetByQuery(query string, value interface{}) []Institution {
	db := db.Conn()
	defer db.Close()

	var entities []Institution

	db.Find(&entities, query, value)
	return entities
}
