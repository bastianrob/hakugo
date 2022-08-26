package notification

import "text/template"

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
	Hi {{.Name}},

	We're very excited for you to join us!
	Please confirm you email address by clicking the link below

	{{.Host}}/register/confirmation?email={{.Email}}&code={{.Code}}

	This verification link will exprire in 24 hours.
	`)

	mailVerificationTemplate = template
}
