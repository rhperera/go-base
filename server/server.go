package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rhperera/go-base/common"
	"github.com/rhperera/go-base/config"
)

var EchoCon *echo.Echo
var EchoRG *echo.Group

func Init() {
	EchoCon = echo.New()
	http.Handle("/", EchoCon)
}

/*
Calling this method will block the main thread
until further callbacks from the echo framework
*/
func Connect(port string) {
	if EchoCon == nil {
		log.Fatal(common.ErrorEchoServerInit)
		return
	}
	EchoCon.Logger.Fatal(EchoCon.Start(":" + port))
}

func InitAPI() {
	if EchoCon == nil {
		log.Fatal(common.ErrorEchoServerInit)
	}
	EchoRG = EchoCon.Group("/api")

	EchoRG.GET("", defaultGetOk)
}

func getCORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
}

func GetJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:       &JWTClaims{},
		SigningKey:   []byte(config.Get("JWT_SECRET")),
		TokenLookup:  common.JWTTokenLookup,
		ContextKey:   common.JWTContextKey,
		ErrorHandler: JWTErrorHandler,
	}
}

func JWTErrorHandler(e error) error {
	response := HttpResponse{}
	response.SetError(common.ErrorCodeInvalidJWT, common.ErrorInvalidJWT)
	return echo.NewHTTPError(http.StatusForbidden, response)
}

func defaultGetOk(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
