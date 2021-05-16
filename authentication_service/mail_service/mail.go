package mail_service 

import (
	"fmt"
	"bytes"
	"net/smtp"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/spf13/viper"
	"text/template"  
  )
  
  func SendVerificationCode(user models.User, code string) {
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

	t, _ := template.ParseFiles("../mail_service/template.html")
	var body bytes.Buffer
  	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Verification Code for Studtor \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	  }{
		Name:    *user.First_name + " " + *user.Last_name,
		Message: code,
	  })

	  // Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}  

	return
  }