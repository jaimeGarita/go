package model

import "time"

type User struct {
	ID        int       `json: "id"`
	Name      string    `json: "name"`
	Email     string    `json: "email"`
	Password  string    `json: "-"`
	CreatedAt time.Time `json: "created_at"`
}

type UserDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
