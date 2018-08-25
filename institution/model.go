package institution

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Institution is a Entity
type Institution struct {
	Name          string  `json:"name"`
	ImageURL      string  `json:"image_url"`
	AverageRating float64 `json:"average_rating"`
	RatedByCount  int     `json:"rated_by_count"`
	Website       string  `json:"website"`
	Cnpj          string  `json:"cnpj"`
	Address       string  `json:"address"`
	City          string  `json:"city"`
	Province      string  `json:"province"`
}

// Emails        []string `json:"emails"`
// Phones        []string `json:"phones"`

// Entity is a course
type Entity struct {
	Institution
	db.CommonModelFields
}

// Entitys is Entity slice
type Entitys []Entity

// GetAll Entitys
func GetAll() Entitys {
	db := db.Conn()
	defer db.Close()
	var entitys Entitys
	db.Find(&entitys)
	return entitys
}

// Save a Entity
func (entity Entity) Save() (Entity, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Entity
func GetByID(id int) Entity {
	db := db.Conn()
	defer db.Close()

	var entity Entity

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Entity
func GetByQuery(query string, value interface{}) Entitys {
	db := db.Conn()
	defer db.Close()

	var entitys Entitys

	db.Find(&entitys, query, value)
	return entitys
}
