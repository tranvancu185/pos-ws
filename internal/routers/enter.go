package routers

import (
	"tranvancu185/vey-pos-ws/internal/routers/app"
	"tranvancu185/vey-pos-ws/internal/routers/manager"
	"tranvancu185/vey-pos-ws/internal/routers/user"
)

type RouterGroup struct {
	App     app.AppRouterGroup
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
