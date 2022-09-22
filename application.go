package tuago

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	defaultHttpPort  = "8088"
	defaultHttpsPort = "8089"
)

type AppEnvPropertyHook func(dest map[string]string) error

type AppHandlerFunc func(app *Application) error

type Application struct {
	*gin.Engine
	envProperties       map[string]string    // 环境变量
	appEnvPropertyHooks []AppEnvPropertyHook // 环境变量钩子

	preAppHandlerFuncs  []AppHandlerFunc // 启动前钩子
	afterAppHandlerFunc []AppHandlerFunc // 启动后钩子

	tableChain []TableInf
	cleanChain []CleanAble
	initChain  []InitAble
}

// UseApplicationPreStartHandler 应用启动前钩子
func (app *Application) UseApplicationPreStartHandler(h AppHandlerFunc) {
	app.preAppHandlerFuncs = append(app.preAppHandlerFuncs, h)
}

func (app *Application) runAppPreHandler() error {
	if app.preAppHandlerFuncs != nil {
		for _, h := range app.preAppHandlerFuncs {
			err := h(app)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"tip": "应用启动前钩子执行异常",
				}).Error(err.Error())
			}
		}
	}
	return nil
}

// UseApplicationAfterStartHandler 应用启动后钩子
func (app *Application) UseApplicationAfterStartHandler(h AppHandlerFunc) {
	app.afterAppHandlerFunc = append(app.afterAppHandlerFunc, h)
}

func (app *Application) runApplicationAfterHandler() error {
	if app.afterAppHandlerFunc != nil {
		for _, h := range app.afterAppHandlerFunc {
			err := h(app)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"tip": "应用启动后钩子执行异常",
				}).Error(err.Error())
			}
		}
	}
	return nil
}

// Start 启动服务
func (app *Application) Start(ctx context.Context) {
	os.Getpid()
}

func (app *Application) httpService(ctx context.Context) error {
	return nil
}

func (app *Application) httpsService(ctx context.Context) error {
	return nil
}
