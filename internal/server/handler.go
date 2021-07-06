package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	tokenPath      = "/token"
	introspectPath = "/token/introspect"
	authorizePath  = "/authorize"
	publicKeysPath = "/keys"
)

type OAuthService interface{}

type OAuthHandler struct {
	oauth OAuthService
}

func New(oauth OAuthService) http.Handler {
	handler := &OAuthHandler{
		oauth: oauth,
	}

	mux := mux.NewRouter()
	mux.Methods("POST").Path(tokenPath).HandlerFunc(handler.Token)
	mux.Methods("POST").Path(authorizePath).HandlerFunc(handler.Authorize)
	mux.Methods("POST").Path(introspectPath).HandlerFunc(handler.TokenIntrospect) // TODO: only POST or GET too?
	mux.Methods("GET").Path(publicKeysPath).HandlerFunc(handler.Keys)

	return mux
}

func (h *OAuthHandler) Token(w http.ResponseWriter, r *http.Request) {

}

func (h *OAuthHandler) Authorize(w http.ResponseWriter, r *http.Request) {

}

func (h *OAuthHandler) TokenIntrospect(w http.ResponseWriter, r *http.Request) {

}

func (h *OAuthHandler) Keys(w http.ResponseWriter, r *http.Request) {

}
