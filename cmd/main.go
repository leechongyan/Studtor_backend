package main

import (
	"github.com/gin-gonic/gin"
	authhandler "github.com/leechongyan/Studtor_backend/authentication_service/controllers"
	authMiddleWare "github.com/leechongyan/Studtor_backend/authentication_service/middleware"
	initialization_helper "github.com/leechongyan/Studtor_backend/helpers/initialization_helpers"
	tuthandler "github.com/leechongyan/Studtor_backend/tuition_service/controllers"
	"github.com/spf13/viper"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	err := initialization_helper.Initialize()
	if err != nil {
		return
	}

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

	authorized.POST("/user/logout", authhandler.Logout())

	// get current user
	home.GET("/user", authhandler.GetCurrentUser())
	// get other user by their user id
	home.GET("/users/:user_id", authhandler.GetUser())

	// get a list of courses
	home.GET("/courses", tuthandler.GetCourses())

	// get a course
	home.GET("/courses/:course_id", tuthandler.GetSingleCourse())

	// get the tutors for a course
	home.GET("/courses/:course_id/tutors/", tuthandler.GetTutorsForCourse())

	// get all schools (this is for filtering in modal)
	home.GET("/schools", tuthandler.GetSchools())

	// get a single tutor
	home.GET("/tutors/:tutor_id", tuthandler.GetSingleTutor())

	// get all the tutors
	home.GET("/tutors", tuthandler.GetAllTutors())

	// get a list of courses taught by a tutor
	home.GET("/tutors/:tutor_id/courses", tuthandler.GetCoursesOfTutor())

	// register for a course for a tutor
	home.POST("/tutors/:tutor_id/courses/:course_id", tuthandler.RegisterCourse())

	// deregister a course for a tutor
	home.DELETE("/tutors/:tutor_id/courses/:course_id", tuthandler.DeregisterCourse())

	// put available time for tutor
	home.POST("/availability", tuthandler.PutAvailableTimeTutor())

	// delete available time for tutor
	home.DELETE("/availability/:availability_id", tuthandler.DeleteAvailableTimeTutor())

	// get availability for a tutor
	home.GET("/availability/tutors/:tutor_id", tuthandler.GetAvailableTimeTutor())

	// book an available time for a tutor
	home.POST("availability/:availability_id/courses/:course_id/tutors/:tutor_id", tuthandler.BookTimeTutor())

	// unbook a booking for a tutor (can be done by a student)
	home.DELETE("/booking/:booking_id", tuthandler.UnbookTimeTutor())

	// get all the booked time for a user or a tutor
	home.GET("/users/:user_id/booking", tuthandler.GetAllBookedTime())
	// end of version v1

	router.Run(viper.GetString("port"))
}
