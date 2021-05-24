package storage_service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	nurl "net/url"

	"cloud.google.com/go/storage"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

type googlestorage struct {
	storageClient *storage.Client
	bucketName    string
	ctx           context.Context
}

func InitGoogleStorage() (gs googlestorage, err error) {
	gs = googlestorage{}
	gs.ctx = context.Background()
	gs.bucketName = viper.GetString("google_bucket_name")
	gs.storageClient, err = storage.NewClient(gs.ctx, option.WithCredentialsFile("../cred.json"))
	if err != nil {
		return gs, errors.New("Cannot add Storage Client")
	}
	return
}

func (gs googlestorage) SaveUserProfilePicture(user_id string, file multipart.File, fileheader multipart.FileHeader) (url string, err error) {
	return gs.saveImage(fileheader.Filename, "users/"+user_id, file)
}

func (gs googlestorage) SaveTutorNotesForACourse(tutor_id string, course_code string, file multipart.File, fileheader multipart.FileHeader) (url string, err error) {
	return gs.saveImage(fileheader.Filename, "notes/"+tutor_id+"/"+course_code, file)
}

func (gs googlestorage) SaveCourseProfilePicture(course_code string, file multipart.File, fileheader multipart.FileHeader) (url string, err error) {
	return gs.saveImage(fileheader.Filename, "courses/"+course_code, file)
}

func (gs googlestorage) saveImage(file_name string, sub_directory string, file multipart.File) (url string, err error) {
	sw := gs.storageClient.Bucket(gs.bucketName).Object(sub_directory + "/" + file_name).NewWriter(gs.ctx)
	fmt.Print("Step 1")
	if _, err = io.Copy(sw, file); err != nil {
		fmt.Print(err.Error())
		return "", errors.New("Cannot copy file over")
	}
	fmt.Print("Step 2")
	if err = sw.Close(); err != nil {
		fmt.Print(err.Error())
		return "", errors.New("Cannot close object writer")
	}
	fmt.Print("Step 3")
	u, err := nurl.Parse("/" + gs.bucketName + "/" + sw.Attrs().Name)
	if err != nil {
		fmt.Print(err.Error())
		return "", errors.New("Cannot parse URL")
	}
	fmt.Print("Step 4")
	return u.EscapedPath(), nil
}
