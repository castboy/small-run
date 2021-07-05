package main

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"small-run/fx_gin/internal/controllers"
	"small-run/fx_gin/internal/models"
	"small-run/fx_gin/shared/libraries"
	"small-run/fx_gin/shared/middlewares"
)

func main() {
	fx.New(Libraries, Models, Controllers, Middlewares, RegisterMiddlewares, RegisterRouters, App).Run()
}

var Libraries = fx.Provide(
	libraries.NewLogger,
	libraries.NewConfig,
	libraries.NewDB,
	libraries.NewEngine,
)

var Models = fx.Provide(
	models.NewStudentModel,
	models.NewSubjectModel,
	models.NewScoreModel,
	models.NewAuthModel,
)

var Controllers = fx.Provide(
	controllers.NewStudentController,
	controllers.NewSubjectController,
	controllers.NewScoreController,
	controllers.NewAuthController,
)

var Middlewares = fx.Provide(
	middlewares.NewDealTime,
)

var RegisterMiddlewares = fx.Invoke(
	func(engine *gin.Engine, mds middlewares.MiddleWares) {
		middlewares.Register(engine, mds)
	},
)

var RegisterRouters = fx.Invoke(
	func(engine *gin.Engine, c controllers.Controllers) {
		group := engine.Group("/v1")
		controllers.Register(group, c)
	},
)

var App = fx.Invoke(
	func(engine *gin.Engine, lf fx.Lifecycle) {
		lf.Append(
			fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := engine.Run(); err != nil {
							panic(errors.WithStack(err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			},
		)
	},
)
