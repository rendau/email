package httpapi

import (
	"net/http"

	"github.com/rendau/email/internal/domain/entities"
)

func (a *St) hSend(w http.ResponseWriter, r *http.Request) {
	var err error

	reqObj := &entities.SendReqSt{}

	if !a.uParseRequestJSON(w, r, reqObj) {
		return
	}

	if reqObj.Sync {
		err = a.cr.SendMail(reqObj)
		if err != nil {
			a.uHandleError(err, w)
			return
		}
	} else {
		go func() { _ = a.cr.SendMail(reqObj) }()
	}

	w.WriteHeader(200)
}
