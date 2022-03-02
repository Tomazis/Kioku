package models

type User struct {
	ID         uint64 `db:"id"`
	UserName   string `db:"username"`
	Password   string `db:"password"`
	Email      string `db:"email"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
}

type UserProgress struct {
	ID     uint64 `db:"id"`
	UserID uint64 `db:"user_id"`
	Level  uint32 `db:"level"`
}
