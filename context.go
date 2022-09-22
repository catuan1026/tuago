package tuago

type HandlerFunc func(c *Context)

type OrderHandler struct {
	Order int //排序
	h     HandlerFunc
}

type Context struct {
	actionLabel string
	roleLabel   string
	groupLabel  string
	
	handlerNext bool
	respChan    chan Result
}

func (c *Context) ActionLabel() string {
	return c.actionLabel
}

func (c *Context) RoleLabel() string {
	return c.roleLabel
}

func (c *Context) GroupLabel() string {
	return c.groupLabel
}

func (c *Context) IsHandlerNext() bool {
	return c.handlerNext
}

func (c *Context) HandlerAbort() {
	c.handlerNext = false
}

func (c *Context) Result(errCode int, errMsg string, data ...any) {
	if len(data) == 0 {
		c.respChan <- Result{
			ErrCode: errCode,
			ErrMsg:  errMsg,
		}
		return
	}
	if len(data) == 1 {
		c.respChan <- Result{
			ErrCode: errCode,
			ErrMsg:  errMsg,
			Data:    data[0],
		}
		return
	}
	c.respChan <- Result{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	}
}
