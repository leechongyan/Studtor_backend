package controller

import (
	"errors"
	"log"
	"os"
	"time"

	database_errors "github.com/leechongyan/Studtor_backend/constants/errors/database_errors"
	db_model "github.com/leechongyan/Studtor_backend/database_service/models"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
	conn.AutoMigrate(&db_model.School{})
	conn.AutoMigrate(&db_model.SchoolCourses{})
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
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetUserByID(userID int) (user db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.First(&user, "id = ?", userID)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
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
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) CreateUser(user db_model.User) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// user model input should not have an id set!
	if user.ID != 0 {
		err = database_errors.ErrInvalidParameters
		return
	}

	var db_user db_model.User

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

	db_user.Verified = user.Verified

	if !user.UserCreatedAt.IsZero() {
		db_user.UserCreatedAt = user.UserCreatedAt
	}
	if !user.UserUpdatedAt.IsZero() {
		db_user.UserUpdatedAt = user.UserUpdatedAt
	}

	// create user record
	result := conn.Create(&db_user)
	if result.Error != nil {
		err = database_errors.ErrCreateRecordFailed
		return
	}

	return
}

func (db SQLiteDB) UpdateUser(user db_model.User) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// user model input must have an id set!
	if user.ID == 0 {
		err = database_errors.ErrInvalidParameters
		return
	}

	var db_user db_model.User
	var found bool
	found = true

	// Get user if exists
	result := conn.First(&db_user, "id = ?", user.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	// if not found, throw error!
	if !found {
		err = database_errors.ErrRecordToBeUpdatedNotFound
		return
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

	db_user.Verified = user.Verified

	if !user.UserCreatedAt.IsZero() {
		db_user.UserCreatedAt = user.UserCreatedAt
	}
	if !user.UserUpdatedAt.IsZero() {
		db_user.UserUpdatedAt = user.UserUpdatedAt
	}

	// update existing record
	result = conn.Save(&db_user)
	if result.Error != nil {
		err = database_errors.ErrUpdateRecordFailed
		return
	}

	return
}

func (db SQLiteDB) DeleteUserByID(userID int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Delete(&db_model.User{}, userID)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrDeleteRecordFailed
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
		err = database_errors.ErrDeleteRecordFailed
		return
	}

	return
}

func (db SQLiteDB) GetCoursesForTutor(tutorID int) (courses []db_model.Course, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("tutor_id = ?", tutorID).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of course IDs
	courseIDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		courseIDs = append(courseIDs, int(tutorCourses[i].CourseID))
	}

	// Using tutorID, get all courses
	result = conn.Where("id IN ?", courseIDs).Find(&courses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourse(courseID int) (tutors []db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("courseID = ?", courseID).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of course IDs
	tutorIDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutorIDs = append(tutorIDs, int(tutorCourses[i].TutorID))
	}

	// Using tutorID, get all courses
	result = conn.Where("ID IN ?", tutorIDs).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourseFromIDOfSize(courseID int, tutorID int, size int) (tutors []db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", courseID).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of course IDs
	tutorIDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutorIDs = append(tutorIDs, int(tutorCourses[i].TutorID))
	}

	// Using tutorID, get all courses
	result = conn.Where("id IN ? AND id >= ?", tutorIDs, tutorID).Limit(size).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourseFromID(courseID, tutorID int) (tutors []db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", courseID).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of course IDs
	tutorIDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutorIDs = append(tutorIDs, int(tutorCourses[i].TutorID))
	}

	// Using tutorID, get all courses
	result = conn.Where("id IN ? AND id >= ?", tutorIDs, tutorID).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetTutorsForCourseOfSize(courseID, size int) (tutors []db_model.User, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}
	var tutorCourses []db_model.TutorCourses

	// Get all tutor courses
	result := conn.Where("course_id = ?", courseID).Find(&tutorCourses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of course IDs
	tutorIDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		tutorIDs = append(tutorIDs, int(tutorCourses[i].TutorID))
	}

	// Using tutorID, get all courses
	result = conn.Where("id IN ?", tutorIDs).Limit(size).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) CreateTutorCourse(tutorID int, courseID int) (err error) {
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
	result := conn.Where("tutor_id = ? AND course_id = ?", tutorID, courseID).First(&db_tutor_course)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	if found {
		// do nothing
		return
	} else {
		// create new record
		result = conn.Create(&db_tutor_course)
		if result.Error != nil {
			err = database_errors.ErrCreateRecordFailed
			return
		}
	}

	return
}

func (db SQLiteDB) DeleteTutorCourse(tutorID int, courseID int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Where("tutorID = ? AND courseID = ?", tutorID, courseID).Delete(db_model.TutorCourses{})
	if result.Error != nil {
		// row not found
		err = database_errors.ErrDeleteRecordFailed
		return
	}

	return
}

