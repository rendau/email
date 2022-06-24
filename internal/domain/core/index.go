package core

import (
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg               logger.Lite
	smtpAddr         string
	smtpAuthUser     string
	smtpAuthPassword string
	smtpAuthHost     string
}

func New(lg logger.Lite, smtpAddr string, smtpAuthUser string, smtpAuthPassword string, smtpAuthHost string) *St {
	return &St{
		lg:               lg,
		smtpAddr:         smtpAddr,
		smtpAuthUser:     smtpAuthUser,
		smtpAuthPassword: smtpAuthPassword,
		smtpAuthHost:     smtpAuthHost,
	}
}
