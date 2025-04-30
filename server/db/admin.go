package db

import (
	"street-art/models"

	"golang.org/x/crypto/bcrypt"
)

func (manager *Manager) AdminLogin(email string, password string) (models.Admin, error) {
	var user models.Admin
	var hashedPassword string

	query := `SELECT id, email, password_hash FROM admin_panel WHERE email = $1;`
	err := manager.Conn.QueryRow(query, email).Scan(&user.Id, &user.Email, &hashedPassword)
	if err != nil {
		return models.Admin{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return models.Admin{}, err
	}

	return user, nil
}