package sqlitecontroller

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// current implementation is for sqlite

func SelectStudentByID(id string) (string, error) {
	// try opening database; if failed, throw errow
	db, err := sql.Open("sqlite3", "./build/studtor.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT student_name FROM students where student_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var name string
	err = stmt.QueryRow(id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name, nil
}

func InsertStudent(name string) error {
	// try opening database; if failed, throw errow
	db, err := sql.Open("sqlite3", "./build/studtor.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO students(student_name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = stmt.Exec(name)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	tx.Commit()

	return nil
}
