package web

import (
	"github.com/gin-gonic/gin"
	"socialbot/internal/web/controller/admin"
	"socialbot/internal/web/controller/front"
	middlewares "socialbot/internal/web/middleware"
	"socialbot/internal/web/service/configService"
)

func RegisterRouter(g *gin.Engine) *gin.Engine{
	g.Use(middlewares.LoggerMiddleware())
	g.Use(middlewares.RecoveryMiddleware())
	g.Use(middlewares.CORSMiddleware())
	g.Use(middlewares.ServeThemeView())

	// storage path
	uploadPath := configService.GetStorageUploadPath()
	g.Static("/storage", uploadPath)

	// front
	api := g.Group("/v1/api")
	api.POST("/user/login", front.Login)
	api.POST("/user/register", front.Register)
	api.GET("/category/list", front.ListCategory)
	api.GET("/category/listWithTags", front.ListCategoryWithTags)


	api.Use(middlewares.UserInfo())
	api.GET("/media/homeRecommend", front.HomeRecommendMedias)
	api.GET("/media/listByCategory", front.ListMedias)
	api.GET("/media/detail",front.MediaDetail)

	apiAut := api.Use(middlewares.Auth())
	apiAut.POST("/media/like", front.LikeMedia)
	apiAut.POST("/user/editProfile", front.EditProfile)


	// admin
	adminApi := g.Group("/v1/adminApi")
	adminApi.POST("/login", admin.Login)
	adminApi.GET("/reverse", admin.Reverse)

	adminApi.Use(middlewares.AuthAdmin())
	adminApi.POST("/upload/single", admin.UploadSingle)

	adminApi.POST("/config/addServer", admin.AddServer)
	adminApi.POST("/config/updateServer", admin.UpdateServer)
	adminApi.GET("/config/listServer", admin.ListServer)
	adminApi.POST("/config/deleteServer", admin.DeleteServer)



	adminApi.POST("/robot/addServer", admin.AddRobotServer)
	adminApi.GET("/robot/listServer", admin.ListRobotServer)
	adminApi.POST("/robot/deleteServer", admin.DeleteRobotServer)


	adminApi.POST("/category/add", admin.AddCategory)
	adminApi.POST("/category/update", admin.UpdateCategory)
	adminApi.GET("/category/list", admin.ListCategory)
	adminApi.POST("/category/delete", admin.DeleteCategory)
	adminApi.GET("/category/listWithTags", admin.ListCategoryWithTags)

	adminApi.POST("/copywriter/add", admin.AddCopywriter)
	adminApi.POST("/copywriter/update", admin.UpdateCopywriter)
	adminApi.GET("/copywriter/list", admin.ListCopywriter)
	adminApi.GET("/copywriter/search", admin.SearchCopywriter)
	adminApi.POST("/copywriter/delete", admin.DeleteCopywriter)

	adminApi.POST("/tag/add", admin.AddTag)
	adminApi.POST("/tag/update", admin.UpdateTag)
	adminApi.GET("/tag/list", admin.ListTag)
	adminApi.POST("/tag/delete", admin.DeleteTag)

	adminApi.POST("/crawler/add", admin.AddCrawler)
	adminApi.POST("/crawler/update", admin.UpdateCrawler)
	adminApi.GET("/crawler/list", admin.ListCrawler)
	adminApi.GET("/crawler/listRandItem", admin.ListRandCrawlerItem)
	adminApi.POST("/crawler/deleteItem", admin.DeleteCrawlerItem)

	adminApi.POST("/media/addCommissionProduct", admin.AddCommissionProduct)
	adminApi.POST("/media/addSocialMediaFromCrawler", admin.AddSocialMediaFromCrawler)

	adminApi.GET("/gallery/listGallery", admin.ListGallery)
	adminApi.POST("/gallery/addGalleryTag", admin.AddGalleryTag)

	return g
}