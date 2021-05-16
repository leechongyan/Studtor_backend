package mail_service 

import (
	"crypto/tls"
	"fmt"

	"github.com/leechongyan/Studtor_backend/authentication_service/models"
	"github.com/spf13/viper"
  
	gomail "gopkg.in/mail.v2"
  )
  
  func SendVerificationCode(user models.User, code string) {
	m := gomail.NewMessage()

	serverEmail := viper.GetString("serverEmail") 
	serverEmailPW := viper.GetString("serverEmailPW") 
  
	// Set E-Mail sender
	m.SetHeader("From", serverEmail)
  
	// Set E-Mail receivers
	m.SetHeader("To", *user.Email)
  
	// Set E-Mail subject
	m.SetHeader("Subject", "Verify your email")
  
	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", fmt.Sprintf("This is your verification code: %s", code))
  
	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, serverEmail, serverEmailPW)
  
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  
	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
	  fmt.Println(err)
	  panic(err)
	}
  
	return
  }