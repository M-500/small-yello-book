package web

import (
	_ "gin-svc/docs"
	"gin-svc/internal"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/service"
	"gin-svc/internal/web/controller"
	"gin-svc/internal/web/middleware"
	"gin-svc/internal/web/middleware/jwt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupWebEngine(app *internal.App) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.CorsMdl())
	engine.Static("/static", "./static")
	rg := engine.Group("/api/v1")
	publicGroup := rg.Group("/na")
	{
		publicGroup.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		InitPubController(app).RegisterRoute(publicGroup)
	}

	privateGroup := rg.Group("")
	{
		InitUserController(app).RegisterRoute(privateGroup)
		InitRoleController(app).RegisterRoute(privateGroup)
		InitFileController(app).RegisterRoute(privateGroup)

	}

	return engine
}
func InitFileController(app *internal.App) *controller.FileController {
	return controller.NewFileController()
}

func InitUserController(app *internal.App) *controller.UserController {
	userDao := dao.NewUserDao(app.DB)
	userRepo := repo.NewUserRepoInterface(userDao)

	codeCache := cache.NewRedisVerificationCodeCache(app.Cli)
	jwtHdl := jwt.NewRedisJWTHandler(app.Cli)
	userSvc := service.NewUserSvc(userRepo, codeCache, jwtHdl)
	return controller.NewUserController(userSvc)
}

func InitPubController(app *internal.App) *controller.PublicController {
	userDao := dao.NewUserDao(app.DB)
	userRepo := repo.NewUserRepoInterface(userDao)
	codeCache := cache.NewRedisVerificationCodeCache(app.Cli)
	jwtHdl := jwt.NewRedisJWTHandler(app.Cli)
	userSvc := service.NewUserSvc(userRepo, codeCache, jwtHdl)
	redisVerificationCodeCache := cache.NewRedisVerificationCodeCache(app.Cli)
	emailSvc := service.NewSMTPService(redisVerificationCodeCache)
	return controller.NewPublicController(userSvc, emailSvc)
}

func InitRoleController(app *internal.App) controller.BaseController {
	perDao := dao.NewPermissionDAO(app.DB)
	roleDao := dao.NewRoleDAO(app.DB)
	cac := cache.NewRoleCache(app.Cli)
	roleRepo := repo.NewRoleRepo(perDao, roleDao, cac)
	roleSvc := service.NewRoleService(roleRepo)
	return controller.NewRoleController(roleSvc)
}

func InitNoteController(app *internal.App) *controller.NoteCtl {
	dao := dao.NewNoteDao(app.DB)
	repo := repo.NewNoteRepo(dao)
	svc := service.NewNoteSvcImpl(repo)
	return controller.NewNoteCtl(svc)
}
