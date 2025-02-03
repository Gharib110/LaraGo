package data

import (
	up "github.com/upper/db/v4"
	"time"
)

type User struct {
	ID        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Active    int       `db:"user_active"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Token     Token     `db:"-"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) GetAll(condition up.Cond) ([]*User, error) {
	collection := up.Collection(condition)

	var users []*User
	res := collection.Find(condition)
	err := res.All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	var user User
	collection := up.Collection(u.Table())
	res := collection.Find(up.Cond{"email": email})
	err := res.One(&user)
	if err != nil {
		return nil, err
	}

	var token Token
	coll := up.Collection(token.Table())
	res = coll.Find(up.Cond{"user_id =": user.ID, "expiry <": time.Now()}).OrderBy("Created_at desc")
	err = res.One(&token)
	if err != nil {
		return nil, err
	}

	user.Token = token

	return &user, nil
}
