package emailer

import (
	"crypto/tls"
	"fmt"
	"log"

	gomail "gopkg.in/mail.v2"
)

func SendOtp(from, to, password, otp, subject string) error {
	emailBody := fmt.Sprintf(`
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f5f5f5;
        }

        .container {
            max-width: 600px;
            margin: 0 auto;
            padding: 30px;
            background-color: #fff;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }

        h2 {
            text-align: center;
            margin-bottom: 20px;
            font-size: 24px;
        }

        p {
            line-height: 1.5;
            margin-bottom: 15px;
        }

        .otp-code {
            font-weight: bold;
            color: #007bff;
            font-size: 18px;
            background-color: #f0f0f0;
            padding: 10px;
            border-radius: 5px;
        }

        a {
            color: #007bff;
            text-decoration: none;
        }

        .signature {
            margin-top: 30px;
        }

        .signature p {
            margin-bottom: 5px;
        }
    </style>
</head>
<body>

<div class="container">
    <h2>Connectr - OTP Verification</h2>
    <p>Hi New User,</p>
    <p>Thank you for reaching out to us. As part of our security measures, we require you to authenticate your identity through a One-Time Password (OTP).</p>
    <p style="font-weight: bold;">Your OTP is: <span class="otp-code">%s</span></p>
    <small>
        Please use this OTP to verify your identity and proceed with the necessary action. For security reasons, please do not share this OTP with anyone. If you did not request this OTP, please disregard this email.
    </small>
    <div class="signature">
        <p>Best regards,</p>
        <p><b>Connectr</b></p>
    </div>
</div>

</body>
</html>
`, subject, otp)

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", emailBody)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
