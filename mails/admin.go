package mails

import (
	"fmt"
	"log"
	"net/mail"

	"github.com/guille1093/G-Base/core"
	"github.com/guille1093/G-Base/mails/templates"
	"github.com/guille1093/G-Base/models"
	"github.com/guille1093/G-Base/tokens"
	"github.com/guille1093/G-Base/tools/mailer"
	"github.com/guille1093/G-Base/tools/rest"
)

// SendAdminPasswordReset sends a password reset request email to the specified admin.
func SendAdminPasswordReset(app core.App, admin *models.Admin) error {
	token, tokenErr := tokens.NewAdminResetPasswordToken(app, admin)
	if tokenErr != nil {
		return tokenErr
	}

	actionUrl, urlErr := rest.NormalizeUrl(fmt.Sprintf(
		"%s/_/#/confirm-password-reset/%s",
		app.Settings().Meta.AppUrl,
		token,
	))
	if urlErr != nil {
		return urlErr
	}

	params := struct {
		AppName   string
		AppUrl    string
		Admin     *models.Admin
		Token     string
		ActionUrl string
	}{
		AppName:   app.Settings().Meta.AppName,
		AppUrl:    app.Settings().Meta.AppUrl,
		Admin:     admin,
		Token:     token,
		ActionUrl: actionUrl,
	}

	mailClient := app.NewMailClient()

	// resolve body template
	body, renderErr := resolveTemplateContent(params, templates.Layout, templates.AdminPasswordResetBody)
	if renderErr != nil {
		return renderErr
	}

	message := &mailer.Message{
		From: mail.Address{
			Name:    app.Settings().Meta.SenderName,
			Address: app.Settings().Meta.SenderAddress,
		},
		To:      mail.Address{Address: admin.Email},
		Subject: "Reset admin password",
		HTML:    body,
	}

	event := &core.MailerAdminEvent{
		MailClient: mailClient,
		Message:    message,
		Admin:      admin,
		Meta:       map[string]any{"token": token},
	}

	sendErr := app.OnMailerBeforeAdminResetPasswordSend().Trigger(event, func(e *core.MailerAdminEvent) error {
		return e.MailClient.Send(e.Message)
	})

	if sendErr == nil {
		if err := app.OnMailerAfterAdminResetPasswordSend().Trigger(event); err != nil && app.IsDebug() {
			log.Println(err)
		}
	}

	return sendErr
}
