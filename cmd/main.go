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
	authorized.POST("/logout", authhandler.Logout())

	// require token
	home := v1.Group("/")
	home.Use(authMiddleWare.Authentication())

	// ping to validate authorized status
	home.GET("/", authhandler.GetMain())

	// get current user
	home.GET("/user", authhandler.GetCurrentUser())
	// get other user by their user id
	home.GET("/user/:user", authhandler.GetUser())

	// get a list of courses
	home.GET("/courses", tuthandler.GetCourses())

	// get a course
	home.GET("/courses/:course", tuthandler.GetSingleCourse())

	// get the tutors for a course
	home.GET("/courses/:course/tutors/", tuthandler.GetCoursesTutors())

	// get all the tutors or get a single tutor
	home.GET("/tutors/*tutor_id", tuthandler.GetTutors())

	// for tutor usage
	home.POST("/tutors/putavailabletime", tuthandler.PutAvailableTimeTutor())
	home.POST("/tutors/deleteavailabletime", tuthandler.DeleteAvailableTimeTutor())

	// for student usage
	home.GET("/students/availabletime/:tutor", tuthandler.GetAvailableTimeTutor())
	home.POST("/students/book", tuthandler.BookTimeTutor())
	home.POST("/students/unbook", tuthandler.UnbookTimeTutor())
	home.GET("/users/bookedtime/*user", tuthandler.GetAllBookedTime())
	// end of version v1

	router.Run(viper.GetString("port"))
}