func (db SQLiteDB) GetCourse(courseID int) (course db_model.Course, n_students int, n_tutors int, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get course if exists
	result := conn.First(&course, "id = ?", courseID)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// Get n_students for course
	var count int64
	result = conn.Model(&db_model.Booking{}).Where("course_id = ?", courseID).Count(&count)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}
	n_students = int(count)

	// Get n_tutors for course
	result = conn.Model(&db_model.TutorCourses{}).Where("course_id = ?", courseID).Count(&count)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
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
		err = database_errors.ErrNoRecordFound
		return
	}

	// Raw SQL for count group by query
	booking_map := make(map[int]int)
	var courseID int
	var count int
	rows, err := conn.Raw("SELECT course_id, COUNT(*) FROM bookings GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&courseID, &count)
		booking_map[courseID] = count
	}

	// Raw SQL for count group by query
	tutor_map := make(map[int]int)
	rows, err = conn.Raw("SELECT course_id, COUNT(*) FROM tutor_courses GROUP BY course_id").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&courseID, &count)
		tutor_map[courseID] = count
	}

	// format booking map and tutor courses map for results
	for i := 0; i < len(courses); i++ {
		n_students = append(n_students, booking_map[int(courses[i].ID)])
		n_tutors = append(n_tutors, tutor_map[int(courses[i].ID)])
	}

	return
}

