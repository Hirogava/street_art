package db

import "street-art/models"

func (manager *Manager) GetAllCategories() ([]models.Category, error){
	rows, err := manager.Conn.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next(){
		var category models.Category
		
		err := rows.Scan(&category.Id, &category.Name, &category.Description, &category.ImageUrl)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}