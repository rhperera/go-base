package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/rhperera/go-base/common"
	"github.com/rhperera/go-base/domain"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

//should be called after JWT middleware
func (mw *Middleware) SetValidateUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		storeVal := c.Get(common.JWTContextKey)
		jwtConfig := storeVal.(*jwt.Token)
		claims := jwtConfig.Claims.(*JWTClaims)
		u := domain.User {
			Name: claims.Name,
			Admin: claims.Admin,
			Id: claims.Id,
			Email: claims.Email,
			Avatar: claims.Avatar,
			Mobile: claims.Mobile,
			Roles: claims.Roles,
		}
		c.Set(common.UserContext, u)
		return next(c)
	}
}

type JWTClaims struct {
	Name   string   `json:"name"`
	Admin  bool     `json:"admin"`
	Id     int      `json:"id"`
	Email  string   `json:"email"`
	Avatar string   `json:"avatar"`
	Mobile string   `json:"mobile"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}
