package errs

import (
	"github.com/rendau/dop/dopErrs"
)

const (
	ReceiversEmpty  = dopErrs.Err("receivers_is_empty")
	BadEmailFormat  = dopErrs.Err("bad_email_format")
	MessageRequired = dopErrs.Err("message_required")
	SubjectRequired = dopErrs.Err("subject_required")
)
