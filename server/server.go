package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
)

// Server struct
type Server struct {
	*echo.Echo
}

// New Initialize the controllers and routers
func New(database *gorm.DB, logger *lumberjack.Logger) (*Server, error) {
	server := Server{echo.New()}

	// Middleware
	server.HTTPErrorHandler = func(err error, context echo.Context) {
		message := err.Error()
		statusCode := context.Response().Status
		context.JSON(statusCode, map[string]map[string]interface{}{ // sub level mapping
			"error": {
				"message": message,
			},
		})
	}

	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	// server.Use(middleware.CORS())
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logger})) // Server header

	// Setup Controller
	InitializeControllers(database)

	// Setup Routers
	InitializeRouters(server)

	return &server, nil
}
