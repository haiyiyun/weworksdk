package svr

import (
	"github.com/gin-gonic/gin"
	"github.com/haiyiyun/weworksdk/pkg/svr/callback"
	"github.com/haiyiyun/weworksdk/pkg/svr/install"
)

func InjectRouter(e *gin.Engine) {

	callbackGroup := e.Group("/callback")
	{
		callbackGroup.GET("/data", callback.DataGetHandler)
		callbackGroup.POST("/data", callback.DataPostHandler)
		callbackGroup.GET("/cmd", callback.CmdGetHandler)
		callbackGroup.POST("/cmd", callback.CmdPostHandler)
		callbackGroup.GET("/customized", callback.CustomizedGetHandler)
		callbackGroup.POST("/customized", callback.CustomizedPostHandler)
	}
	suite := e.Group("/suite")
	{
		suite.GET("/install", install.SuiteInstall)
		suite.GET("/install/auth", install.SuiteInstallAuth)
	}
}
