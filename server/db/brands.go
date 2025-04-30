package db

import (
	"database/sql"
	"fmt"
	"street-art/models"
)

func (manager *Manager) GetMiniBrands() ([]models.BrandMini, error) {
	rows, err := manager.Conn.Query("SELECT id, name FROM brands")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no value")
		}
		return nil, err
	}
	defer rows.Close()

	var brands []models.BrandMini

	for rows.Next() {
		var brand models.BrandMini

		err := rows.Scan(&brand.Id, &brand.Name)
		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	return brands, nil
}