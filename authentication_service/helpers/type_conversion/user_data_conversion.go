package type_conversion

import (
	authModel "github.com/leechongyan/Studtor_backend/authentication_service/models"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/models"
)

// this is used for returning for getting user profile information from database
func ConvertFromDatabaseUserToUserProfile(databaseuser databaseModel.User) (profile authModel.Userprofile) {
	profile = authModel.Userprofile{}
	profile.ID = int(databaseuser.ID)
	profile.FirstName = databaseuser.FirstName
	profile.LastName = databaseuser.LastName
	profile.Email = databaseuser.Email
	if databaseuser.ProfilePicture.Valid {
		profile.ProfilePicture = databaseuser.ProfilePicture.String
	} else {
		profile.ProfilePicture = ""
	}
	return profile
}

// this is only used for initial sign up for loading the user credential into database user object
func ConvertFromAuthUserToDatabaseUser(authuser authModel.User) (databaseuser databaseModel.User) {
	databaseuser = databaseModel.User{}
	databaseuser.UserCreatedAt = authuser.UserCreatedAt
	databaseuser.UserUpdatedAt = authuser.UserUpdatedAt
	if authuser.Id != nil {
		databaseuser.ID = uint(*authuser.Id)
	} else {
		databaseuser.ID = 0
	}
	databaseuser.FirstName = *authuser.FirstName
	databaseuser.LastName = *authuser.LastName
	databaseuser.Password = *authuser.Password
	databaseuser.Email = *authuser.Email
	if authuser.ProfilePicture != nil {
		databaseuser.ProfilePicture.String = *authuser.ProfilePicture
		databaseuser.ProfilePicture.Valid = true
	} else {
		databaseuser.ProfilePicture.Valid = false
	}
	if authuser.Token != nil {
		databaseuser.Token.String = *authuser.Token
		databaseuser.Token.Valid = true
	} else {
		databaseuser.Token.Valid = false
	}
	databaseuser.UserType = *authuser.UserType
	if authuser.RefreshToken != nil {
		databaseuser.RefreshToken.String = *authuser.RefreshToken
		databaseuser.RefreshToken.Valid = true
	} else {
		databaseuser.RefreshToken.Valid = false
	}
	if authuser.VKey != nil {
		databaseuser.VKey.String = *authuser.VKey
		databaseuser.VKey.Valid = true
	} else {
		databaseuser.VKey.Valid = false
	}
	return databaseuser
}
