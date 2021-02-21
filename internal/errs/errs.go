package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ReceiversEmpty  = Err("receivers_is_empty")
	BadEmailFormat  = Err("bad_email_format")
	MessageRequired = Err("message_required")
	SubjectRequired = Err("subject_required")
	ServiceNA       = Err("service_not_available")
)
