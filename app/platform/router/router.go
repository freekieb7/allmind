package router

import (
	"encoding/gob"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"viavia.io/platform"
	"viavia.io/platform/authenticator"
	"viavia.io/platform/middleware"
	"viavia.io/web/app/callback"
	"viavia.io/web/app/home"
	"viavia.io/web/app/landing"
	"viavia.io/web/app/login"
	"viavia.io/web/app/logout"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(platform.Profile{})

	// Create authenticator
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")

	// Public pages
	router.GET("/", landing.Handler)

	router.GET("/api/health", func(ctx *gin.Context) { ctx.Status(200) })

	// Authentication pages
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)

	router.GET("/home", middleware.IsAuthenticated, home.Handler)

	return router
}
