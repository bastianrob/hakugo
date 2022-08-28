package notification

import "html/template"

var mailVerificationTemplate *template.Template

func init() {
	initializeMailVerificationTemplate()
}

type MailVerificationTemplate struct {
	Name  string
	Email string
	Code  string
	Host  string
}

func initializeMailVerificationTemplate() {
	template, _ := template.New("mail-verification").Parse(`
	<!doctype html>
	<html>

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
	</head>

	<body
		style="font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;">
		<div style="color: black !important;">
			<h2>
				Hi {{.Name}},
			</h2>

			<p>
				We're very excited for you to join us!<br />
				Please confirm you email address by clicking the button below
			</p><br />
		</div>

		<div style="color: white !important;">
			<a href="{{.Host}}/register/confirmation?email={{.Email}}&code={{.Code}}"
				style="text-decoration: none; color: white; background: rgb(9, 169, 222); padding: 12px 20px; border-radius: 6px;">
				Verify Now
			</a>
		</div>

		<br />
		<div style="color: black !important;">
			<p>
				This verification link will expire in 24 hours.
			</p>
		</div>
	</body>

	</html>
	`)

	mailVerificationTemplate = template
}
