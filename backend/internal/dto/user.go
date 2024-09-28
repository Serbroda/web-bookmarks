package dto

type UserDto struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
}
