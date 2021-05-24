package controller

import (
	"errors"
	"log"
	"os"

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

// func (db SQLiteDB) GetAllCourses(db_options DB_options) (courses []interface{}, err error) {
// 	// create return object: courses
// 	courses = make([]interface{}, 0)

// 	// convert from string to int
// 	from_id, err := strconv.Atoi(*db_options.From_id)
// 	if err != nil {
// 		return
// 	}

// 	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
// 	if err != nil {
// 		// err opening database
// 		log.Println(err)
// 		return
// 	}

// 	var db_courses []db_model.Course
// 	var db_tutor_courses []db_model.TutorCourses
// 	var db_bookings []db_model.Booking

// 	// fetch courses
// 	result := conn.Where("id >= ? AND id <= ?", from_id, from_id+*db_options.Size).Find(&db_courses)
// 	if result.Error != nil {
// 		// row not found
// 		err = result.Error
// 		return
// 	}
// 	rows_affected := int(result.RowsAffected)

// 	// save db_courses to courses
// 	course := make(map[string]interface{})
// 	for i := 0; i < rows_affected; i++ {
// 		course["course_code"] = db_courses[i].CourseCode
// 		course["course_name"] = db_courses[i].CourseName

// 		// get number of tutors for course
// 		result = conn.Where("course_id = ?", db_courses[i].ID).Find(&db_tutor_courses)
// 		if result.Error != nil {
// 			// row not found
// 			err = result.Error
// 			return
// 		}
// 		course["n_tutors"] = result.RowsAffected

// 		// get number of students for course
// 		result = conn.Where("tutor = ?", db_courses[i].ID).Find(&db_tutor_courses)
// 		if result.Error != nil {
// 			// row not found
// 			err = result.Error
// 			return
// 		}
// 		course["n_tutors"] = result.RowsAffected

// 		courses = append(courses, db_courses[i].CourseCode)
// 	}
// 	return
// }

// func (db SQLiteDB) GetAllTutors(db_options DB_options) (tutors []string, err error) {
// 	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
// 	if err != nil {
// 		// err opening database
// 		log.Println(err)
// 		return
// 	}

// 	var db_tutors []db_model.Tutor
// 	tutors = make([]string, 0)

// 	// fetch courses
// 	result := conn.Where("id >= ? AND id <= ?", i, i+size).Find(&db_tutors)
// 	if result.Error != nil {
// 		// row not found
// 		err = result.Error
// 		return
// 	}

// 	// save db_courses to courses
// 	for i := 0; i < int(result.RowsAffected); i++ {
// 		name := db_tutors[i].FirstName + " " + db_tutors[i].LastName
// 		tutors = append(tutors, name)
// 	}
// 	return
// }

// func (db SQLiteDB) GetAllTutorsForACourse(courseID string, from string, size int) (tutors []string, err error) {
// 	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
// 	if err != nil {
// 		// err opening database
// 		log.Println(err)
// 		return
// 	}

// 	var db_user db_model.User
// 	found := true

// 	// Get user if exists
// 	result := conn.First(&db_user, "email = ?", *user.Email)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		// row not found
// 		found = false
// 	}

// 	// Update user
// 	db_user.FirstName = *user.First_name
// 	db_user.LastName = *user.Last_name
// 	db_user.Password = *user.Password
// 	db_user.Email = *user.Email
// 	db_user.UserType = *user.User_type
// 	if user.Token != nil {
// 		db_user.Token.String = *user.Token
// 		db_user.Token.Valid = true
// 	}
// 	if user.Refresh_token != nil {
// 		db_user.RefreshToken.String = *user.Refresh_token
// 		db_user.RefreshToken.Valid = true
// 	}
// 	if user.V_key != nil {
// 		db_user.VKey.String = *user.V_key
// 		db_user.VKey.Valid = true
// 	}

// 	if user.Verified {
// 		db_user.Verified = 1
// 	} else {
// 		db_user.Verified = 0
// 	}

// 	if user.Created_at.IsZero() {
// 		db_user.UserCreatedAt.Time = user.Created_at
// 		db_user.UserCreatedAt.Valid = true
// 	}
// 	if user.Updated_at.IsZero() {
// 		db_user.UserUpdatedAt.Time = user.Updated_at
// 		db_user.UserUpdatedAt.Valid = true
// 	}

// 	if found {
// 		// update existing record
// 		result = conn.Save(&db_user)
// 		if result.Error != nil {
// 			err = result.Error
// 			return
// 		}
// 	} else {
// 		// create new record
// 		result = conn.Create(&db_user)
// 		if result.Error != nil {
// 			err = result.Error
// 			return
// 		}
// 	}

// 	return
// }

// func (db SQLiteDB) SaveTutorAvailableTimes(slot tut_model.TimeFrame_query) (err error) {
// 	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
// 	if err != nil {
// 		// err opening database
// 		log.Println(err)
// 		return
// 	}

// 	var db_tutor_availabilities []db_model.Availability
// 	found := true

// 	// Get availability if exists
// 	result := conn.Where("email = ? AND available_from = ? AND available_to = ?", slot.Email, slot.From, slot.To).Find(&db_tutor_availabilities)
// 	if result.RowsAffected != 0 {
// 		// row not found
// 		found = false
// 	}

// 	// Update availability
// 	var db_availability db_model.Availability
// 	db_availability.Email = *slot.Email
// 	db_availability.AvailableFrom = slot.From
// 	db_availability.AvailableTo = slot.To

// 	if found {
// 		// If found, throw an error
// 		err = errors.New("This timeslot for this tutor has already been saved.")
// 		return
// 	} else {
// 		// create new record
// 		result = conn.Create(&db_availability)
// 		if result.Error != nil {
// 			err = result.Error
// 			return
// 		}
// 	}

// 	return
// }

// func (db SQLiteDB) DeleteTutorAvailableTimes(slot tut_model.TimeFrame_query) (err error) {
// 	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
// 	if err != nil {
// 		// err opening database
// 		log.Println(err)
// 		return
// 	}

// 	// Batch delete availability
// 	result := conn.Where("email = ? AND available_from = ? AND available_to = ?", slot.Email, slot.From, slot.To).Delete(db_model.Availability{})
// 	if result.RowsAffected == 0 {

// 		err = errors.New("This timeslot for this tutor has already been saved.")
// 	}
// 	log.Printf("%d rows were deleted.\n", result.RowsAffected)

// 	return
// }

// func (db SQLiteDB) GetTutorBookedTimes(slot tut_model.TimeFrame_query) (bookedTimes Timeslots, err error) {
// 	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
// 	if err != nil {
// 		// err opening database
// 		log.Println(err)
// 		return
// 	}

// 	bookedTimes = make(Timeslots)

// 	var db_bookings []db_model.Booking
// 	var db_availability []db_model.Availability

// 	// Get all availability for tutor
// 	result := conn.Where("email = ?", slot.Email).Find(&db_availability)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			// no record found, just return
// 			return
// 		}
// 		// row not found
// 		err = result.Error
// 		return
// 	}

// 	// should have course code for the booked time slots as well
// 	slots := make(map[string][]time.Time)
// 	slots["CZ1003"] = []time.Time{time.Now(), time.Now()}
// 	slots["CZ1004"] = []time.Time{time.Now(), time.Now()}
// 	bookedTimes = make(Timeslots)
// 	bookedTimes["first_name"] = "Jeff"
// 	bookedTimes["last_name"] = "Lee"
// 	bookedTimes["email"] = "clee051@e.ntu.edu.sg"
// 	bookedTimes["time_slots"] = slots
// 	return
// }

// func (db SQLiteDB) GetTutorAvailableTimes(slot tut_model.TimeFrame_query) (availableTimes Timeslots, err error) {
// 	// extract from
// 	// query database from and to
// 	// from := slot.From
// 	// to := slot.To

// 	// create 10 timeslots for testing
// 	slots := make([][]time.Time, 10)
// 	for i := range slots {
// 		slots[i] = []time.Time{time.Now(), time.Now()}
// 	}
// 	availableTimes = make(Timeslots)

// 	availableTimes["first_name"] = "Jeff"
// 	availableTimes["last_name"] = "Lee"
// 	availableTimes["email"] = "clee051@e.ntu.edu.sg"
// 	availableTimes["time_slots"] = slots
// 	return
// }

// func (db SQLiteDB) BookTutorTime(student_email string, slot tut_model.TimeFrame_query) (err error) {
// 	return nil
// }

// func (db SQLiteDB) UnBookTutorTime(student_email string, slot tut_model.TimeFrame_query) (err error) {
// 	return nil
// }

// func (db SQLiteDB) GetStudentBookedTimes(slot tut_model.TimeFrame_query) (bookedTimes Timeslots, err error) {
// 	// should have course code for the booked time slots as well
// 	slots := make(map[string][]time.Time)
// 	slots["CZ1003"] = []time.Time{time.Now(), time.Now()}
// 	slots["CZ1004"] = []time.Time{time.Now(), time.Now()}
// 	bookedTimes = make(Timeslots)
// 	bookedTimes["first_name"] = "Jeff"
// 	bookedTimes["last_name"] = "Lee"
// 	bookedTimes["email"] = "clee051@e.ntu.edu.sg"
// 	bookedTimes["time_slots"] = slots
// 	return
// }
