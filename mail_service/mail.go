package mail_service

import (
	"bytes"
	"fmt"
	"net/smtp"
	"time"

	"text/template"

	systemError "github.com/leechongyan/Studtor_backend/constants/errors/system_errors"
	userModel "github.com/leechongyan/Studtor_backend/database_service/client_models"
	"github.com/spf13/viper"
)

var CurrentMailService MailService

type MailService struct {
	serverEmail string
	smtpHost    string
	smtpPort    string
	smtpAuth    smtp.Auth
	mimeHeaders string
}

func InitMailService() {
	CurrentMailService = MailService{}
	CurrentMailService.serverEmail = viper.GetString("serverEmail")
	CurrentMailService.smtpHost = "smtp.gmail.com"
	CurrentMailService.smtpPort = "587"
	CurrentMailService.mimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	CurrentMailService.smtpAuth = smtp.PlainAuth("", CurrentMailService.serverEmail, viper.GetString("serverEmailPW"), CurrentMailService.smtpHost)
}

func (mailService MailService) SendVerificationCode(user userModel.User, code string) (err error) {
	// Sender data.
	// Receiver email address.
	to := []string{
		*user.Email(),
	}

	t, _ := template.ParseFiles("../mail_service/templates/verification_template.html")
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: Verification Code for Studtor \n%s\n\n", mailService.mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    *user.FirstName() + " " + *user.LastName(),
		Message: code,
	})

	// Sending email.
	err = smtp.SendMail(mailService.smtpHost+":"+mailService.smtpPort, mailService.smtpAuth, mailService.serverEmail, to, body.Bytes())
	if err != nil {
		return systemError.ErrEmailSendingFailure
	}

	return
}

func (mailService MailService) SendBookingConfirmation(student userModel.User, tutor userModel.User, courseName string, date time.Time, time string) (err error) {
	// send to student
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: Confirmation of Appointment\n%s\n\n", mailService.mimeHeaders)))
	to := []string{
		*student.Email(),
	}
	t, _ := template.ParseFiles("../mail_service/templates/student_confirmation_template.html")
	t.Execute(&body, struct {
		Name   string
		Tutor  string
		Course string
		Date   string
		Time   string
	}{
		Name:   *student.FirstName() + " " + *student.LastName(),
		Tutor:  *tutor.FirstName() + " " + *tutor.LastName(),
		Course: courseName,
		Date:   date.Format("Jan 2, 2006"),
		Time:   time,
	})

	err = smtp.SendMail(mailService.smtpHost+":"+mailService.smtpPort, mailService.smtpAuth, mailService.serverEmail, to, body.Bytes())
	if err != nil {
		return systemError.ErrEmailSendingFailure
	}

	body.Reset()

	to = []string{
		*tutor.Email(),
	}
	t, _ = template.ParseFiles("../mail_service/templates/tutor_confirmation_template.html")
	t.Execute(&body, struct {
		Name    string
		Student string
		Course  string
		Date    string
		Time    string
	}{
		Name:    *tutor.FirstName() + " " + *tutor.LastName(),
		Student: *student.FirstName() + " " + *student.LastName(),
		Course:  courseName,
		Date:    date.Format("Jan 2, 2006"),
		Time:    time,
	})

	err = smtp.SendMail(mailService.smtpHost+":"+mailService.smtpPort, mailService.smtpAuth, mailService.serverEmail, to, body.Bytes())
	if err != nil {
		return systemError.ErrEmailSendingFailure
	}
	return
}

func (mailService MailService) SendBookingCancellation(student userModel.User, tutor userModel.User, courseName string, date time.Time, time string) (err error) {
	// send to student
	var body bytes.Buffer
	body.Write([]byte(fmt.Sprintf("Subject: Cancellation of Appointment\n%s\n\n", mailService.mimeHeaders)))
	to := []string{
		*student.Email(),
	}
	t, _ := template.ParseFiles("../mail_service/templates/student_cancellation_template.html")
	t.Execute(&body, struct {
		Name   string
		Tutor  string
		Course string
		Date   string
		Time   string
	}{
		Name:   *student.FirstName() + " " + *student.LastName(),
		Tutor:  *tutor.FirstName() + " " + *tutor.LastName(),
		Course: courseName,
		Date:   date.Format("Jan 2, 2006"),
		Time:   time,
	})

	err = smtp.SendMail(mailService.smtpHost+":"+mailService.smtpPort, mailService.smtpAuth, mailService.serverEmail, to, body.Bytes())
	if err != nil {
		return systemError.ErrEmailSendingFailure
	}

	body.Reset()

	to = []string{
		*tutor.Email(),
	}
	t, _ = template.ParseFiles("../mail_service/templates/tutor_cancellation_template.html")
	t.Execute(&body, struct {
		Name    string
		Student string
		Course  string
		Date    string
		Time    string
	}{
		Name:    *tutor.FirstName() + " " + *tutor.LastName(),
		Student: *student.FirstName() + " " + *student.LastName(),
		Course:  courseName,
		Date:    date.Format("Jan 2, 2006"),
		Time:    time,
	})

	err = smtp.SendMail(mailService.smtpHost+":"+mailService.smtpPort, mailService.smtpAuth, mailService.serverEmail, to, body.Bytes())
	if err != nil {
		return systemError.ErrEmailSendingFailure
	}
	return
}
