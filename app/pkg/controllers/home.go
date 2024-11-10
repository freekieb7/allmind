package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type HomeController struct {
	CookieStore *sessions.CookieStore
}

const name = "allmind-app"

var (
	tracer  = otel.Tracer(name)
	meter   = otel.Meter(name)
	logger  = otelslog.NewLogger(name)
	rollCnt metric.Int64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter("dice.rolls",
		metric.WithDescription("The number of rolls by roll value"),
		metric.WithUnit("{roll}"))
	if err != nil {
		panic(err)
	}
}

// Handler for our home page.
func (controller *HomeController) ShowHome(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "roll")
	defer span.End()

	logger.InfoContext(ctx, "home-to-ip", "result", r.RemoteAddr)

	// session, _ := controller.CookieStore.Get(r, "session-name")
	// profile := session.Values["profile"].(auth.Profile)
	// templ.Handler(layout.Base("Home", template.Home(profile))).ServeHTTP(w, r)
}
