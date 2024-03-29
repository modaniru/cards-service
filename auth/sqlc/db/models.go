// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type User struct {
	ID       int32
	Email    sql.NullString
	Password sql.NullString
	Username sql.NullString
}

type UsersAuth struct {
	UserID   sql.NullInt32
	AuthType sql.NullString
	AuthID   sql.NullString
}
