package controller

import (
	"errors"
	"log"
	"os"
	"time"

	db_model "github.com/leechongyan/Studtor_backend/database_service/models"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO:Commented snippet below describes how to change a db_model user model to an auth_model user model. To be used as reference, as we migrate to db_model.
// user = models.User{}

// // Convert variables to their golang datatypes
// user.First_name = &db_user.FirstName
// user.Last_name = &db_user.LastName
// user.Password = &db_user.Password
// email_ref := email
// user.Email = &email_ref
// user.User_type = &db_user.UserType

// if db_user.Token.Valid {
// 	// token is not null
// 	user.Token = &db_user.Token.String
// }

// if db_user.RefreshToken.Valid {
// 	// refresh_token is not null
// 	user.Refresh_token = &db_user.RefreshToken.String
// }
// if db_user.VKey.Valid {
// 	// v_key is not null
// 	user.V_key = &db_user.VKey.String
// }
// if db_user.Verified == 0 {
// 	user.Verified = false
// } else {
// 	user.Verified = true

// }

// if db_user.UserCreatedAt.Valid {
// 	// Created_at is not null
// 	user.Created_at = db_user.UserCreatedAt.Time
// }
// if db_user.UserUpdatedAt.Valid {
// 	// Created_at is not null
// 	user.Updated_at = db_user.UserUpdatedAt.Time
// }

type SQLiteDB struct {
	DatabaseFilename string
}

func (db *SQLiteDB) Init() {
	db.DatabaseFilename = "./build/studtor.db"

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
		err = db.migrateSchema()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Database initialized successfully.")
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		log.Fatal(err)
	}
}

// migrateSchema instantiates the database tables following the schema defined by our models.
func (db *SQLiteDB) migrateSchema() (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		return
	}
	// Migrate the schema
	conn.AutoMigrate(&db_model.User{})
	conn.AutoMigrate(&db_model.Course{})
	conn.AutoMigrate(&db_model.Tutor{})
	conn.AutoMigrate(&db_model.Availability{})
	conn.AutoMigrate(&db_model.Booking{})
	conn.AutoMigrate(&db_model.TutorCourses{})
	return nil
}

