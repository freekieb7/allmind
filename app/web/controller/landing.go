package controller

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"viavia.io/platform"
	"viavia.io/web/view"
	"viavia.io/web/view/layout"
)

type LandingController struct {
	CookieStore *sessions.CookieStore
}

// Handler for our home page.
func (controller *LandingController) ShowLanding(w http.ResponseWriter, r *http.Request) {
	session, _ := controller.CookieStore.Get(r, "session-name")

	profile, exists := session.Values["profile"].(platform.Profile)

	if exists {
		templ.Handler(layout.Base("Welcome to ALLMINDS", view.Landing(profile))).ServeHTTP(w, r)
	} else {
		templ.Handler(layout.Base("Welcome to ALLMINDS", view.Landing(profile))).ServeHTTP(w, r)
	}
}
