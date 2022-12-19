// go:build wireinject
//  +build wireinject

package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	lg "github.com/sirupsen/logrus"

	"github.com/bayusamudra5502/go-backend-template/config"
	cc "github.com/bayusamudra5502/go-backend-template/controller"
	acc "github.com/bayusamudra5502/go-backend-template/controller/app"
	log "github.com/bayusamudra5502/go-backend-template/lib/log"
	"github.com/bayusamudra5502/go-backend-template/lib/log/logrus"
	"github.com/bayusamudra5502/go-backend-template/middleware"
	mc "github.com/bayusamudra5502/go-backend-template/middleware/chi"
	rc "github.com/bayusamudra5502/go-backend-template/routes/chi"
	rcApp "github.com/bayusamudra5502/go-backend-template/routes/chi/app"
	rcBase "github.com/bayusamudra5502/go-backend-template/routes/chi/interface"
	serviceApp "github.com/bayusamudra5502/go-backend-template/service/app"
)


func CreateHandler(logger log.Log) (http.Handler, error) {
	wire.Build(
		// Middleware
		wire.NewSet(
			mc.NewChiMiddlware,
			wire.Bind(new(middleware.Middleware), new(*mc.ChiMiddleware)),
		),

		// Repositories
		// wire.NewSet(),
		
		// Services
		wire.NewSet(
			// App Service
			wire.NewSet(
				wire.Struct(new(serviceApp.ServiceAppV1), "*"),
				wire.Bind(new(serviceApp.AppService), new(*serviceApp.ServiceAppV1)),
			),
		),
		
		// Controllers
		wire.NewSet(
			// App Controller
			wire.NewSet(
				wire.Struct(new(acc.AppControllerV1), "*"),
				wire.Bind(new(acc.AppController), new(*acc.AppControllerV1)),
			),

			// Controller Collections
			wire.Struct(new(cc.ControllerImpl), "*"),
			wire.Bind(new(cc.Controller), new(*cc.ControllerImpl)),
		),
		
		// Routes
		wire.NewSet(
			// App Routes
			wire.Struct(new(rcApp.AppRouteV1), "*"),

			// Routes Collections
			wire.Struct(new(rc.RoutesImplement), "*"),
			wire.Bind(new(rcBase.BaseChiRoute), new(rc.RoutesImplement)),
		),

		// Router
		wire.NewSet(
			NewChi,
			wire.Bind(new(http.Handler), new(*chi.Mux)),
		),

	)

	return nil, nil
}

func CreateServer(cfg *config.Config) (*App, error) {
	wire.Build(
		// Config Fields
		wire.FieldsOf(
			new(*config.Config), 
			"LogtailToken", 
			"ProductionMode",
		),

		// Logger
		wire.NewSet(
			logrus.NewLogtailHooks,
			logrus.NewFormatter,
			logrus.New,
			wire.Bind(new(lg.Formatter), new(*logrus.LogrusFormatter)),
			wire.Bind(new(log.Log), new(*logrus.LogrusLog)),
		),

		wire.NewSet(
			CreateHandler,
		),

		// Application
		New,
	)
	return nil, nil
}