func (db SQLiteDB) GetSchools() (schools []db_model.School, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all schools
	result := conn.Order("school_code ASC").Find(&schools)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetSchoolByInstitutionAndSchoolCode(institution string, schoolCode string) (school db_model.School, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get school if exists
	result := conn.First(&school, "institution = ? AND school_code = ?", institution, schoolCode)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}

func (db SQLiteDB) GetCoursesForSchool(school_id int) (schoolCoursesDetails db_model.SchoolCoursesDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var schoolCourses []db_model.SchoolCourses
	var school db_model.School

	// Get school if exists
	result := conn.First(&school, "id = ?", school_id)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// Get courses
	result = conn.Where("school_id = ?", school_id).Find(&schoolCourses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	courseIDs := make([]int, 0)

	// Put courses into list
	for i := 0; i < int(result.RowsAffected); i++ {
		courseIDs = append(courseIDs, int(schoolCourses[i].CourseID))
	}

	// Find courses corresponding to list
	var courses []db_model.Course
	result = conn.Where("id IN ? ", courseIDs).Find(&courses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	var courseCodes []string
	var CourseNames []string

	for i := 0; i < int(result.RowsAffected); i++ {
		courseCodes = append(courseCodes, courses[i].CourseCode)
		CourseNames = append(CourseNames, courses[i].CourseName)
	}

	// Format result
	schoolCoursesDetails.Institution = school.Institution
	schoolCoursesDetails.SchoolCode = school.SchoolCode
	schoolCoursesDetails.SchoolName = school.SchoolName
	schoolCoursesDetails.CourseCodes = courseCodes
	schoolCoursesDetails.CourseNames = CourseNames

	return
}

func (db SQLiteDB) GetBookingsByID(userID int) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", userID).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of availability_IDs and courseIDs
	availability_IDs := make([]int, 0)
	courseIDs := make([]int, 0)
	userIDs := make([]int, 0)

	for i := 0; i < int(result.RowsAffected); i++ {
		availability_IDs = append(availability_IDs, int(bookings_no_details[i].TutorAvailabilityID))
		courseIDs = append(courseIDs, int(bookings_no_details[i].CourseID))
		userIDs = append(userIDs, int(bookings_no_details[i].UserID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? ", availability_IDs).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get tutors corresponding to availabilities
	tutorIDs := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutorIDs = append(tutorIDs, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", courseIDs).Find(&courses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", userIDs).Find(&users)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding tutors
	var tutors []db_model.User
	result = conn.Where("id IN ? ", tutorIDs).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.User)
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
func (db SQLiteDB) GetBookingsByIDFrom(userID int, fromTime time.Time) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", userID).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of availability_IDs
	availability_IDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		availability_IDs = append(availability_IDs, int(bookings_no_details[i].TutorAvailabilityID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("ID IN ? AND available_from >= ?", availability_IDs, fromTime).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
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

	// create list of courseIDs and userIDs
	courseIDs := make([]int, 0)
	userIDs := make([]int, 0)
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		courseIDs = append(courseIDs, int(bookings_no_details_after_filter[i].CourseID))
		userIDs = append(userIDs, int(bookings_no_details_after_filter[i].UserID))
	}

	// get tutors corresponding to availabilities
	tutorIDs := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutorIDs = append(tutorIDs, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", courseIDs).Find(&courses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", userIDs).Find(&users)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding tutors
	var tutors []db_model.User
	result = conn.Where("id IN ? ", tutorIDs).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.User)
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
func (db SQLiteDB) GetBookingsByIDTo(userID int, toTime time.Time) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", userID).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of availability_IDs
	availability_IDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		availability_IDs = append(availability_IDs, int(bookings_no_details[i].TutorAvailabilityID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? AND available_to <= ?", availability_IDs, toTime).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
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

	// create list of courseIDs and userIDs
	courseIDs := make([]int, 0)
	userIDs := make([]int, 0)
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		courseIDs = append(courseIDs, int(bookings_no_details_after_filter[i].CourseID))
		userIDs = append(userIDs, int(bookings_no_details_after_filter[i].UserID))
	}

	// get tutors corresponding to availabilities
	tutorIDs := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutorIDs = append(tutorIDs, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("id IN ? ", courseIDs).Find(&courses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("id IN ? ", userIDs).Find(&users)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding tutors
	var tutors []db_model.User
	result = conn.Where("id IN ? ", tutorIDs).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.User)
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

func (db SQLiteDB) GetBookingsByIDFromTo(userID int, fromTime time.Time, toTime time.Time) (bookings []db_model.BookingDetails, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get bookings
	var bookings_no_details []db_model.Booking
	result := conn.Where("user_id = ?", userID).Find(&bookings_no_details)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// create list of availability_IDs
	availability_IDs := make([]int, 0)
	for i := 0; i < int(result.RowsAffected); i++ {
		availability_IDs = append(availability_IDs, int(bookings_no_details[i].TutorAvailabilityID))
	}

	// get corresponding availabilities
	var availabilities []db_model.Availability
	result = conn.Where("id IN ? AND available_from >= ? AND available_to <= ?", availability_IDs, fromTime, toTime).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
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

	// create list of courseIDs and userIDs
	courseIDs := make([]int, 0)
	userIDs := make([]int, 0)
	for i := 0; i < len(bookings_no_details_after_filter); i++ {
		courseIDs = append(courseIDs, int(bookings_no_details_after_filter[i].CourseID))
		userIDs = append(userIDs, int(bookings_no_details_after_filter[i].UserID))
	}

	// get tutors corresponding to availabilities
	tutorIDs := make([]int, 0)
	for i := 0; i < len(availabilities); i++ {
		tutorIDs = append(tutorIDs, int(availabilities[i].TutorID))
	}

	// get corresponding course codes
	var courses []db_model.Course
	result = conn.Where("ID IN ? ", courseIDs).Find(&courses)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding course codes
	var users []db_model.User
	result = conn.Where("ID IN ? ", userIDs).Find(&users)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// get corresponding tutors
	var tutors []db_model.User
	result = conn.Where("ID IN ? ", tutorIDs).Find(&tutors)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	// format into map for fast reference
	availability_map := make(map[int]db_model.Availability)
	courses_map := make(map[int]db_model.Course)
	users_map := make(map[int]db_model.User)
	tutors_map := make(map[int]db_model.User)
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
func (db SQLiteDB) CreateBooking(availabilityID int, userID int, courseID int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	var db_booking db_model.Booking
	var found bool
	found = true

	// Get booking if exists
	result := conn.Where("tutor_availability_ID = ? AND user_id = ? AND course_id = ?", availabilityID, userID, courseID).First(&db_booking)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	if found {
		err = database_errors.ErrRecordAlreadyExists
		return
	}

	// Update user fields
	db_booking.TutorAvailabilityID = availabilityID
	db_booking.UserID = userID
	db_booking.CourseID = courseID

	// create record
	result = conn.Create(&db_booking)
	if result.Error != nil {
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
func (db SQLiteDB) DeleteBookingByID(bookingID int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Delete(&db_model.Booking{}, bookingID)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrDeleteRecordFailed
		return
	}

	return
}

func (db SQLiteDB) GetAvailabilityByID(tutorID int) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ?", tutorID).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
func (db SQLiteDB) GetAvailabilityByIDFrom(tutorID int, fromTime time.Time) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ? AND available_from >= ?", tutorID, fromTime).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
func (db SQLiteDB) GetAvailabilityByIDTo(tutorID int, toTime time.Time) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ? AND available_to <= ?", tutorID, toTime).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
func (db SQLiteDB) GetAvailabilityByIDFromTo(tutorID int, fromTime time.Time, toTime time.Time) (availabilities []db_model.Availability, err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get all users
	result := conn.Where("tutor_id = ? AND available_from >= ? AND available_to = ?", tutorID, fromTime, toTime).Find(&availabilities)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
func (db SQLiteDB) CreateTutorAvailability(tutorID int, fromTime time.Time, toTime time.Time) (err error) {
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
	result := conn.Where("tutor_id = ? AND available_from = ? AND available_to = ?", tutorID, fromTime, toTime).First(&db_availability)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// row not found
		found = false
	}

	if found {
		err = database_errors.ErrRecordAlreadyExists
		return
	}

	// Update user fields
	db_availability.TutorID = uint(tutorID)
	db_availability.AvailableFrom = fromTime
	db_availability.AvailableTo = toTime

	// create new record
	result = conn.Create(&db_availability)
	if result.Error != nil {
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
func (db SQLiteDB) DeleteTutorAvailabilityByID(availabilityID int) (err error) {
	conn, err := gorm.Open(sqlite.Open(db.DatabaseFilename), &gorm.Config{})
	if err != nil {
		// err opening database
		log.Println(err)
		return
	}

	// Get user if exists
	result := conn.Delete(&db_model.Availability{}, availabilityID)
	if result.Error != nil {
		// row not found
		err = database_errors.ErrNoRecordFound
		return
	}

	return
}
