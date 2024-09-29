package home

import (
	"github.com/a-h/templ"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"viavia.io/platform"
	"viavia.io/web/view"
	"viavia.io/web/view/layout"
)

// Handler for our home page.
func Handler(ctx *gin.Context) {
	profile := sessions.Default(ctx).Get("profile").(platform.Profile)
	templ.Handler(layout.Base("Home", view.Home(profile))).ServeHTTP(ctx.Writer, ctx.Request)
}
