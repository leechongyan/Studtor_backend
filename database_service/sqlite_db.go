package database_service

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
)

// sqlite db
// this sqlite db has to implement all the methods which will be used by DatabaseConnector

type SQLiteDB struct {
	DatabaseFilename string
}

func (db *SQLiteDB) Init() {
	db.DatabaseFilename = "./build/studtor.db"
	initDatabaseScript := "./scripts/init-sqlite-db.sql"

	// check if database already exists
	if _, err := os.Stat(db.DatabaseFilename); err == nil {
		// file exists
		log.Println("Database file " + db.DatabaseFilename + " exists.")
		log.Println("Loading existing database file...")
	} else if os.IsNotExist(err) {
		// file does not exist
		log.Println("Database file " + db.DatabaseFilename + " does not exist.")
		log.Println("Creating new database file...")

		// Run database initialization scripts
		log.Println("Reading database initialization scripts from " + initDatabaseScript + "...")
		c, ioErr := ioutil.ReadFile(initDatabaseScript)
		if ioErr != nil {
			// err reading file
			log.Fatal("err reading database initialization scripts.")
		}

		// Initialize database
		conn, err := sql.Open("sqlite3", db.DatabaseFilename)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		stmt := string(c)
		_, err = conn.Exec(stmt)
		if err != nil {
			// handle err.
			log.Fatal(err)
		}
		log.Println("Database initialized successfully.")
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		log.Fatal(err)
	}
}

func (db SQLiteDB) SaveUser(user models.User) (err error) {
	// try opening database
	conn, err := sql.Open("sqlite3", db.DatabaseFilename)
	if err != nil {
		// err opening database
		log.Fatal(err)
		return
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		// err starting connection
		log.Fatal(err)
		return
	}

	stmt, err := tx.Prepare(
		"INSERT INTO students(first_name,last_name,password,email,token,user_type,refresh_token,v_key,verified,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = stmt.Exec(
		user.First_name,
		user.Last_name,
		user.Password,
		user.Email,
		user.Token,
		user.User_type,
		user.Refresh_token,
		user.V_key,
		convertBooleanToInt(user.Verified),
		user.Created_at.String(),
		user.Updated_at.String())
	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()
	tx.Commit()

	return
}

func (db SQLiteDB) GetUser(email string) (user models.User, err error) {
	// try opening database
	conn, err := sql.Open("sqlite3", db.DatabaseFilename)
	if err != nil {
		// err opening database
		log.Fatal(err)
		return
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}

	var verified string
	var created_at string
	var updated_at string

	stmt, err := tx.Prepare("SELECT first_name,last_name,password,token,user_type,refresh_token,v_key,verified,created_at,updated_at FROM users where email = ?")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	user = models.User{}
	err = stmt.QueryRow(email).Scan(
		&user.First_name,
		&user.Last_name,
		&user.Password,
		&user.Token,
		&user.User_type,
		&user.Refresh_token,
		&user.V_key,
		&verified,
		&created_at,
		&updated_at)
	if err != nil {
		log.Println(err)
		return
	}

	// Convert variables to their golang datatypes
	i, err := strconv.Atoi(verified)
	user.Verified = convertIntToBool(i)

	// Use default format: https://golang.org/pkg/time/#Time.String
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	user.Created_at, err = time.Parse(layout, created_at)
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Updated_at, err = time.Parse(layout, updated_at)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (db SQLiteDB) GetAllCourses(from string, size int) (courses []string, err error) {
	// convert from to int
	i, _ := strconv.Atoi(from)

	// try opening database
	conn, err := sql.Open("sqlite3", db.DatabaseFilename)
	if err != nil {
		// err opening database
		log.Fatal(err)
		return
	}
	defer conn.Close()

	courses = make([]string, 0)

	// fetch courses
	stmt := fmt.Sprintf("SELECT course_code FROM courses WHERE course_id >= %d AND course_id <= %d", i, i+size)
	rows, err := conn.Query(stmt)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var course_code string
		err = rows.Scan(&course_code)
		if err != nil {
			log.Println(err)
			return
		}
		courses = append(courses, course_code)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (db SQLiteDB) GetAllTutors(from string, size int) (tutors []string, err error) {
	// convert from to int
	i, _ := strconv.Atoi(from)

	// try opening database
	conn, err := sql.Open("sqlite3", db.DatabaseFilename)
	if err != nil {
		// err opening database
		log.Fatal(err)
		return
	}
	defer conn.Close()

	tutors = make([]string, 0)

	// fetch courses
	stmt := fmt.Sprintf("SELECT tutor_name FROM tutors WHERE tutor_id >= %d AND tutor_id <= %d", i, i+size)
	rows, err := conn.Query(stmt)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tutor_name string
		err = rows.Scan(&tutor_name)
		if err != nil {
			log.Println(err)
			return
		}
		tutors = append(tutors, tutor_name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func convertBooleanToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func convertIntToBool(i int) bool {
	if i != 0 {
		return true
	} else {
		return false
	}
}
