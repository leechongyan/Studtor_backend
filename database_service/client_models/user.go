package client_models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	httpError "github.com/leechongyan/Studtor_backend/constants/errors/http_errors"
	databaseModel "github.com/leechongyan/Studtor_backend/database_service/database_models"
)

var validate = validator.New()

type jsonUser struct {
	Id             *int    `json:"id"`
	FirstName      *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName       *string `json:"last_name" validate:"required,min=2,max=100"`
	Password       *string `json:"password" validate:"required,min=6"`
	Email          *string `json:"email" validate:"email,required"`
	Token          *string `json:"token"`
	UserType       *string `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken   *string `json:"refresh_token"`
	VKey           *string
	ProfilePicture *string
	Verified       bool
	UserCreatedAt  time.Time `json:"created_at"`
	UserUpdatedAt  time.Time `json:"updated_at"`
}

type User struct {
	id             *int
	firstName      *string
	lastName       *string
	password       *string
	email          *string
	token          *string
	userType       *string
	refreshToken   *string
	vKey           *string
	profilePicture *string
	verified       bool
	userCreatedAt  time.Time
	userUpdatedAt  time.Time
}

func (user *User) ID() *int {
	return user.id
}

func (user *User) Email() *string {
	return user.email
}

func (user *User) Password() *string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = &password
}

func (user *User) VKey() *string {
	return user.vKey
}

func (user *User) SetVKey(vkey string) {
	if vkey != "" {
		user.vKey = &vkey
	} else {
		user.vKey = nil
	}
}

func (user *User) Verified() bool {
	return user.verified
}

func (user *User) UserCreatedAt() time.Time {
	return user.userCreatedAt
}

func (user *User) UserUpdatedAt() time.Time {
	return user.userUpdatedAt
}

func (user *User) SetVerified(verified bool) {
	user.verified = verified
}

func (user *User) FirstName() *string {
	return user.firstName
}

func (user *User) LastName() *string {
	return user.lastName
}

func (user *User) UserType() *string {
	return user.userType
}

func (user *User) RefreshToken() *string {
	return user.refreshToken
}

// need to set to empty
func (user *User) SetRefreshToken(refreshToken string) {
	if refreshToken != "" {
		user.refreshToken = &refreshToken
	} else {
		user.refreshToken = nil
	}
}

func (user *User) Token() *string {
	return user.token
}

func (user *User) SetToken(token string) {
	if token != "" {
		user.token = &token
	} else {
		user.token = nil
	}
}

func (user *User) ProfilePicture() *string {
	return user.profilePicture
}

func (user *User) SetProfilePicture(url string) {
	if url != "" {
		user.profilePicture = &url
	} else {
		user.profilePicture = nil
	}
}

func (user *User) SetUserCreatedAt(createdAt time.Time) {
	user.userCreatedAt = createdAt
}

func (user *User) SetUserUpdatedAt(updatedAt time.Time) {
	user.userUpdatedAt = updatedAt
}

func convertFromjsonUserToUser(jsUser jsonUser) (user User) {
	user = User{}
	user.id = jsUser.Id
	user.firstName = jsUser.FirstName
	user.lastName = jsUser.LastName
	user.password = jsUser.Password
	user.email = jsUser.Email
	user.token = jsUser.Token
	user.userType = jsUser.UserType
	user.refreshToken = jsUser.RefreshToken
	user.vKey = jsUser.VKey
	user.profilePicture = jsUser.ProfilePicture
	user.verified = jsUser.Verified
	user.userCreatedAt = jsUser.UserCreatedAt
	user.userUpdatedAt = jsUser.UserUpdatedAt
	return user
}

func UnmarshalJson(c *gin.Context) (user User, err error) {
	jsUser := jsonUser{}
	err = c.BindJSON(&jsUser)
	if err != nil {
		return User{}, httpError.ErrJsonParsingFailure
	}
	err = validate.Struct(&jsUser)
	if err != nil {
		return User{}, httpError.ErrJsonValidationError
	}
	return convertFromjsonUserToUser(jsUser), err
}

type UserProfile struct {
	id             int
	firstName      string
	lastName       string
	email          string
	profilePicture string
}

// conversion method from Userprofile
func ConvertFromDatabaseUserToUserProfile(databaseuser databaseModel.User) (profile UserProfile) {
	profile = UserProfile{}
	profile.id = int(databaseuser.ID)
	profile.firstName = databaseuser.FirstName
	profile.lastName = databaseuser.LastName
	profile.email = databaseuser.Email
	if databaseuser.ProfilePicture.Valid {
		profile.profilePicture = databaseuser.ProfilePicture.String
	} else {
		profile.profilePicture = ""
	}
	return profile
}

// this is only used for initial sign up for loading the user credential into database user object
func ConvertFromAuthUserToDatabaseUser(user User) (databaseuser databaseModel.User) {
	databaseuser = databaseModel.User{}
	databaseuser.UserCreatedAt = user.userCreatedAt
	databaseuser.UserUpdatedAt = user.userUpdatedAt
	if user.ID() != nil {
		databaseuser.ID = uint(*user.ID())
	} else {
		databaseuser.ID = 0
	}
	databaseuser.FirstName = *user.firstName
	databaseuser.LastName = *user.lastName
	databaseuser.Password = *user.password
	databaseuser.Email = *user.email
	if user.profilePicture != nil {
		databaseuser.ProfilePicture.String = *user.profilePicture
		databaseuser.ProfilePicture.Valid = true
	} else {
		databaseuser.ProfilePicture.Valid = false
	}
	if user.token != nil {
		databaseuser.Token.String = *user.token
		databaseuser.Token.Valid = true
	} else {
		databaseuser.Token.Valid = false
	}
	databaseuser.UserType = *user.userType
	if user.refreshToken != nil {
		databaseuser.RefreshToken.String = *user.refreshToken
		databaseuser.RefreshToken.Valid = true
	} else {
		databaseuser.RefreshToken.Valid = false
	}
	if user.vKey != nil {
		databaseuser.VKey.String = *user.vKey
		databaseuser.VKey.Valid = true
	} else {
		databaseuser.VKey.Valid = false
	}
	databaseuser.Verified = user.verified
	return databaseuser
}

func ConvertFromToDatabaseUserToAuthUser(databaseuser databaseModel.User) (usr User) {
	usr = User{}
	usr.userCreatedAt = databaseuser.UserCreatedAt
	usr.userUpdatedAt = databaseuser.UserUpdatedAt
	databaseuserID := int(databaseuser.ID)
	usr.id = &databaseuserID
	usr.firstName = &databaseuser.FirstName
	usr.lastName = &databaseuser.LastName
	usr.password = &databaseuser.Password
	usr.email = &databaseuser.Email
	if databaseuser.ProfilePicture.Valid {
		usr.profilePicture = &databaseuser.ProfilePicture.String
	}
	if databaseuser.Token.Valid {
		usr.token = &databaseuser.Token.String
	}
	usr.userType = &databaseuser.UserType
	if databaseuser.VKey.Valid {
		usr.vKey = &databaseuser.VKey.String
	}
	usr.verified = databaseuser.Verified
	return usr
}
