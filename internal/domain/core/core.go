package core

import (
	"net/smtp"
	"regexp"
	"strings"

	"github.com/rendau/dop/dopErrs"
	"github.com/rendau/email/internal/domain/types"
	"github.com/rendau/email/internal/errs"
)

var emailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,10}$`)

func (c *St) SendMail(pars *types.SendReqSt) error {
	var err error

	err = c.validateValues(pars)
	if err != nil {
		return err
	}

	headers := []string{
		`From: ` + c.smtpAuthUser,
		`To: ` + strings.Join(pars.Receivers, `,`),
		`Subject:` + pars.Subject,
		`MIME-version: 1.0;`,
		`Content-Type: text/html; charset="UTF-8"`,
	}

	msg := []byte(strings.Join(headers, "\r\n") + "\r\n\r\n" + pars.Message)

	err = smtp.SendMail(
		c.smtpAddr,
		smtp.PlainAuth(
			"",
			c.smtpAuthUser,
			c.smtpAuthPassword,
			c.smtpAuthHost,
		),
		c.smtpAuthUser,
		pars.Receivers,
		msg,
	)
	if err != nil {
		c.lg.Errorw("fail to send mail:", err)
		return dopErrs.ServiceNA
	}

	// c.lg.Info("Messages have been sent successfully")

	return nil
}

func (c *St) validateValues(pars *types.SendReqSt) error {
	for _, item := range pars.Receivers {
		if !c.validateEmail(item) {
			c.lg.Warnw("Bad email format", errs.BadEmailFormat)
			return errs.BadEmailFormat
		}
	}
	if len(pars.Receivers) == 0 {
		c.lg.Warnw("Receivers is empty", errs.ReceiversEmpty)
		return errs.ReceiversEmpty
	}
	if len(pars.Message) == 0 {
		c.lg.Warnw("Message is empty", errs.MessageRequired)
		return errs.MessageRequired
	}
	if len(pars.Subject) == 0 {
		c.lg.Warnw("Subject is empty", errs.SubjectRequired)
		return errs.SubjectRequired
	}
	return nil
}

func (c *St) validateEmail(v string) bool {
	return emailRegexp.MatchString(v)
}
