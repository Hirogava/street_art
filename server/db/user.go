package db

import (
	"street-art/models"

	"golang.org/x/crypto/bcrypt"
)

func (manager *Manager) Login(email string, password string) (models.User, error) {
	var user models.User
	var hashedPassword string

	query := `SELECT id, name, email, password_hash, phone, address, balance FROM users WHERE email = $1;`
	err := manager.Conn.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &hashedPassword, &user.Phone, &user.Address, &user.Balance)
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

	query := `INSERT INTO users (name, email, password_hash, phone, address) VALUES ($1, $2, $3, $4, $5) RETURNING id, balance;`
	err = manager.Conn.QueryRow(query, user.Name, user.Email, hashedPassword, user.Phone, user.Address).Scan(&user.Id, &user.Balance)
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

func (manager *Manager) GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, name, email, phone, address, balance FROM users;`

	rows, err := manager.Conn.Query(query)
	if err != nil {
		return nil, err
	}
   	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Address, &user.Balance)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}