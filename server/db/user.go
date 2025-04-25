package db

import (
	"street-art/models"

	"golang.org/x/crypto/bcrypt"
)

func (manager *Manager) Login(email string, password string) (models.User, error) {
	var user models.User
	var hashedPassword string

	query := `SELECT id, name, email, password_hash, phone, address FROM users WHERE email = $1;`
	err := manager.Conn.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &hashedPassword, &user.Phone, &user.Address)
	if err != nil {
		return models.User{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (manager *Manager) Register(user *models.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (name, email, password_hash, phone, address) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err = manager.Conn.QueryRow(query, user.Name, user.Email, hashedPassword, user.Phone, user.Address).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) UpdateUserInfo(user *models.User) error {
	query := `UPDATE users SET name = $1, phone = $2, address = $3 WHERE id = $4;`

	_, err := manager.Conn.Exec(query, user.Name, user.Phone, user.Address, user.Id)
	if err != nil {
		return err
	}

	return nil
}