package authorization

import (
	"context"
	"database/sql"
	"log"

	"github.com/volatiletech/authboss/v3"
)

type CreatingUserStorer struct {
	Database *sql.DB
}

func (creatingStorer CreatingUserStorer) Load(_ context.Context, username string) (authboss.User, error) {
	rows, err := creatingStorer.Database.Query(`SELECT id, name, email, password FROM users WHERE name = $1`, username)
	if err != nil {
		log.Fatal(err)
	}
	if rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		return user, err
	}

	return nil, authboss.ErrUserNotFound
}

func (creatingStorer CreatingUserStorer) Save(_ context.Context, user authboss.User) error {
	u := user.(*User)
	_, err := creatingStorer.Database.Exec(`UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`, u.ID, u.Username, u.Email, u.Password, u.ID)
	return err
}

func (creatingStorer CreatingUserStorer) New(_ context.Context) authboss.User {
	return &User{}
}

func (creatingStorer CreatingUserStorer) Create(ctx context.Context, user authboss.User) error {
	u := user.(*User)
	err := creatingStorer.Database.QueryRow(`INSERT INTO users (name, email, password) VALUES($1, $2, $3)`, u.Username, u.Email, u.Password)
	if err != nil {
		return authboss.ErrUserFound
	}
	return nil
}
