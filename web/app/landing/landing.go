package landing

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
	var profile platform.Profile
	if sesProfile := sessions.Default(ctx).Get("profile"); sesProfile != nil {
		profile = sesProfile.(platform.Profile)
		templ.Handler(layout.Base("Welcome to ALLMINDS", view.Landing(profile))).ServeHTTP(ctx.Writer, ctx.Request)
	} else {
		templ.Handler(layout.Base("Welcome to ALLMINDS", view.Landing(profile))).ServeHTTP(ctx.Writer, ctx.Request)
	}

}
