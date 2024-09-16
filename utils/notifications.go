package utils

import (
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
    "github.com/twilio/twilio-go"
    "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendEmail(to string, subject string, body string) error {
    from := mail.NewEmail("Movie Reservation System", "no-reply@moviereservation.com")
    toEmail := mail.NewEmail("", to)
    message := mail.NewSingleEmail(from, subject, toEmail, body, body)
    client := sendgrid.NewSendClient("your_sendgrid_api_key")
    _, err := client.Send(message)
    return err
}

func SendSMS(to string, body string) error {
    client := twilio.NewRestClient("your_twilio_account_sid", "your_twilio_auth_token")
    params := &v2010.CreateMessageParams{}
    params.SetTo(to)
    params.SetFrom("your_twilio_phone_number")
    params.SetBody(body)
    _, err := client.ApiV2010.CreateMessage(params)
    return err
}
