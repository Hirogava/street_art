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

func (manager *Manager) AddCategory(category *models.Category) error {
	query := `INSERT INTO categories(name, description, image_url) VALUES($1, $2, $3)`
	_, err := manager.Conn.Exec(query, category.Name, category.Description, category.ImageUrl)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) GetCategoryById(id int) (models.Category, error) {
	var category models.Category

	query := `SELECT * FROM categories WHERE id = $1`
	err := manager.Conn.QueryRow(query, id).Scan(&category.Id, &category.Name, &category.Description, &category.ImageUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, fmt.Errorf("no value")
     	}
	 	return models.Category{}, err
	}

	return category, nil
}

func (manager *Manager) DeleteCategoryById(id int) (string, error) {
	var imageUrl string
	query := `DELETE FROM categories WHERE id = $1 RETURNING image_url`
	err := manager.Conn.QueryRow(query, id).Scan(&imageUrl)
	if err != nil {
		return "", err
	}

	return imageUrl, nil
}

func (manager *Manager) UpdateCategory(category *models.Category) error {
	query := `UPDATE categories SET name = $1, description = $2, image_url = $3 WHERE id = $4`
	_, err := manager.Conn.Exec(query, category.Name, category.Description, category.ImageUrl, category.Id)
	if err != nil {
		return err
	}

	return nil
}