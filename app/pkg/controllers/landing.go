package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type LandingController struct {
	CookieStore *sessions.CookieStore
}

// Handler for our home page.
func (controller *LandingController) ShowLanding(w http.ResponseWriter, r *http.Request) {
	// session, _ := controller.CookieStore.Get(r, "session-name")

	// profile, exists := session.Values["profile"].(auth.Profile)

	// if exists {
	// 	templ.Handler(layout.Base("Welcome to ALLMINDS", template.Landing(profile))).ServeHTTP(w, r)
	// } else {
	// 	templ.Handler(layout.Base("Welcome to ALLMINDS", template.Landing(profile))).ServeHTTP(w, r)
	// }
}
