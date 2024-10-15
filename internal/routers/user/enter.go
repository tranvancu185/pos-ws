package user

type UserRouterGroup struct {
	AuthRouter
	UserRouter
	ProductRouter
	TableRouter
}
