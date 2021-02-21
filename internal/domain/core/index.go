package core

import (
	"github.com/rendau/email/internal/interfaces"
)

type St struct {
	lg               interfaces.Logger
	smtpAddr         string
	smtpAuthUser     string
	smtpAuthPassword string
	smtpAuthHost     string
}

func New(lg interfaces.Logger, smtpAddr string, smtpAuthUser string, smtpAuthPassword string, smtpAuthHost string) *St {
	return &St{
		lg:               lg,
		smtpAddr:         smtpAddr,
		smtpAuthUser:     smtpAuthUser,
		smtpAuthPassword: smtpAuthPassword,
		smtpAuthHost:     smtpAuthHost,
	}
}
