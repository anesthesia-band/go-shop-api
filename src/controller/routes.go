package controller

var (
	baseRoutes = map[string]string{
		"client": "/api",
		"admin":  "/admin/api",
	}
	itemRoutes = map[string]string{
		"item":     "/items",
		"itemById": "/items/:id",
	}
	itemTypeRoutes = map[string]string{
		"type":     "/items/types",
		"typeById": "/items/types/:id",
	}
	groupRoutes = map[string]string{
		"group":     "/groups",
		"groupById": "/groups/:id",
	}
	groupTypeRoutes = map[string]string{
		"type":     "/group/types",
		"typeById": "/group/types/:id",
	}
	groupItemRoutes = map[string]string{
		"groupItem": "/items/groups",
	}
)
