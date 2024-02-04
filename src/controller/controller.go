package controller

import (
	"go-shop-api/src/services/group"
	groupitem "go-shop-api/src/services/group-item"
	grouptype "go-shop-api/src/services/group-type"
	"go-shop-api/src/services/item"
	itemtype "go-shop-api/src/services/item-type"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	v1 := app.Group("v1")
	initAdminControllers(v1)
	initClientControllers(v1)
}

func initAdminControllers(app fiber.Router) {
	admin := app.Group(baseRoutes["admin"])

	admin.Get(itemTypeRoutes["type"], itemtype.GetAll)
	admin.Get(itemTypeRoutes["typeById"], itemtype.GetById)
	admin.Post(itemTypeRoutes["type"], itemtype.Insert)
	admin.Put(itemTypeRoutes["typeById"], itemtype.UpdateById)
	admin.Delete(itemTypeRoutes["typeById"], itemtype.DeleteById)

	admin.Get(itemRoutes["item"], item.GetAll)
	admin.Get(itemRoutes["itemById"], item.GetById)
	admin.Post(itemRoutes["item"], item.Insert)
	admin.Put(itemRoutes["itemById"], item.UpdateById)
	admin.Delete(itemRoutes["itemById"], item.DeleteById)

	admin.Get(groupTypeRoutes["type"], grouptype.GetAll)
	admin.Get(groupTypeRoutes["typeById"], grouptype.GetById)
	admin.Post(groupTypeRoutes["type"], grouptype.Insert)
	admin.Put(groupTypeRoutes["typeById"], grouptype.UpdateById)
	admin.Delete(groupTypeRoutes["typeById"], grouptype.DeleteById)

	admin.Get(groupRoutes["group"], group.GetAll)
	admin.Get(groupRoutes["groupById"], group.GetById)
	admin.Post(groupRoutes["group"], group.Insert)
	admin.Put(groupRoutes["groupById"], group.UpdateById)
	admin.Delete(groupRoutes["groupById"], group.DeleteById)

	admin.Get(groupItemRoutes["groupItem"], groupitem.GetByGroupId)
	admin.Post(groupItemRoutes["groupItem"], groupitem.Insert)
	admin.Delete(groupItemRoutes["groupItem"], groupitem.DeleteByGroupIdAndItempId)
}

func initClientControllers(app fiber.Router) {
	_ = app.Group(baseRoutes["client"])
}
