package api

import (
	"github.com/gin-gonic/gin"
	"go-api/lib"
)

func MpAuth(c *gin.Context) {
	code := c.PostForm("code")
	encryptedData := c.PostForm("encrypted_data")
	iv := c.PostForm("iv")

	session, _ := lib.Minprogram.GetAuth().Code2Session(code)
	data, _ := lib.Minprogram.GetEncryptor().Decrypt(session.SessionKey, encryptedData, iv)
	lib.SuccessJson(c, data)
}
