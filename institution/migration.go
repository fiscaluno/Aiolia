package institution

import "github.com/fiscaluno/pandorabox/db"

// Migrate migration entity BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	entity := new(Entity)

	entity.Name = "Impacta"
	entity.ImageURL = "https://www.impacta.edu.br/themes/wc_agenciar3/images/logo.png"
	entity.AverageRating = 5
	entity.RatedByCount = 1
	entity.Website = "https://www.impacta.edu.br/"
	entity.Cnpj = "59.069.914/0001-51"
	entity.Address = "Avenida Rudge, 315 - Barra Funda"
	entity.City = "São Paulo"
	entity.Province = "SP"

	// Migrate the schema
	db.AutoMigrate(&entity)

	// Create
	db.Create(&entity)

	// Read
	// var entity Entity
	db.First(&entity, 1) // find entity with id 1
	// db.First(&entity, "name = ?", "JC") // find entity with name JC

	// Update - update entity's Name to SI
	// db.Model(&entity).Update("Name", "SI")

	// Delete - delete entity
	// db.Delete(&entity)
}
