package db

import (
	"database/sql"
	"fmt"
	"street-art/models"
)

func (manager *Manager) GetProductById(id int) (models.Product, error) {
	var product models.Product

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
			WHERE p.id = $1`
	err := manager.Conn.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageUrl, &product.Brand, &product.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product not found")
		} else {
			return product, err
		}
	}

	return product, nil
}

func (manager *Manager) GetAllProducts() ([]models.Product, error) {
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
			LEFT JOIN categories c ON p.category_id = c.id`
	rows, err := manager.Conn.Query(query)
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

func (manager *Manager) AddToCart(product *models.ProductToCart) error {
	query := `INSERT INTO cart (user_id, product_id, count) VALUES ($1, $2, $3)`
	_, err := manager.Conn.Exec(query, &product.UserId, &product.Id, &product.Count)
	if err != nil {
		return err
	}

	return nil
}

func (Manager *Manager) GetTopProducts() ([]models.Product, error) {
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
			FROM 
				products p
			INNER JOIN 
				top_products tp ON p.id = tp.product_id
			LEFT JOIN 
				brands b ON p.brand_id = b.id
			JOIN 
				categories c ON p.category_id = c.id;`
	rows, err := Manager.Conn.Query(query)
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

func (manager *Manager) AddProduct(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, stock, image_url, brand_id, category_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := manager.Conn.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.ImageUrl, product.BrandId, product.CategoryId)
	if err != nil {
		return err
	}

	return nil
}