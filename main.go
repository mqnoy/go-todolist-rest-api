package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/mqnoy/go-todolist-rest-api/config"
	"github.com/mqnoy/go-todolist-rest-api/handler"
	"github.com/mqnoy/go-todolist-rest-api/middleware"
	"github.com/mqnoy/go-todolist-rest-api/model"
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
	"gorm.io/gorm"

	_taskHttpDelivery "github.com/mqnoy/go-todolist-rest-api/task/delivery/http"
	_taskRepoMySQL "github.com/mqnoy/go-todolist-rest-api/task/repository/mysql"
	_taskUseCase "github.com/mqnoy/go-todolist-rest-api/task/usecase"
	_userHttpDelivery "github.com/mqnoy/go-todolist-rest-api/user/delivery/http"
	_userRepoMySQL "github.com/mqnoy/go-todolist-rest-api/user/repository/mysql"
	_userUseCase "github.com/mqnoy/go-todolist-rest-api/user/usecase"
)

var (
	appCfg config.Configuration
)

type AppCtx struct {
	mysqlDB *gorm.DB
}

func init() {
	appCfg = config.AppConfig

	clogger.Setup(clogger.CloggerOpt{
		Level:       appCfg.LoggerLevel,
		Environment: appCfg.App.Environment,
	})
}

func main() {
	appCtx := AppCtx{
		mysqlDB: config.InitMySQLDatabase(appCfg),
	}

	// Auto migration
	if appCfg.MigrateConfig.AutoMigrate {
		errAutoMigrate := appCtx.mysqlDB.AutoMigrate(
			&model.User{},
			&model.Task{},
			&model.Member{},
			&model.MemberTask{},
		)

		if errAutoMigrate != nil {
			clogger.Logger().Error(errAutoMigrate)
		}
	}

	// The HTTP Server
	addr := appCfg.Server.Address()
	server := &http.Server{
		Addr:    addr,
		Handler: AppHandler(appCtx),
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	clogger.Logger().Debugf("server running on %s", addr)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				clogger.Logger().Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			clogger.Logger().Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		clogger.Logger().Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

}

func AppHandler(appctx AppCtx) http.Handler {
	mux := chi.NewRouter()

	// Setup middleware
	mux.Use(chiMiddleware.RealIP)
	mux.Use(middleware.NewStructuredLogger(clogger.Logger()))
	mux.Use(middleware.PanicRecoverer)

	// Initialize Repository
	userRepoMySQL := _userRepoMySQL.New(appctx.mysqlDB)
	taskRepoMySQL := _taskRepoMySQL.New(appctx.mysqlDB)

	// Initialize UseCase
	userUseCase := _userUseCase.New(userRepoMySQL)
	taskUseCase := _taskUseCase.New(taskRepoMySQL, userUseCase)

	// Initialize Middleware
	authMiddleware := middleware.NewAuthorizationMiddleware(userUseCase)

	// Fallback
	mux.NotFound(handler.FallbackHandler)

	// Initialize handler
	_userHttpDelivery.New(mux, userUseCase)
	_taskHttpDelivery.New(mux, authMiddleware, taskUseCase)

	// Print all routes
	chi.Walk(mux, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		clogger.Logger().Debugf("[%s]: '%s' has %d middlewares\n", method, route, len(middlewares))
		return nil
	})

	return mux
}
