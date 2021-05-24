package storage_service

import (
	"mime/multipart"
	"strconv"

	"github.com/spf13/viper"
)

// store it by user/ course profile picture/ notes also
// storage of user profile picture
// :user:profile
// storage of user notes for what course
// :tutor:notes
// storage of course photos
// :course_id
var CurrentStorageConnector StorageConnector

type StorageConnector interface {
	SaveUserProfilePicture(user_id string, file multipart.File, fileheader multipart.FileHeader) (url string, err error)
	SaveTutorNotesForACourse(tutor_id string, course_code string, file multipart.File, fileheader multipart.FileHeader) (url string, err error)
	SaveCourseProfilePicture(course_code string, file multipart.File, fileheader multipart.FileHeader) (url string, err error)
}

func InitStorage() (err error) {
	isMock, _ := strconv.ParseBool(viper.GetString("mock_storage"))
	if isMock {
		CurrentStorageConnector = InitMock()
		return
	}
	// future add cred file
	CurrentStorageConnector, err = InitGoogleStorage()
	return
}
