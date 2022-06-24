package rest

import (
	"github.com/gin-gonic/gin"
	dopHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/email/internal/domain/types"
)

// @Router   /send [post]
// @Tags     general
// @Param    body  body  types.SendReqSt  false  "body"
// @Produce  json
// @Success  200
// @Failure  400  {object}  dopTypes.ErrRep
func (s *St) hSend(c *gin.Context) {
	reqObj := &types.SendReqSt{}
	if !dopHttps.BindJSON(c, reqObj) {
		return
	}

	if reqObj.Sync {
		dopHttps.Error(c, s.core.SendMail(reqObj))
	} else {
		go func() { _ = s.core.SendMail(reqObj) }()
	}
}
