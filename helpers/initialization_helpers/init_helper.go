package initialization_helpers

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	tokenHelper "github.com/leechongyan/Studtor_backend/authentication_service/helpers/account"
	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	mailService "github.com/leechongyan/Studtor_backend/mail_service"
	storageService "github.com/leechongyan/Studtor_backend/storage_service"
	"github.com/spf13/viper"
)

type config struct {
	JwtKey                *string `mapstructure:"jwtKey" validate:"required"`
	AccessExpirationTime  *int    `mapstructure:"accessExpirationTime"`
	RefreshExpirationTime *int    `mapstructure:"refreshExpirationTime"`
	ServerEmail           *string `mapstructure:"serverEmail" validate:"required,email"`
	ServerEmailPW         *string `mapstructure:"serverEmailPW" validate:"required"`
	GoogleBucketName      *string `mapstructure:"google_bucket_name" validate:"required"`
	IsMockDB              *bool   `mapstructure:"mock_database"`
	IsMockStorage         *bool   `mapstructure:"mock_storage"`
	DatabaseConfig        string  `mapstructure:"database_config"`
}

func getConfiguration() (conf config, err error) {
	validate := validator.New()

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return conf, systemError.ErrInitFailure
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return conf, systemError.ErrInitFailure
	}

	if err := validate.Struct(&conf); err != nil {
		return conf, err
	}

	if conf.AccessExpirationTime == nil {
		defaultAET := 1
		conf.AccessExpirationTime = &defaultAET
	}

	if conf.RefreshExpirationTime == nil {
		defaultRET := 2
		conf.RefreshExpirationTime = &defaultRET
	}

	return
}

func initLogging() (err error) {
	f, err := os.Create("../logs/log.log")
	if err != nil {
		return systemError.ErrInitFailure
	}
	multiWriter := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multiWriter)
	gin.DefaultWriter = multiWriter
	return
}

func Initialize() (err error) {
	err = initLogging()
	if err != nil {
		return
	}
	conf, err := getConfiguration()
	if err != nil {
		return
	}

	mailService.InitMailService(*conf.ServerEmail, *conf.ServerEmailPW)

	tokenHelper.InitJWT(*conf.JwtKey, *conf.AccessExpirationTime, *conf.RefreshExpirationTime)

	// default is mock db
	err = databaseService.InitDatabase(conf.IsMockDB == nil || *conf.IsMockDB, conf.DatabaseConfig)
	if err != nil {
		return
	}
	// default is mock storage
	return storageService.InitStorage(conf.IsMockStorage == nil || *conf.IsMockStorage)
}
