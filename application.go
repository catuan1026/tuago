package tuago

import "github.com/gin-gonic/gin"

type AppEnvPropertyHook func(dest map[string]string) error

type AppHandlerFunc func(app *Application)

type Application struct {
	*gin.Engine
	envProperties       map[string]string    // 环境变量
	appEnvPropertyHooks []AppEnvPropertyHook // 环境变量钩子
	preAppHandlerFuncs  []AppHandlerFunc     // 启动前钩子
	afterAppHandlerFunc []AppHandlerFunc     // 启动后钩子

	tableChain []TableInf
	cleanChain []CleanAble
	initChain  []InitAble
}

func (app *Application) UsePre(h AppHandlerFunc) {
	app.preAppHandlerFuncs = append(app.preAppHandlerFuncs, h)
}
