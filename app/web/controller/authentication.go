package controller

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	"viavia.io/platform"
	"viavia.io/platform/authenticator"
)

type AuthenticationController struct {
	CookieStore   *sessions.CookieStore
	Authenticator *authenticator.Authenticator
}

func (controller *AuthenticationController) Login(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	state := base64.StdEncoding.EncodeToString(b)

	session, err := controller.CookieStore.Get(r, "session-name")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Save the state inside the session.
	session.Values["state"] = state
	if err := session.Save(r, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Location", controller.Authenticator.AuthCodeURL(state))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (controller *AuthenticationController) Callback(w http.ResponseWriter, r *http.Request) {
	session, _ := controller.CookieStore.Get(r, "session-name")

	log.Println(session.Values["state"])

	if r.URL.Query().Get("state") != session.Values["state"] {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid state parameter."))
		return
	}

	// Exchange an authorization code for a token.
	token, err := controller.Authenticator.Exchange(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed to convert an authorization code into a token."))
		return
	}

	idToken, err := controller.Authenticator.VerifyIDToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to verify ID Token."))
		return
	}

	var profile platform.Profile
	if err := idToken.Claims(&profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	if err := session.Save(r, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Redirect to logged in page.
	w.Header().Add("Location", "/home")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (controller *AuthenticationController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := controller.CookieStore.Get(r, "session-name")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	returnTo, err := url.Parse("https://" + r.Host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	session.Values = make(map[interface{}]interface{}) // Clear
	session.Save(r, w)

	// Redirect
	w.Header().Add("Location", logoutUrl.String())
	w.WriteHeader(http.StatusTemporaryRedirect)
}
