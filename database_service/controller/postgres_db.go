package controller

import (
	"sort"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/clause"

	csvModel "github.com/leechongyan/Studtor_backend/database_service/csv_models"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
	databaseError "github.com/leechongyan/Studtor_backend/database_service/errors"
	"gorm.io/gorm"
)

type postgresdb struct {
	db *gorm.DB
}

func InitPostGres(config string) (pgdb postgresdb, err error) {
	pgdb.db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return pgdb, err
	}
	err = pgdb.db.AutoMigrate(&databaseModel.User{})
	if err != nil {
		return pgdb, err
	}
	err = pgdb.db.AutoMigrate(&databaseModel.Faculty{})
	if err != nil {
		return pgdb, err
	}
	err = pgdb.db.AutoMigrate(&databaseModel.Availability{})
	if err != nil {
		return pgdb, err
	}
	err = pgdb.db.AutoMigrate(&databaseModel.Booking{})
	if err != nil {
		return pgdb, err
	}
	err = pgdb.db.AutoMigrate(&databaseModel.School{})
	if err != nil {
		return pgdb, err
	}
	err = pgdb.db.AutoMigrate(&databaseModel.Course{})
	if err != nil {
		return pgdb, err
	}
	return pgdb, createCourse(pgdb.db)
}

func (pgdb postgresdb) GetUserByID(userID int) (user databaseModel.User, err error) {
	result := pgdb.db.Preload("Availabilities").Preload("Bookings").Preload("Courses").First(&user, userID)
	return user, result.Error
}

// GetUserByEmail retrieves a user model object by the user's email from the database.
func (pgdb postgresdb) GetUserByEmail(email string) (user databaseModel.User, err error) {
	result := pgdb.db.Where(&databaseModel.User{Email: email}).First(&user)
	return user, result.Error
}

// CreateUser saves an user object into the database.
func (pgdb postgresdb) CreateUser(user databaseModel.User) (id int, err error) {
	result := pgdb.db.Create(&user)
	return int(user.ID), result.Error
}

// UpdateUser updates an user object into the database.
func (pgdb postgresdb) UpdateUser(user databaseModel.User) (id int, err error) {
	result := pgdb.db.Save(&user)
	return int(user.ID), result.Error
}

// DeleteUserByID deletes an user object by the user's ID from the database.
func (pgdb postgresdb) DeleteUserByID(userID int) (err error) {
	result := pgdb.db.Delete(&databaseModel.User{}, userID)
	return result.Error
}

// DeleteUserByEmail deletes an user object by the user's email from the database.
func (pgdb postgresdb) DeleteUserByEmail(email string) (err error) {
	result := pgdb.db.Delete(&databaseModel.User{}, "Email = ?", email)
	return result.Error
}

func (pgdb postgresdb) GetCoursesForTutor(tutorID int) (courses []databaseModel.Course, nStudents []int, nTutors []int, err error) {
	user := databaseModel.User{}
	result := pgdb.db.Preload("Courses").Preload("Courses.Tutors").First(&user, tutorID)
	if result.Error != nil {
		return courses, nStudents, nTutors, result.Error
	}
	// now query bookings
	courses = user.Courses
	nStudents = make([]int, len(user.Courses))
	nTutors = make([]int, len(user.Courses))
	for i, v := range courses {
		var count int64
		result = pgdb.db.Model(&databaseModel.Booking{}).Where(&databaseModel.Booking{CourseID: v.ID}).Distinct("StudentID").Count(&count)
		if result.Error != nil {
			return courses, nStudents, nTutors, result.Error
		}
		nStudents[i] = int(count)
		nTutors[i] = len(v.Tutors)
	}
	return
}

