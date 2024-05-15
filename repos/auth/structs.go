package auth

type User struct {
	Id        string `db:"user_id"`
	Nip       int64  `db:"nip"`
	Name      string `db:"name"`
	Password  string `db:"password"`
	Active    bool   `db:"active"`
	ImageUrl  string `db:"image_url"`
	CreatedAt string `db:"created_at"`
}
