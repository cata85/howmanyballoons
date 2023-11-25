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
 * Creates the Postgres balloon_objects_collection table
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
 * This is the Saving function for the balloon object type.
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
 * This prepares to save a single balloon object.
 * The balloon object is sent over to the balloon's "Save()" method.
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
