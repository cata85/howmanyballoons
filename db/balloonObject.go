package db

import (
	"log"
	"time"

	"github.com/cata85/balloons/types"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type BalloonObject struct {
	tableName  struct{}  `sql:"balloon_objects_collection"`
	ID         int       `sql:"id,pk"`
	Name       string    `sql:"name,unique"`
	Weight     string    `sql:"weight"`
	Balloons   string    `sql:"balloons"`
	WeightType string    `sql:"weight_type"`
	CreatedAt  time.Time `sql:"created_at,unique"`
	UpdatedAt  time.Time `sql:"updated_at"`
	IsActive   bool      `sql:"is_active"`
}

/**
 * Creates the Postgres balloon_objects_collection table
 */
func CreateBalloonObjectTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&BalloonObject{}, opts)
	if createErr != nil {
		log.Printf("%v\n", createErr)
		return createErr
	}
	log.Printf("Created balloonObjects table.\n")
	return nil
}

/**
 * This is the Saving function for the balloon object type.
 */
func (balloonObject *BalloonObject) Save(db *pg.DB) error {
	_, insertErr := db.Model(balloonObject).
		OnConflict("(name) DO UPDATE").
		Set("weight = ?weight, balloons = ?balloons, updated_at = ?updated_at, is_active = ?is_active").
		Insert()
	if insertErr != nil {
		log.Printf("Insert Error: %v\n", insertErr)
		return insertErr
	}
	return nil
}

/** NOT IN USE
 * This saves and returns a balloon object.
 */
func (balloonObject *BalloonObject) SaveAndReturn(db *pg.DB) (*BalloonObject, error) {
	_, insertError := db.Model(balloonObject).Returning("*").Insert()
	if insertError != nil {
		log.Printf("Insert Error: %v\n", insertError)
		return nil, insertError
	}
	return balloonObject, nil
}

/** TODO
 * This saves many balloon objects.
 * (May use an alternative Save() method and use this one to prepare data. Currently broken.)
 */
func (balloonObject *BalloonObject) SaveMany(db *pg.DB, balloonObjects []*BalloonObject) error {
	_, insertErr := db.Model(balloonObjects).Insert()
	if insertErr != nil {
		log.Printf("Insert Errors: %v\n", insertErr)
		return insertErr
	}
	return nil
}

/**
 * This prepares to save a single balloon object.
 * The balloon object is sent over to the balloon's "Save()" method.
 */
func SaveBalloonObject(balloonObject types.BalloonObject) {
	if balloonObject.Name != "" {
		var db *pg.DB = Connect()
		defer db.Close()
		newBalloonObject := &BalloonObject{
			Name:       balloonObject.Name,
			Weight:     balloonObject.Weight,
			Balloons:   balloonObject.Balloons,
			WeightType: balloonObject.WeightType,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			IsActive:   true,
		}
		newBalloonObject.Save(db)
	}
	return
}

/**
 * This grabs just one balloon object from the given id value.
 */
func GetOneBalloonObject(id string) (*BalloonObject, error) {
	var balloonObject = new(BalloonObject)
	var db *pg.DB = Connect()
	defer db.Close()
	err := db.Model(balloonObject).Where("id = ?", id).Select()
	if err != nil {
		log.Printf("Query Single Error: %v\n", err)
		return nil, err
	}
	return balloonObject, nil
}

/**
 * This grabs all balloon objects within the balloon_objects_collection table.
 */
func GetAllBalloonObjects() ([]*BalloonObject, error) {
	var balloonObjects []*BalloonObject
	var db *pg.DB = Connect()
	defer db.Close()
	_, err := db.Query(&balloonObjects, `SELECT * FROM balloon_objects_collection`)
	if err != nil {
		log.Printf("Query Error: %v\n", err)
		return nil, err
	}
	return balloonObjects, nil
}

/**
 * This "deletes" a single balloon object by setting the "is_active" status to false.
 * This is a soft deletion.
 */
func DeleteSingle(id string) {
	balloonObject, _ := GetOneBalloonObject(id)
	balloonObject.IsActive = false
	var db *pg.DB = Connect()
	defer db.Close()
	balloonObject.Save(db)
	return
}
