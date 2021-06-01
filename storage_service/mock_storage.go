package storage_service

import (
	"mime/multipart"
)

type mockstorage struct {
}

func InitMock() (ms mockstorage) {
	ms = mockstorage{}
	return ms
}

func (storage mockstorage) SaveUserProfilePicture(user_id string, file multipart.File) (url string, err error) {
	return "www.sample.com/" + user_id, nil
}

func (storage mockstorage) SaveTutorNotesForACourse(tutor_id string, course_code string, file multipart.File, fileheader multipart.FileHeader) (url string, err error) {
	return "www.sample.com/" + fileheader.Filename, nil
}
