package tuago

type RoleInf interface {
	RoleLabel() string
	Call(c *Context)
}

type Role struct {
	roleLabel string
	handlers  []OrderHandler
}

func NewRole(roleLabel string) *Role {
	return &Role{
		roleLabel: roleLabel,
		handlers:  make([]OrderHandler, 0),
	}
}

func (r *Role) RoleLabel() string {
	return r.roleLabel
}

func (r *Role) Call(c *Context) {
	for _, h := range r.handlers {
		h.h(c)
	}
}

func (r *Role) UseHandler(order int, h HandlerFunc) {
	if r.handlers == nil {
		r.handlers = make([]OrderHandler, 0)
	}
	r.handlers = append(r.handlers, OrderHandler{Order: order, h: h})
	//按照order排序
	amount := len(r.handlers)
	for i := 0; i < amount; i++ {
		for j := i + 1; j < amount; j++ {
			if r.handlers[i].Order > r.handlers[j].Order {
				r.handlers[i], r.handlers[j] = r.handlers[j], r.handlers[i]
			}
		}
	}
}
