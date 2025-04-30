package db

import (
	"database/sql"
	"fmt"
	"street-art/models"
)

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

func (manager *Manager) GetMiniCategories() ([]models.CategoryMini, error){
	rows, err := manager.Conn.Query("SELECT id, name FROM categories")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no value")
		}
		return nil, err
	}
	defer rows.Close()

	var categories []models.CategoryMini

	for rows.Next(){
		var category models.CategoryMini
		
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (manager *Manager) GetAllProductsByCategoryId(id int) ([]models.Product, error){
	var products []models.Product

	query := `SELECT 
				p.id,
				p.name,
				p.description,
				p.price,
				p.stock,
				p.image_url,
				b.name AS brand,
				c.name AS category
			FROM products p
			LEFT JOIN brands b ON p.brand_id = b.id
			LEFT JOIN categories c ON p.category_id = c.id
			WHERE p.category_id = $1`
	
	rows, err := manager.Conn.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageUrl, &product.Brand, &product.Category)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}