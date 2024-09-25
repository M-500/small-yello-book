package web

import (
	_ "gin-svc/docs"
	"gin-svc/internal"
	"gin-svc/internal/repo"
	"gin-svc/internal/repo/cache"
	"gin-svc/internal/repo/dao"
	"gin-svc/internal/service"
	"gin-svc/internal/types"
	"gin-svc/internal/web/controller"
	"gin-svc/internal/web/middleware"
	"gin-svc/internal/web/middleware/jwt"
	"gin-svc/pkg/ginx"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupWebEngine(app *internal.App) *gin.Engine {
	engine := gin.Default()

	jwtHdl := jwt.NewRedisJWTHandler(app.Cli, app.Cfg)
	builder := middleware.NewJwtMiddlewareBuilder(jwtHdl, app.Cfg.Jwt).
		IgnorePath("/api/v1/file/upload").
		IgnorePath("/api/v1/feed/notes").
		IgnorePath("/api/v1/notes/detail/:uuid").
		Build()
	engine.Use(middleware.CorsMdl())
	engine.Static("/static", "./static")
	rg := engine.Group("/api/v1")
	publicGroup := rg.Group("/na")
	p := InitPubController(app)
	{
		publicGroup.POST("/login", ginx.WrapJsonBody[types.EmailLoginForm](p.EmailLoginCtl))
		publicGroup.POST("/email/send", ginx.WrapJsonBody[types.EmailForm](p.EmailSendCtl))
		publicGroup.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	privateGroup := rg.Group("")
	privateGroup.Use(builder)
	u := InitUserController(app)
	f := InitFileController(app)
	n := InitNoteController(app)
	intrCtl := SetupInteracticeController(app)
	//social:=
	//roleCtl := InitRoleController(app)
	commentCtl := SetupCommentController(app)
	{
		privateGroup.POST("/file/upload", ginx.WrapResponse(f.UploadImgCtl))
		privateGroup.GET("/file", ginx.WrapResponse(f.ReadFileCtl))

		privateGroup.PUT("/users", ginx.WrapJsonBody[types.UserForm](u.UpdateUserInfo))
		privateGroup.GET("/users/:uuid", ginx.WrapResponse(u.FindUserInfo))
		privateGroup.GET("/users/info", ginx.WrapJWT[jwt.UserClaims](u.AdminUserDetail)) // 查询自身用户信息
		privateGroup.DELETE("/users/:id", ginx.WrapJsonBody[types.UserForm](u.DeleteUserCtl))

		//privateGroup.GET("/roles", ginx.WrapQueryBody[types.RolePageListReq](r.RolePageList)) // 获取角色列表
		//privateGroup.GET("/roles/:id", ginx.WrapResponse(r.DetailRoleCtl))                    // 查询角色详情
		//privateGroup.POST("/roles", ginx.WrapJsonBody[types.CreateRoleReq](r.CreateRoleCtl))  // 创建角色
		//privateGroup.DELETE("/roles/:id", ginx.WrapResponse(r.DeleteRoleCtl))                 // 删除角色
		//privateGroup.PUT("/roles", ginx.WrapJsonBody[types.UpdateRoleReq](r.UpdateRoleCtl))   //	更新角色
		privateGroup.POST("/notes", ginx.WrapJsonBodyAndClaims[types.CreateNoteForm, jwt.UserClaims](n.CreateNoteCtl)) // 发布文章
		privateGroup.PUT("/notes/:uuid/pass", ginx.WrapResponse(n.PassNoteCtl))
		privateGroup.GET("/notes", ginx.WrapQueryBody[types.QueryNoteForm](n.NoteListCtl))              // 获取文章列表
		privateGroup.GET("/notes/detail/:uuid", ginx.WrapResponse(n.NoteDetail))                        // 获取文章详情信息
		privateGroup.GET("/feed/notes", ginx.WrapQueryBody[types.FeedNoteQueryForm](n.FeedNoteListCtl)) // 获取推荐文章列表  后续要改成feed流模式

		// 点赞收藏相关
		privateGroup.POST("/interactive/like", ginx.WrapJsonBodyAndClaims[types.InteractiveForm, jwt.UserClaims](intrCtl.LikeCtl))
		//group.POST("/collect", ginx.WrapJsonBody[types.CollectForm](s.CollectCtl))
		//group.POST("/follow", ginx.WrapJsonBody[types.FollowForm](s.FollowCtl))

		// 评论相关
		privateGroup.POST("/comments", ginx.WrapJsonBodyAndClaims[types.AddCommentForm, jwt.UserClaims](commentCtl.AddCommentCtl))
		privateGroup.GET("/comments/:resource_id", ginx.WrapResponse(commentCtl.CommentListCtl))
	}

	return engine
}
func InitFileController(app *internal.App) *controller.FileController {
	return controller.NewFileController(app.Cfg)
}

func InitUserController(app *internal.App) *controller.UserController {
	userDao := dao.NewUserDao(app.DB)
	userRepo := repo.NewUserRepoInterface(userDao)

	codeCache := cache.NewRedisVerificationCodeCache(app.Cli)
	jwtHdl := jwt.NewRedisJWTHandler(app.Cli, app.Cfg)
	userSvc := service.NewUserSvc(userRepo, codeCache, jwtHdl)
	return controller.NewUserController(userSvc)
}

func InitPubController(app *internal.App) *controller.PublicController {
	userDao := dao.NewUserDao(app.DB)
	userRepo := repo.NewUserRepoInterface(userDao)
	codeCache := cache.NewRedisVerificationCodeCache(app.Cli)
	jwtHdl := jwt.NewRedisJWTHandler(app.Cli, app.Cfg)
	userSvc := service.NewUserSvc(userRepo, codeCache, jwtHdl)
	redisVerificationCodeCache := cache.NewRedisVerificationCodeCache(app.Cli)
	emailSvc := service.NewSMTPService(redisVerificationCodeCache)
	return controller.NewPublicController(userSvc, emailSvc)
}

//func InitRoleController(app *internal.App) controller. {
//	perDao := dao.NewPermissionDAO(app.DB)
//	roleDao := dao.NewRoleDAO(app.DB)
//	cac := cache.NewRoleCache(app.Cli)
//	roleRepo := repo.NewRoleRepo(perDao, roleDao, cac)
//	roleSvc := service.NewRoleService(roleRepo)
//	return controller.NewRoleController(roleSvc)
//}

func InitNoteController(app *internal.App) *controller.NoteCtl {
	noteDao := dao.NewNoteDao(app.DB)
	userDao := dao.NewUserDao(app.DB)
	imgDao := dao.NewImageDao(app.DB)
	noteRepo := repo.NewNoteRepo(noteDao, userDao, imgDao)
	interactiveCache := cache.NewInteractiveCache(app.Cli)
	interactiveDao := dao.NewInteractiveDao(app.DB)
	intrRepo := repo.NewInteractiveRepo(interactiveCache, interactiveDao)
	svc := service.NewNoteSvcImpl(noteRepo, app.Lg, intrRepo)
	return controller.NewNoteCtl(svc, app.Cfg)
}

func SetupCommentController(app *internal.App) *controller.CommentCtl {
	commentDao := dao.NewCommentDAO(app.DB)
	userDao := dao.NewUserDao(app.DB)
	userRepo := repo.NewUserRepoInterface(userDao)
	commentRepo := repo.NewCommentRepo(commentDao)
	commentSvc := service.NewCommentSvc(commentRepo, userRepo)
	return controller.NewCommentCtl(commentSvc)
}

func SetupInteracticeController(app *internal.App) *controller.InteractiveCtl {
	userDao := dao.NewUserDao(app.DB)
	userRepo := repo.NewUserRepoInterface(userDao)
	interactiveCache := cache.NewInteractiveCache(app.Cli)
	interactiveDao := dao.NewInteractiveDao(app.DB)
	interactiveRepo := repo.NewInteractiveRepo(interactiveCache, interactiveDao)
	notifyDao := dao.NewNotificationDao(app.DB)
	notifyRepo := repo.NewNotifyRepo(notifyDao)
	interactiveSvc := service.NewInteractiveSvc(interactiveRepo, notifyRepo, userRepo)
	return controller.NewInteractiveController(interactiveSvc)
}
