package server

import (
	"fmt"
	"food-comming-api/config"
	"food-comming-api/user/handlers"
	"food-comming-api/user/repositories"
	"food-comming-api/user/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) InitUserHttpHandler() {
	userPostgresRepository := repositories.NewUserPostgresRepository(s.db)
	userFCMessagingRepository := repositories.NewUserFCMessaging()

	userUseCase := usecases.NewUserUseCaseImpl(userPostgresRepository, userFCMessagingRepository)

	userHttpHandler := handlers.NewUserHttpHandler(userUseCase)

	userRouters := s.app.Group("v1/user")
	userRouters.POST("", userHttpHandler.DetectUser)

}
