package emails

import (
	"bytes"
	"html/template"
	"net/http"
	"ragamaya-api/emails/dto"
	"ragamaya-api/models"
	"ragamaya-api/pkg/config"
	"ragamaya-api/pkg/exceptions"
	"ragamaya-api/pkg/logger"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(data dto.EmailRequest) *exceptions.Exception {
	email := config.GetEmail()
	password := config.GetEmailPassword()
	server := config.GetEmailServer()
	smtpPort := config.GetEmailPort()

	i, err := strconv.Atoi(smtpPort)
	if err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrInternalServer)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	d := gomail.NewDialer(server, i, email, password)

	if err := d.DialAndSend(m); err != nil {
		return exceptions.NewException(http.StatusBadGateway, err.Error())
	}

	return nil
}

func ExampleEmail(data dto.EmailExample) *exceptions.Exception {
	tmpl, exc := template.ParseFiles("emails/templates/example.html")
	if exc != nil {
		return exceptions.NewException(http.StatusInternalServerError, exc.Error())
	}

	var body bytes.Buffer
	if exc := tmpl.Execute(&body, data); exc != nil {
		return exceptions.NewException(http.StatusInternalServerError, exc.Error())
	}

	emailData := dto.EmailRequest{
		Email:   data.Email,
		Subject: data.Subject,
		Body:    body.String(),
	}

	err := SendEmail(emailData)
	if err != nil {
		return err
	}

	logger.Info("Email sent successfully to %s", data.Email)

	return nil
}

func VerificationEmail(data dto.EmailVerification) *exceptions.Exception {
	tmpl, exc := template.ParseFiles("emails/templates/verification.html")
	if exc != nil {
		return exceptions.NewException(http.StatusInternalServerError, exc.Error())
	}

	var body bytes.Buffer
	if exc := tmpl.Execute(&body, data); exc != nil {
		return exceptions.NewException(http.StatusInternalServerError, exc.Error())
	}

	emailData := dto.EmailRequest{
		Email:   data.Email,
		Subject: "[Ragamaya] Email Verification for Account Activation",
		Body:    body.String(),
	}

	err := SendEmail(emailData)
	if err != nil {
		return err
	}

	return nil
}

func PayoutNotificationEmail(email string, data models.WalletPayoutRequest) *exceptions.Exception {
	tmpl, exc := template.ParseFiles("emails/templates/payout_notification.html")
	if exc != nil {
		return exceptions.NewException(http.StatusInternalServerError, exc.Error())
	}

	var body bytes.Buffer
	if exc := tmpl.Execute(&body, data); exc != nil {
		return exceptions.NewException(http.StatusInternalServerError, exc.Error())
	}

	emailData := dto.EmailRequest{
		Email:   email,
		Subject: "[Ragamaya] Payout Notification",
		Body:    body.String(),
	}

	err := SendEmail(emailData)
	if err != nil {
		return err
	}

	return nil
}