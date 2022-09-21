package tuago

import "github.com/gin-gonic/gin"

type AppEnvPropertyHook func(dest map[string]string) error

type Application struct {
	*gin.Engine
	envProperties       map[string]string    // 环境变量
	appEnvPropertyHooks []AppEnvPropertyHook // 环境变量钩子

	tableChain []TableInf
	cleanChain []CleanAble
	initChain  []InitAble
}
