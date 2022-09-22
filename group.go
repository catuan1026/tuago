package tuago

type GroupInf interface {
	GroupLabel() string
	RoleLabel() string
	Call(c *Context)
	FindAction(actionLabel string) (HandlerFunc, bool)
	BindAction(actionLabel string, h HandlerFunc)
}

type Group struct {
	groupLabel string
	roleLabel  string
	actions    map[string]HandlerFunc

	beforeHandlers []OrderHandler
	afterHandlers  []OrderHandler
}

func NewGroup(groupLabel string, roleLabel string) *Group {
	return &Group{
		groupLabel:     groupLabel,
		roleLabel:      roleLabel,
		actions:        make(map[string]HandlerFunc),
		beforeHandlers: make([]OrderHandler, 0),
		afterHandlers:  make([]OrderHandler, 0),
	}
}

func (g *Group) GroupLabel() string {
	return g.groupLabel
}

func (g *Group) RoleLabel() string {
	return g.roleLabel
}

func (g *Group) Call(c *Context) {
	for _, h := range g.beforeHandlers {
		h.h(c)
	}
	if !c.IsHandlerNext() {
		return
	}
	actionLabel := c.ActionLabel()
	if h, ok := g.actions[actionLabel]; ok {
		h(c)
	}
	if !c.IsHandlerNext() {
		return
	}
	for _, h := range g.afterHandlers {
		h.h(c)
	}
}

func (g *Group) FindAction(actionLabel string) (HandlerFunc, bool) {
	h, ok := g.actions[actionLabel]
	return h, ok
}

func (g *Group) BindAction(actionLabel string, h HandlerFunc) {
	if _, ok := g.actions[actionLabel]; ok {
		panic("actionLabel is already exist")
	}
	g.actions[actionLabel] = h
}

func (g *Group) UseBeforeHandler(order int, h HandlerFunc) {
	if g.beforeHandlers == nil {
		g.beforeHandlers = make([]OrderHandler, 0)
	}
	g.beforeHandlers = append(g.beforeHandlers, OrderHandler{Order: order, h: h})
	//按照order排序
	amount := len(g.beforeHandlers)
	for i := 0; i < amount; i++ {
		for j := i + 1; j < amount; j++ {
			if g.beforeHandlers[i].Order > g.beforeHandlers[j].Order {
				g.beforeHandlers[i], g.beforeHandlers[j] = g.beforeHandlers[j], g.beforeHandlers[i]
			}
		}
	}
}

func (g *Group) UseAfterHandler(order int, h HandlerFunc) {
	if g.afterHandlers == nil {
		g.afterHandlers = make([]OrderHandler, 0)
	}
	g.afterHandlers = append(g.afterHandlers, OrderHandler{Order: order, h: h})
	//按照order排序
	amount := len(g.afterHandlers)
	for i := 0; i < amount; i++ {
		for j := i + 1; j < amount; j++ {
			if g.afterHandlers[i].Order > g.afterHandlers[j].Order {
				g.afterHandlers[i], g.afterHandlers[j] = g.afterHandlers[j], g.afterHandlers[i]
			}
		}
	}
}