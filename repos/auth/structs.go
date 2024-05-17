package auth

import "time"

type User struct {
	Id        string    `db:"user_id"`
	Nip       int64     `db:"nip"`
	Name      string    `db:"name"`
	Admin     bool      `db:"is_admin"`
	Password  string    `db:"password"`
	Active    bool      `db:"active"`
	ImageUrl  string    `db:"image_url"`
	CreatedAt time.Time `db:"created_at"`
}

type UserFilter struct {
	UserId        *string
	Limit         int
	Offset        int
	Name          *string
	Nip           *int64
	Role          *string
	CreatedAtSort *string
}

type AllUsersFilter struct {
	UserIds []string
}

type ActivateUserReq struct {
	Id       string `db:"user_id"`
	Password string `db:"password"`
}
