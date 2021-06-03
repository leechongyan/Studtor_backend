package main

import (
	"log"

	"github.com/gin-gonic/gin"
	authhandler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"
	authMiddleWare "github.com/leechongyan/Studtor_backend/authentication_service/middleware"
	initialization_helper "github.com/leechongyan/Studtor_backend/helpers/initialization_helpers"
	tuthandler "github.com/leechongyan/Studtor_backend/tuition_service/controllers"
)

func main() {
	err := initialization_helper.Initialize()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	router := gin.Default()
	// current version is v1
	v1 := router.Group("/v1")

	// does not require token
	authorized := v1.Group("/auth")
	authorized.POST("/signup", authhandler.SignUp())
	authorized.POST("/verify", authhandler.Verify())
	authorized.POST("/login", authhandler.Login())
	authorized.POST("/refresh", authhandler.RefreshToken())

	// require token
	home := v1.Group("/")
	home.Use(authMiddleWare.Authentication())

	home.POST("/user/logout", authhandler.Logout())

	// get current user
	home.GET("/user", authhandler.GetCurrentUser())
	// get other user by their user id
	home.GET("/users/:user_id", authhandler.GetUser())

	// student usage
	// get the filtering field for student
	home.GET("/schools", tuthandler.GetSchools()) // done

	// get a list of courses
	home.GET("/courses", tuthandler.GetCourses()) // done

	// get a course
	home.GET("/courses/:course_id", tuthandler.GetSingleCourse()) // done

	// get the tutors for a course
	home.GET("/courses/:course_id/tutors", tuthandler.GetTutorsForCourse()) // done

	// get a tutor for a course
	home.GET("/courses/:course_id/tutors/:user_id", authhandler.GetUser()) // done

	// tutors usage
	// get a list of courses taught by a tutor
	home.GET("/tutors/:tutor_id/courses", tuthandler.GetCoursesOfTutor()) // done

	// register for a course for a tutor
	home.POST("/tutors/:tutor_id/courses/:course_id", tuthandler.RegisterCourse()) // done

	// deregister a course for a tutor
	home.DELETE("/tutors/:tutor_id/courses/:course_id", tuthandler.DeregisterCourse()) // done

	// making appointment
	// put available time for tutor
	home.POST("/tutors/:tutor_id/availability", tuthandler.PutAvailableTimeTutor()) // done

	// delete available time for tutor
	home.DELETE("/tutors/:tutor_id/availability/:availability_id", tuthandler.DeleteAvailableTimeTutor()) // done

	// get availability for a tutor
	home.GET("/tutors/:tutor_id/availability", tuthandler.GetAvailableTimeTutor()) // done

	// book an available time for a tutor
	home.POST("/courses/:course_id/tutors/:tutor_id/availability/:availability_id", tuthandler.BookTimeTutor()) // done

	// unbook a booking for a tutor (can be done by a student)
	home.DELETE("/users/:user_id/bookings/:booking_id", tuthandler.UnbookTimeTutor()) // done

	// get all the booked time for a user or a tutor
	home.GET("/users/:user_id/bookings", tuthandler.GetAllBookedTime()) // done
	// end of version v1

	router.Run(":3000")
}
