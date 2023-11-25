package db

import (
	"log"
	"time"

	"github.com/cata85/balloons/types"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	tableName struct{}  `sql:"users_collection"`
	ID        int       `sql:"id,pk"`
	Name      string    `sql:"name,unique"` // This is for demonstration, in production it would not be unique and would use email.
	Password  string    `sql:"password"`
	CreatedAt time.Time `sql:"created_at,unique"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

/**
 * Creates the Postgres users_collection table
 */
func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&User{}, opts)
	if createErr != nil {
		log.Printf("%v\n", createErr)
		return createErr
	}
	log.Printf("Created user table.\n")
	return nil
}

/**
 * This is the Saving function for the user type.
 */
func (user *User) Save(db *pg.DB) error {
	_, insertErr := db.Model(user).
		OnConflict("(name) DO UPDATE").
		Set("updated_at = ?updated_at, is_active = ?is_active").
		Insert()
	if insertErr != nil {
		log.Printf("Insert Error: %v\n", insertErr)
		return insertErr
	}
	return nil
}

/**
 * This prepares to save a single user.
 * The user is sent over to the user's "Save()" method.
 */
func SaveUser(user types.User) {
	if user.Name != "" && user.Password != "" {
		var db *pg.DB = Connect()
		defer db.Close()
		newUser := &User{
			Name:      user.Name,
			Password:  user.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsActive:  true,
		}
		newUser.Save(db)
	}
	return
}

/**
 * This gives a user given a username.
 */
func GetOneUser(name string) *User {
	var user = new(User)
	var db *pg.DB = Connect()
	defer db.Close()
	err := db.Model(user).Where("name = ?", name).Select()
	if err != nil {
		log.Printf("User query error: %v\n", err)
		return nil
	}
	return user
}
