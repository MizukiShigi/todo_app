package models

import (
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (uuid, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd, CreateUUID(), u.Name, u.Email, Encript(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}