func (db *SQLiteDB) GetUsers() (users []db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Find(&users)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetUserById(user_id int) (user db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.First(&user, "id = ?", user_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetUserByEmail(email string) (user db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.First(&user, "email = ?", email)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) SaveUser(user db_model.User) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_user db_model.User
	var found bool
	found = true

	// Get user if exists
	result := conn.First(&db_user, "email = ?", user.Email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	// Update user fields
	db_user.FirstName = user.FirstName
	db_user.LastName = user.LastName
	db_user.Password = user.Password
	db_user.Email = user.Email
	db_user.UserType = user.UserType

	// Update optional fields
	if user.Token.Valid {
		db_user.Token.String = user.Token.String
		db_user.Token.Valid = true
	}
	if user.RefreshToken.Valid {
		db_user.RefreshToken.String = user.RefreshToken.String
		db_user.RefreshToken.Valid = true
	}
	if user.VKey.Valid {
		db_user.VKey.String = user.VKey.String
		db_user.VKey.Valid = true
	}

	if user.Verified == 0 {
		db_user.Verified = 0
	} else {
		db_user.Verified = 1
	}

	if user.UserCreatedAt.Valid {
		db_user.UserCreatedAt.Time = user.UserCreatedAt.Time
		db_user.UserCreatedAt.Valid = true
	}
	if user.UserUpdatedAt.Valid {
		db_user.UserUpdatedAt.Time = user.UserUpdatedAt.Time
		db_user.UserUpdatedAt.Valid = true
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

func (db SQLiteDB) DeleteUserById(user_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Delete(&db_model.User{}, user_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) DeleteUserByEmail(email string) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Where("email = ?", email).Delete(db_model.User{})
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetCoursesForTutor(tutor_id int) (courses []db_model.Course, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("tutor_id = ?", tutor_id).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of course ids
	course_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		course_ids = append(course_ids, int(tutorCourses[i].CourseID))
	}

	// Using tutor_id, get all courses
	result = conn.Where("id IN ?", course_ids).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourse(course_id int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", course_id).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of course ids
	tutor_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutor_ids = append(tutor_ids, int(tutorCourses[i].TutorID))
	}

	// Using tutor_id, get all courses
	result = conn.Where("id IN ?", tutor_ids).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourseFromIdOfSize(course_id int, tut_id int, size int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", course_id).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of course ids
	tutor_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutor_ids = append(tutor_ids, int(tutorCourses[i].TutorID))
	}

	// Using tutor_id, get all courses
	result = conn.Where("id IN ? AND id >= ?", tutor_ids, tut_id).Limit(size).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourseFromId(course_id, tut_id int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", course_id).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of course ids
	tutor_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutor_ids = append(tutor_ids, int(tutorCourses[i].TutorID))
	}

	// Using tutor_id, get all courses
	result = conn.Where("id IN ? AND id >= ?", tutor_ids, tut_id).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourseOfSize(course_id, size int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", course_id).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of course ids
	tutor_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutor_ids = append(tutor_ids, int(tutorCourses[i].TutorID))
	}

	// Using tutor_id, get all courses
	result = conn.Where("id IN ?", tutor_ids).Limit(size).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) SaveTutorCourse(tutor_id int, course_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_tutor_course db_model.TutorCourses
	var found bool
	found = true

	// Get user if exists
	result := conn.Where("tutor_id = ? AND course_id = ?", tutor_id, course_id).First(&db_tutor_course)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	if found {
		// update existing record
		result = conn.Save(&db_tutor_course)
		if result.Error != nil {
			err = result.Error
			return
		}
	} else {
		// create new record
		result = conn.Create(&db_tutor_course)
		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return
}

func (db SQLiteDB) DeleteTutorCourse(tutor_id int, course_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Where("tutor_id = ? AND course_id = ?", tutor_id, course_id).Delete(db_model.TutorCourses{})
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetCourse(course_id int) (course db_model.Course, n_students int, n_tutors int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get course if exists
	result := conn.First(&course, "id = ?", course_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// Get n_students for course
	var count int64
	result = conn.Model(&db_model.Booking{}).Where("course_id = ?", course_id).Count(&count)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}
	n_students = int(count)

	// Get n_tutors for course
	result = conn.Model(&db_model.TutorCourses{}).Where("course_id = ?", course_id).Count(&count)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}
	n_tutors = int(count)

	return
}

func (db SQLiteDB) GetCourses() (courses []db_model.Course, n_students []int, n_tutors []int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all courses
	result := conn.Order("course_code ASC").Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// Raw SQL for count group by query
	booking_map := make(map[int]int)
	var course_id int
	var count int
	rows, err := conn.Raw("SELECT course_id, COUNT(*) FROM bookings GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		booking_map[course_id] = count
	}

	// Raw SQL for count group by query
	tutor_map := make(map[int]int)
	rows, err = conn.Raw("SELECT course_id, COUNT(*) FROM tutor_courses GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		tutor_map[course_id] = count
	}

	// format booking map and tutor courses map for results
	for i := 0; i < len(courses); i++ {
		n_students = append(n_students, booking_map[int(courses[i].ID)])
		n_tutors = append(n_tutors, tutor_map[int(courses[i].ID)])
	}

	return
}

func (db SQLiteDB) GetCoursesForSchool(school_code string) (courses []db_model.Course, n_students []int, n_tutors []int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all courses for school
	result := conn.Where("course_code LIKE ?", school_code+"%").Order("course_code ASC").Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// Raw SQL for count group by query
	booking_map := make(map[int]int)
	var course_id int
	var count int
	rows, err := conn.Raw("SELECT course_id, COUNT(*) FROM bookings GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		booking_map[course_id] = count
	}

	// Raw SQL for count group by query
	tutor_map := make(map[int]int)
	rows, err = conn.Raw("SELECT course_id, COUNT(*) FROM tutor_courses GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		tutor_map[course_id] = count
	}

	// format booking map and tutor courses map for results
	for i := 0; i < len(courses); i++ {
		n_students = append(n_students, booking_map[int(courses[i].ID)])
		n_tutors = append(n_tutors, tutor_map[int(courses[i].ID)])
	}

	return
}

func (db SQLiteDB) GetCoursesForSchoolWithOffset(school_code string, offset int) (courses []db_model.Course, n_students []int, n_tutors []int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all courses for school with offset
	result := conn.Where("course_code LIKE ?", school_code+"%").Order("course_code ASC").Offset(offset).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// Raw SQL for count group by query
	booking_map := make(map[int]int)
	var course_id int
	var count int
	rows, err := conn.Raw("SELECT course_id, COUNT(*) FROM bookings GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		booking_map[course_id] = count
	}

	// Raw SQL for count group by query
	tutor_map := make(map[int]int)
	rows, err = conn.Raw("SELECT course_id, COUNT(*) FROM tutor_courses GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		tutor_map[course_id] = count
	}

	// format booking map and tutor courses map for results
	for i := 0; i < len(courses); i++ {
		n_students = append(n_students, booking_map[int(courses[i].ID)])
		n_tutors = append(n_tutors, tutor_map[int(courses[i].ID)])
	}

	return
}

func (db SQLiteDB) GetCoursesForSchoolOfSize(school_code string, size int) (courses []db_model.Course, n_students []int, n_tutors []int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all courses for school with limit
	result := conn.Where("course_code LIKE ?", school_code+"%").Order("course_code ASC").Limit(size).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// Raw SQL for count group by query
	booking_map := make(map[int]int)
	var course_id int
	var count int
	rows, err := conn.Raw("SELECT course_id, COUNT(*) FROM bookings GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		booking_map[course_id] = count
	}

	// Raw SQL for count group by query
	tutor_map := make(map[int]int)
	rows, err = conn.Raw("SELECT course_id, COUNT(*) FROM tutor_courses GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		tutor_map[course_id] = count
	}

	// format booking map and tutor courses map for results
	for i := 0; i < len(courses); i++ {
		n_students = append(n_students, booking_map[int(courses[i].ID)])
		n_tutors = append(n_tutors, tutor_map[int(courses[i].ID)])
	}

	return
}
func (db SQLiteDB) GetCoursesForSchoolOfSizeWithOffset(school_code string, offset int, size int) (courses []db_model.Course, n_students []int, n_tutors []int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all courses for school with offset and limit
	result := conn.Where("course_code LIKE ?", school_code+"%").Order("course_code ASC").Offset(offset).Limit(size).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// Raw SQL for count group by query
	booking_map := make(map[int]int)
	var course_id int
	var count int
	rows, err := conn.Raw("SELECT course_id, COUNT(*) FROM bookings GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		booking_map[course_id] = count
	}

	// Raw SQL for count group by query
	tutor_map := make(map[int]int)
	rows, err = conn.Raw("SELECT course_id, COUNT(*) FROM tutor_courses GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&course_id, &count)
		tutor_map[course_id] = count
	}

	// format booking map and tutor courses map for results
	for i := 0; i < len(courses); i++ {
		n_students = append(n_students, booking_map[int(courses[i].ID)])
		n_tutors = append(n_tutors, tutor_map[int(courses[i].ID)])
	}

	return
}

func (db SQLiteDB) GetTutors() (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all tutors
	result := conn.Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetTutorsFromId(tutor_id int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all tutors from id
	result := conn.Where("id >= ?", tutor_id).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetTutorsFromIdOfSize(tutor_id int, size int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all tutors from id with size
	result := conn.Where("id >= ?", tutor_id).Limit(size).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetTutorsOfSize(size int) (tutors []db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all tutors from start with size
	result := conn.Limit(size).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetTutorById(tutor_id int) (tutor db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get tutor if exists
	result := conn.First(&tutor, "id = ?", tutor_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetTutorByEmail(email string) (tutor db_model.Tutor, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get tutor if exists
	result := conn.First(&tutor, "email = ?", email)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) SaveTutor(tutor db_model.Tutor) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_tutor db_model.Tutor
	var found bool
	found = true

	// Get user if exists
	result := conn.First(&db_tutor, "id = ?", tutor.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	// Update user fields
	db_tutor.FirstName = tutor.FirstName
	db_tutor.LastName = tutor.LastName
	db_tutor.Email = tutor.Email

	if found {
		// update existing record
		result = conn.Save(&db_tutor)
		if result.Error != nil {
			err = result.Error
			return
		}
	} else {
		// create new record
		result = conn.Create(&db_tutor)
		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return
}
func (db SQLiteDB) DeleteTutorById(tutor_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Delete tutor if exists
	result := conn.Delete(&db_model.Tutor{}, tutor_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) DeleteTutorByEmail(email string) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Delete tutor if exists
	result := conn.Where("email LIKE ?", email).Delete(db_model.Tutor{})
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetBookingsById(user_id int) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", user_id).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of availability_ids and course_ids
	availability_ids := make([]int, 0)
	course_ids := make([]int, 0)
	user_ids := make([]int, 0)

	for i := 0; i < int(result.RowsAffected); i++ {
		availability_ids = append(availability_ids, int(bookings_no_details[i].TutorAvailabilityID))
		course_ids = append(course_ids, int(bookings_no_details[i].CourseID))
		user_ids = append(user_ids, int(bookings_no_details[i].UserID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? ", availability_ids).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get tutors corresponding to availabilities
	tutor_ids := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutor_ids = append(tutor_ids, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", course_ids).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", user_ids).Find(&users)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding tutors
	var tutors []db_model.Tutor
	result = conn.Where("id IN ? ", tutor_ids).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.Tutor)
	for i := 0; i < len(bookings_no_details); i++ {
		users_map[int(users[i].ID)] = users[i]
		courses_map[int(courses[i].ID)] = courses[i]
		availability_map[int(availabilities[i].ID)] = availabilities[i]
		tutors_map[int(tutors[i].ID)] = tutors[i]
	}

	// format booking details to be returned
	var booking_details db_model.BookingDetails
	for i := 0; i < len(bookings_no_details); i++ {
		booking_details.StudentName = users_map[bookings_no_details[i].UserID].FirstName + " " + users_map[bookings_no_details[i].UserID].LastName
		booking_details.CourseCode = courses_map[bookings_no_details[i].CourseID].CourseCode
		booking_details.TutorName = tutors_map[int(availability_map[bookings_no_details[i].TutorAvailabilityID].TutorID)].FirstName + " " + tutors_map[int(availability_map[bookings_no_details[i].TutorAvailabilityID].TutorID)].LastName
		booking_details.FromTime = availability_map[bookings_no_details[i].TutorAvailabilityID].AvailableFrom
		booking_details.ToTime = availability_map[bookings_no_details[i].TutorAvailabilityID].AvailableTo

		// append to result
		bookings = append(bookings, booking_details)
	}

	return
}
func (db SQLiteDB) GetBookingsByIdFrom(user_id int, from_time time.Time) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", user_id).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of availability_ids
	availability_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		availability_ids = append(availability_ids, int(bookings_no_details[i].TutorAvailabilityID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? AND available_from >= ?", availability_ids, from_time).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// modify bookings_no_details after filter
	var bookings_no_details_after_filter []db_model.Booking
	for i := 0; i < len(bookings_no_details); i++ {
		for j := 0; j < len(availabilities); j++ {
			if availabilities[j].ID == uint(bookings_no_details[i].TutorAvailabilityID) {
				bookings_no_details_after_filter = append(bookings_no_details_after_filter, bookings_no_details[i])
				break
			}
		}
	}

	// create list of course_ids and user_ids
	course_ids := make([]int, 0)
	user_ids := make([]int, 0)
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		course_ids = append(course_ids, int(bookings_no_details_after_filter[i].CourseID))
		user_ids = append(user_ids, int(bookings_no_details_after_filter[i].UserID))
	}

	// get tutors corresponding to availabilities
	tutor_ids := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutor_ids = append(tutor_ids, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", course_ids).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", user_ids).Find(&users)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding tutors
	var tutors []db_model.Tutor
	result = conn.Where("id IN ? ", tutor_ids).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.Tutor)
	for i := 0; i < len(bookings_no_details); i++ {
		users_map[int(users[i].ID)] = users[i]
		courses_map[int(courses[i].ID)] = courses[i]
		availability_map[int(availabilities[i].ID)] = availabilities[i]
		tutors_map[int(tutors[i].ID)] = tutors[i]
	}

	// format booking details to be returned
	var booking_details db_model.BookingDetails
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		// check if
		booking_details.StudentName = users_map[bookings_no_details_after_filter[i].UserID].FirstName + " " + users_map[bookings_no_details[i].UserID].LastName
		booking_details.CourseCode = courses_map[bookings_no_details_after_filter[i].CourseID].CourseCode
		booking_details.TutorName = tutors_map[int(availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].TutorID)].FirstName + " " + tutors_map[int(availability_map[bookings_no_details[i].TutorAvailabilityID].TutorID)].LastName
		booking_details.FromTime = availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].AvailableFrom
		booking_details.ToTime = availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].AvailableTo

		// append to result
		bookings = append(bookings, booking_details)
	}

	return
}
func (db SQLiteDB) GetBookingsByIdTo(user_id int, to_time time.Time) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", user_id).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of availability_ids
	availability_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		availability_ids = append(availability_ids, int(bookings_no_details[i].TutorAvailabilityID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? AND available_to <= ?", availability_ids, to_time).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// modify bookings_no_details after filter
	var bookings_no_details_after_filter []db_model.Booking
	for i := 0; i < len(bookings_no_details); i++ {
		for j := 0; j < len(availabilities); j++ {
			if availabilities[j].ID == uint(bookings_no_details[i].TutorAvailabilityID) {
				bookings_no_details_after_filter = append(bookings_no_details_after_filter, bookings_no_details[i])
				break
			}
		}
	}

	// create list of course_ids and user_ids
	course_ids := make([]int, 0)
	user_ids := make([]int, 0)
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		course_ids = append(course_ids, int(bookings_no_details_after_filter[i].CourseID))
		user_ids = append(user_ids, int(bookings_no_details_after_filter[i].UserID))
	}

	// get tutors corresponding to availabilities
	tutor_ids := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutor_ids = append(tutor_ids, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", course_ids).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", user_ids).Find(&users)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding tutors
	var tutors []db_model.Tutor
	result = conn.Where("id IN ? ", tutor_ids).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.Tutor)
	for i := 0; i < len(bookings_no_details); i++ {
		users_map[int(users[i].ID)] = users[i]
		courses_map[int(courses[i].ID)] = courses[i]
		availability_map[int(availabilities[i].ID)] = availabilities[i]
		tutors_map[int(tutors[i].ID)] = tutors[i]
	}

	// format booking details to be returned
	var booking_details db_model.BookingDetails
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		// check if
		booking_details.StudentName = users_map[bookings_no_details_after_filter[i].UserID].FirstName + " " + users_map[bookings_no_details[i].UserID].LastName
		booking_details.CourseCode = courses_map[bookings_no_details_after_filter[i].CourseID].CourseCode
		booking_details.TutorName = tutors_map[int(availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].TutorID)].FirstName + " " + tutors_map[int(availability_map[bookings_no_details[i].TutorAvailabilityID].TutorID)].LastName
		booking_details.FromTime = availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].AvailableFrom
		booking_details.ToTime = availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].AvailableTo

		// append to result
		bookings = append(bookings, booking_details)
	}
	return
}
func (db SQLiteDB) GetBookingsByIdFromTo(user_id int, from_time time.Time, to_time time.Time) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", user_id).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// create list of availability_ids
	availability_ids := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		availability_ids = append(availability_ids, int(bookings_no_details[i].TutorAvailabilityID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? AND available_from >= ? AND available_to <= ?", availability_ids, from_time, to_time).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// modify bookings_no_details after filter
	var bookings_no_details_after_filter []db_model.Booking
	for i := 0; i < len(bookings_no_details); i++ {
		for j := 0; j < len(availabilities); j++ {
			if availabilities[j].ID == uint(bookings_no_details[i].TutorAvailabilityID) {
				bookings_no_details_after_filter = append(bookings_no_details_after_filter, bookings_no_details[i])
				break
			}
		}
	}

	// create list of course_ids and user_ids
	course_ids := make([]int, 0)
	user_ids := make([]int, 0)
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		course_ids = append(course_ids, int(bookings_no_details_after_filter[i].CourseID))
		user_ids = append(user_ids, int(bookings_no_details_after_filter[i].UserID))
	}

	// get tutors corresponding to availabilities
	tutor_ids := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutor_ids = append(tutor_ids, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", course_ids).Find(&courses)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", user_ids).Find(&users)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// get corresponding tutors
	var tutors []db_model.Tutor
	result = conn.Where("id IN ? ", tutor_ids).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.Tutor)
	for i := 0; i < len(bookings_no_details); i++ {
		users_map[int(users[i].ID)] = users[i]
		courses_map[int(courses[i].ID)] = courses[i]
		availability_map[int(availabilities[i].ID)] = availabilities[i]
		tutors_map[int(tutors[i].ID)] = tutors[i]
	}

	// format booking details to be returned
	var booking_details db_model.BookingDetails
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		// check if
		booking_details.StudentName = users_map[bookings_no_details_after_filter[i].UserID].FirstName + " " + users_map[bookings_no_details[i].UserID].LastName
		booking_details.CourseCode = courses_map[bookings_no_details_after_filter[i].CourseID].CourseCode
		booking_details.TutorName = tutors_map[int(availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].TutorID)].FirstName + " " + tutors_map[int(availability_map[bookings_no_details[i].TutorAvailabilityID].TutorID)].LastName
		booking_details.FromTime = availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].AvailableFrom
		booking_details.ToTime = availability_map[bookings_no_details_after_filter[i].TutorAvailabilityID].AvailableTo

		// append to result
		bookings = append(bookings, booking_details)
	}
	return
}
func (db SQLiteDB) SaveBooking(availability_id int, user_id int, course_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_booking db_model.Booking
	var found bool
	found = true

	// Get user if exists
	result := conn.Where("tutor_availability_id = ? AND user_id = ? AND course_id = ?", availability_id, user_id, course_id).First(&db_booking)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	// Update user fields
	db_booking.TutorAvailabilityID = availability_id
	db_booking.UserID = user_id
	db_booking.CourseID = course_id

	if found {
		// update existing record
		result = conn.Save(&db_booking)
		if result.Error != nil {
			err = result.Error
			return
		}
	} else {
		// create new record
		result = conn.Create(&db_booking)
		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return
}
func (db SQLiteDB) DeleteBookingById(booking_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Delete(&db_model.Booking{}, booking_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}

func (db SQLiteDB) GetAvailabilityById(tutor_id int) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ?", tutor_id).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetAvailabilityByIdFrom(tutor_id int, from_time time.Time) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ? AND available_from >= ?", tutor_id, from_time).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetAvailabilityByIdTo(tutor_id int, to_time time.Time) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ? AND available_to <= ?", tutor_id, to_time).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) GetAvailabilityByIdFromTo(tutor_id int, from_time time.Time, to_time time.Time) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ? AND available_from >= ? AND available_to = ?", tutor_id, from_time, to_time).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
func (db SQLiteDB) SaveTutorAvailability(tutor_id int, from_time time.Time, to_time time.Time) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_availability db_model.Availability
	var found bool
	found = true

	// Get user if exists
	result := conn.Where("tutor_id = ? AND available_from = ? AND available_to = ?", tutor_id, from_time, to_time).First(&db_availability)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	// Update user fields
	db_availability.TutorID = uint(tutor_id)
	db_availability.AvailableFrom = from_time
	db_availability.AvailableTo = to_time

	if found {
		// update existing record
		result = conn.Save(&db_availability)
		if result.Error != nil {
			err = result.Error
			return
		}
	} else {
		// create new record
		result = conn.Create(&db_availability)
		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return
}
func (db SQLiteDB) DeleteTutorAvailabilityById(availability_id int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Delete(&db_model.Availability{}, availability_id)
	if result.Error != nil {
		// row not found
		err = result.Error
		return
	}

	return
}