// GetTutorsForCourse retrieves a list of all tutors for a particular course from the database.
func (pgdb postgresdb) GetTutorsForCourse(courseID int) (tutors []databaseModel.User, err error) {
	course := databaseModel.Course{}
	result := pgdb.db.Preload("Tutors").First(&course, courseID)
	if result.Error != nil {
		return tutors, result.Error
	}
	return course.Tutors, err
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetTutorsForCourseFromIDOfSize retrieves a list of tutors for a particular course from the database,
// starting from tut_ID to tut_ID + size
func (pgdb postgresdb) GetTutorsForCourseFromIDOfSize(courseID int, tutorID int, size int) (tutors []databaseModel.User, err error) {
	course := databaseModel.Course{}
	result := pgdb.db.Preload("Tutors").First(&course, courseID)
	if result.Error != nil {
		return tutors, result.Error
	}
	sort.Slice(course.Tutors, func(i, j int) bool {
		return tutors[i].ID < tutors[j].ID
	})
	idx := sort.Search(len(course.Tutors), func(i int) bool {
		return int(course.Tutors[i].ID) >= tutorID
	})

	return course.Tutors[idx:min(len(course.Tutors), idx+size)], err
}

// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
// starting from tut_ID to the end
func (pgdb postgresdb) GetTutorsForCourseFromID(courseID int, tutorID int) (tutors []databaseModel.User, err error) {
	course := databaseModel.Course{}
	result := pgdb.db.Preload("Tutors").First(&course, courseID)
	if result.Error != nil {
		return tutors, result.Error
	}
	sort.Slice(course.Tutors, func(i, j int) bool {
		return tutors[i].ID < tutors[j].ID
	})
	idx := sort.Search(len(course.Tutors), func(i int) bool {
		return int(course.Tutors[i].ID) >= tutorID
	})

	return course.Tutors[idx:len(course.Tutors)], err
}

// GetTutorsForCourseFromID retrieves a list of tutors for a particular course from the database,
// starting from 0 to size
func (pgdb postgresdb) GetTutorsForCourseOfSize(courseID int, size int) (tutors []databaseModel.User, err error) {
	course := databaseModel.Course{}
	result := pgdb.db.Preload("Tutors").First(&course, courseID)
	if result.Error != nil {
		return tutors, result.Error
	}
	sort.Slice(course.Tutors, func(i, j int) bool {
		return tutors[i].ID < tutors[j].ID
	})

	return course.Tutors[:min(len(course.Tutors), size)], err
}

// CreateTutorCourse saves a tutor_course model object into the database.
// This function is called when a tutor registers interest to teach a course.
func (pgdb postgresdb) CreateTutorCourse(tutorID int, courseID int) (err error) {
	// retrieve Course from CourseID
	// first retrieve tutor
	course := databaseModel.Course{}
	result := pgdb.db.First(&course, courseID)
	if result.Error != nil {
		return result.Error
	}
	user := databaseModel.User{}
	result = pgdb.db.First(&user, tutorID)
	if result.Error != nil {
		return result.Error
	}
	return pgdb.db.Model(&course).Association("Tutors").Append(&user)
}

// DeleteTutorCourse deletes an tutor course object from the database.
func (pgdb postgresdb) DeleteTutorCourse(tutorID int, courseID int) (err error) {
	course := databaseModel.Course{}
	result := pgdb.db.First(&course, courseID)
	if result.Error != nil {
		return result.Error
	}
	user := databaseModel.User{}
	result = pgdb.db.First(&user, tutorID)
	if result.Error != nil {
		return result.Error
	}
	return pgdb.db.Model(&course).Association("Tutors").Delete(user)
}

func (pgdb postgresdb) GetCourse(courseID int) (course databaseModel.Course, nStudents int, nTutors int, err error) {
	result := pgdb.db.Preload("Tutors").First(&course, courseID)
	if result.Error != nil {
		return course, nStudents, nTutors, result.Error
	}
	var count int64
	result = pgdb.db.Model(&databaseModel.Booking{}).Where(&databaseModel.Booking{CourseID: uint(courseID)}).Distinct("StudentID").Count(&count)
	if result.Error != nil {
		return course, nStudents, nTutors, result.Error
	}
	return course, int(count), len(course.Tutors), err
}

func (pgdb postgresdb) CreateCourse(course databaseModel.Course) (id int, err error) {
	result := pgdb.db.Create(&course)
	return id, result.Error
}

// GetCourses retrieves a list of all courses, along with the number of students
// enrolled in the course and the number of tutors for the course, from the database.
// Sorted by course code.
func (pgdb postgresdb) GetCourses() (courses []databaseModel.Course, nStudents []int, nTutors []int, err error) {
	// get all the Courses
	result := pgdb.db.Preload("Tutors").Find(&courses)
	nStudents = make([]int, len(courses))
	nTutors = make([]int, len(courses))
	if result.Error != nil {
		return courses, nStudents, nTutors, result.Error
	}
	for i, v := range courses {
		var count int64
		result = pgdb.db.Model(&databaseModel.Booking{}).Where(&databaseModel.Booking{CourseID: uint(v.ID)}).Distinct("StudentID").Count(&count)
		if result.Error != nil {
			return courses, nStudents, nTutors, result.Error
		}
		nStudents[i] = int(count)
		nTutors[i] = len(v.Tutors)
	}
	return
}

// GetSingleBooking gets a single booking details for a booking ID
func (pgdb postgresdb) GetSingleBooking(bookingID int) (booking databaseModel.Booking, err error) {
	book := databaseModel.Booking{}
	result := pgdb.db.Preload(clause.Associations).First(&book, bookingID)
	return booking, result.Error
}

// GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID, with no time constraints
func (pgdb postgresdb) GetBookingsByID(userID int) (bookings []databaseModel.Booking, err error) {
	user := databaseModel.User{}
	result := pgdb.db.Preload(clause.Associations).Preload("Bookings."+clause.Associations).Preload("Bookings.Availability."+clause.Associations).First(&user, userID)
	return bookings, result.Error
}

// // GetBookingsByIDFrom retrieves a list of all bookings by a user, as indicated by userID, starting from time fromTime
// GetBookingsByIDFrom(userID int, fromTime time.Time) (bookings []databaseModel.BookingDetails, err error)
// // GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID, ending with time toTime
// GetBookingsByIDTo(userID int, toTime time.Time) (bookings []databaseModel.BookingDetails, err error)
// // GetBookingsByID retrieves a list of all bookings by a user, as indicated by userID,
// // starting from time fromTime and ending with time toTime
// GetBookingsByIDFromTo(userID int, fromTime time.Time, toTime time.Time) (bookings []databaseModel.BookingDetails, err error)

// GetBookingsByIDFromDateForSize retrieves a list of all bookings for a user from a date up to x days
func (pgdb postgresdb) GetBookingsByIDFromDateForSize(userID int, date time.Time, days int) (bookings []databaseModel.Booking, err error) {
	return
}

// CreateBooking saves a booking model object into the database
func (pgdb postgresdb) CreateBooking(availabilityID int, userID int, courseID int) (id int, err error) {
	// get the availability
	avail := databaseModel.Availability{}
	result := pgdb.db.First(&avail, availabilityID)
	if result.Error != nil {
		return id, result.Error
	}
	// check whether this availability is occupied
	if avail.Occupied {
		return id, databaseError.ErrInvalidAvailability
	}
	avail.Occupied = true
	result = pgdb.db.Save(&avail)
	if result.Error != nil {
		return id, result.Error
	}
	// course := Course{}
	// result = db.First(&course, courseID)
	// if result.Error != nil {
	// 	return result.Error
	// }
	// student := User{}
	// result = db.First(&student, userID)
	// if result.Error != nil {
	// 	return result.Error
	// }
	booking := databaseModel.Booking{StudentID: uint(userID), AvailabilityID: uint(availabilityID), CourseID: uint(courseID)}

	// create a booking
	result = pgdb.db.Create(&booking)
	return int(booking.ID), result.Error
}

// DeleteBooking deletes a booking model object into the database
func (pgdb postgresdb) DeleteBookingByID(bookingID int) (err error) {
	booking := databaseModel.Booking{}
	result := pgdb.db.Preload("Availability").First(&booking, bookingID)
	if result.Error != nil {
		return result.Error
	}
	result = pgdb.db.Delete(&databaseModel.Booking{}, bookingID)
	if result.Error != nil {
		return result.Error
	}
	avail := booking.Availability
	avail.Occupied = false
	result = pgdb.db.Save(&avail)
	return result.Error
}

/*
	TutorAvailability model
*/

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, with no time constraints
func (pgdb postgresdb) GetAvailabilityByID(tutorID int) (availabilities []databaseModel.Availability, err error) {
	user := databaseModel.User{}
	result := pgdb.db.Preload("Availabilities").First(&user, tutorID)
	return user.Availabilities, result.Error
}

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, starting from time fromTime
// GetAvailabilityByIDFrom(tutorID int, fromTime time.Time) (availabilities []databaseModel.Availability, err error)
// GetAvailabilityByID retrieves a list of all available timeslots for a tutor, ending with time toTime
// GetAvailabilityByIDTo(tutorID int, toTime time.Time) (availabilities []databaseModel.Availability, err error)

// GetSingleAvailability gets an availability information based on the availability ID
func (pgdb postgresdb) GetSingleAvailability(availabilityID int) (availability databaseModel.Availability, err error) {
	avail := databaseModel.Availability{}
	result := pgdb.db.Preload("Tutor").First(&avail, availabilityID)
	return avail, result.Error
}

// GetAvailabilityByID retrieves a list of all available timeslots for a tutor,
// starting from time fromTime and ending with time toTime
// GetAvailabilityByIDFromTo(tutorID int, fromTime time.Time, toTime time.Time) (availabilities []databaseModel.Availability, err error)

// GetAvailabiltyByIDFromDateForSize retrieves a list of all available timeslots for a tutor from a date up to x days
func (pgdb postgresdb) GetAvailabilityByIDFromDateForSize(tutorId int, date time.Time, days int) (availabilities []databaseModel.Availability, err error) {
	return
}

// CreateTutorAvailability saves a tutor availability model object into the database
func (pgdb postgresdb) CreateTutorAvailability(tutorID int, date time.Time, timeID int) (id int, err error) {
	avail := databaseModel.Availability{}
	avail.Occupied = false
	avail.TutorID = uint(tutorID)
	avail.Date = date
	avail.TimeSlot = timeID
	result := pgdb.db.Create(&avail)
	return int(avail.ID), result.Error
}

// DeleteTutorAvailability deletes a tutor availability model object by ID from the database
func (pgdb postgresdb) DeleteTutorAvailabilityByID(availabilityID int) (err error) {
	result := pgdb.db.Delete(&databaseModel.Availability{}, availabilityID)
	return result.Error
}

func (pgdb postgresdb) GetSchoolsFacultiesCourses() (schools []databaseModel.School, err error) {
	result := pgdb.db.Preload(clause.Associations).Preload("Faculties." + clause.Associations).Find(&schools)
	return schools, result.Error
}

func createCourse(db *gorm.DB) (err error) {
	schools, err := csvModel.ImportSchool()
	if err != nil {
		return err
	}
	faculties, err := csvModel.ImportFaculty()
	if err != nil {
		return err
	}
	courses, err := csvModel.ImportCourse()
	if err != nil {
		return err
	}
	// add schools first
	result := db.Create(schools)
	if result.Error != nil {
		return result.Error
	}
	// add faculties next
	result = db.Create(faculties)
	if result.Error != nil {
		return result.Error
	}
	// add courses last
	result = db.Create(courses)
	return result.Error
}
