package database_service

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	db_model "github.com/leechongyan/Studtor_backend/database_service/models"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// sqlite db
// this sqlite db has to implement all the methods which will be used by DatabaseConnector

type SQLiteDB struct {
	DatabaseFilename string
}

func (db *SQLiteDB) Init() {
	db.DatabaseFilename = "../studtor.db"

	// check if database already exists
	if _, err := os.Stat(db.DatabaseFilename); err == nil {
		// file exists
		log.Println("Database file " + db.DatabaseFilename + " exists.")
		log.Println("Loading existing database file...")
	} else if os.IsNotExist(err) {
		// file does not exist
		log.Println("Database file " + db.DatabaseFilename + " does not exist.")
		log.Println("Creating new database file...")

		// Initialize database
		log.Println("Migrating schema from GORM...")
		conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		// Migrate the schema
		conn.AutoMigrate(&db_model.User{})
		conn.AutoMigrate(&db_model.Course{})
		conn.AutoMigrate(&db_model.Tutor{})
		log.Println("Database initialized successfully.")
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		log.Fatal(err)
	}
}

// SaveUser saves the user to the database.
// Note that if the user is a new user, several fields
// would not have been initialized yet.
func (db SQLiteDB) SaveUser(user models.User) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_user db_model.User
	found := true

	// Get user if exists
	result := conn.First(&db_user, "email = ?", *user.Email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	// Update user
	db_user.First_name = *user.First_name
	db_user.Last_name = *user.Last_name
	db_user.Password = *user.Password
	db_user.Email = *user.Email
	db_user.User_type = *user.User_type
	if user.Token != nil {
		db_user.Token.String = *user.Token
		db_user.Token.Valid = true
	}
	if user.Refresh_token != nil {
		db_user.Refresh_token.String = *user.Refresh_token
		db_user.Refresh_token.Valid = true
	}
	if user.V_key != nil {
		db_user.V_key.String = *user.V_key
		db_user.V_key.Valid = true
	}

	if user.Verified {
		db_user.Verified = 1
	} else {
		db_user.Verified = 0
	}

	if user.Created_at.IsZero() {
		db_user.Created_at.Time = user.Created_at
		db_user.Created_at.Valid = true
	}
	if user.Updated_at.IsZero() {
		db_user.Updated_at.Time = user.Updated_at
		db_user.Updated_at.Valid = true
	}

	if found {
		// update existing record
		result = conn.Save(&db_user)
		if result.Error != nil {
			err = result.Error
			return
		}
	} else {
		// create new record
		result = conn.Create(&db_user)
		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return
}

// GetUser retrieves the user to the database.
func (db SQLiteDB) GetUser(email string) (user models.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_user db_model.User

	// Get user if exists
	result := conn.First(&db_user, "email = ?", email)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	user = models.User{}

	// Convert variables to their golang datatypes
	user.First_name = &db_user.First_name
	user.Last_name = &db_user.Last_name
	user.Password = &db_user.Password
	email_ref := email
	user.Email = &email_ref
	user.User_type = &db_user.User_type

	if db_user.Token.Valid {
		// token is not null
		user.Token = &db_user.Token.String
	}

	if db_user.Refresh_token.Valid {
		// refresh_token is not null
		user.Refresh_token = &db_user.Refresh_token.String
	}
	if db_user.V_key.Valid {
		// v_key is not null
		user.V_key = &db_user.V_key.String
	}
	if db_user.Verified == 0 {
		user.Verified = false
	} else {
		user.Verified = true

	}

	if db_user.Created_at.Valid {
		// Created_at is not null
		user.Created_at = db_user.Created_at.Time
	}
	if db_user.Updated_at.Valid {
		// Created_at is not null
		user.Updated_at = db_user.Updated_at.Time
	}

	return
}

func (db SQLiteDB) GetAllCourses(from string, size int) (courses []string, err error) {
	// convert from to int
	i, _ := strconv.Atoi(from)

	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_courses []db_model.Course
	courses = make([]string, 0)

	// fetch courses
	result := conn.Where("id >= ? AND id <= ?", i, i+size).Find(&db_courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// save db_courses to courses
	for i := 0; i < int(result.RowsAffected); i++ {
		courses = append(courses, db_courses[i].Course_code)
	}
	return
}

func (db SQLiteDB) GetAllTutors(from string, size int) (tutors []string, err error) {
	// convert from to int
	i, _ := strconv.Atoi(from)

	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_tutors []db_model.Tutor
	tutors = make([]string, 0)

	// fetch courses
	result := conn.Where("id >= ? AND id <= ?", i, i+size).Find(&db_tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// save db_courses to courses
	for i := 0; i < int(result.RowsAffected); i++ {
		tutors = append(tutors, db_tutors[i].Tutor_name)
	}
	return
}
