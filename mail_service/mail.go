package mail_service

import (
	"bytes"
	"fmt"
	"net/smtp"

	"text/template"

	authModel "github.com/leechongyan/Studtor_backend/authentication_service/models"
	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
	"github.com/spf13/viper"
)

func SendVerificationCode(user authModel.User, code string) (err error) {
	serverEmail := viper.GetString("serverEmail")
	serverEmailPW := viper.GetString("serverEmailPW")

	// Sender data.
	from := serverEmail
	password := serverEmailPW

	// Receiver email address.
	to := []string{
		*user.Email,
	}
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("../mail_service/templates/verification_template.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Verification Code for Studtor \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    *user.FirstName + " " + *user.LastName,
		Message: code,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return systemError.ErrEmailSendingFailure
	}

	return
}
