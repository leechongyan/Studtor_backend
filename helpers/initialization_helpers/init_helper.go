package initialization_helpers

import (
	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
	databaseService "github.com/leechongyan/Studtor_backend/database_service/controller"
	storageService "github.com/leechongyan/Studtor_backend/storage_service"
	"github.com/spf13/viper"
)

func initializeViper() (err error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return systemError.ErrInitFailure
	}
	return
}

// func checkConfigInputs() (error error) {

// }

func Initialize() (err error) {
	err = initializeViper()
	if err != nil {
		return
	}
	err = databaseService.InitDatabase()
	if err != nil {
		return
	}
	err = storageService.InitStorage()
	return err
}
