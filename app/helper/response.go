package helper

import (
	"github.com/gin-gonic/gin"
)

type ErrorResp struct {
	ErrorString    string                 `json:"error"`
	Code           int                    `json:"code"`
	Message        string                 `json:"message"`
	AdditionalData map[string]interface{} `json:"-"`
	Tags           map[string]string      `json:"-"`
	Error          error                  `json:"-"`
}

func (err *ErrorResp) HandleError(c *gin.Context) {
	err.ErrorString = err.Error.Error()
	c.AbortWithStatusJSON(err.Code, err)
}
