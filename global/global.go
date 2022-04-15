package global

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

const(
	Success = 200

	ErrorGeneral = 500 //常规Error
	ErrorJWTCheck = 10001 //JWT验证错误
	ErrorParams = 10002 //参数获取错误
	ErrorUsers = 10003 //用户类操作错误
	ErrorChats = 10004 //聊天操作错误
)


func UnifiedReturn(c *gin.Context,code int,msg interface{},data interface{}, token string){
	str := msg
	if code == Success||msg == nil{
		str = "ok"
	}

	c.Header("AuthToken",token) //token为空直接没有这个header，好评
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":    str,
		"data":   data,
		// "token":  token,
	})
}