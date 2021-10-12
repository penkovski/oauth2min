package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

const (
	tokenPath      = "/token"
	introspectPath = "/token/introspect"
	authorizePath  = "/authorize"
	publicKeysPath = "/keys"
)

type OAuthService interface {
	Token(ctx context.Context, request *TokenRequest) (*TokenResponse, error)
}

type OAuthHandler struct {
	oauth  OAuthService
	logger zerolog.Logger
}

func New(oauth OAuthService, logger zerolog.Logger) http.Handler {
	handler := &OAuthHandler{
		oauth:  oauth,
		logger: logger,
	}

	mux := mux.NewRouter()
	mux.Methods("POST").Path(tokenPath).HandlerFunc(handler.Token)
	mux.Methods("POST").Path(introspectPath).HandlerFunc(handler.TokenIntrospect) // TODO: only POST or GET too?
	mux.Methods("POST").Path(authorizePath).HandlerFunc(handler.Authorize)
	mux.Methods("GET").Path(publicKeysPath).HandlerFunc(handler.Keys)

	return mux
}

func (h *OAuthHandler) Token(w http.ResponseWriter, r *http.Request) {
	req := new(TokenRequest)
	if err := req.Decode(r.Body); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	token, err := h.oauth.Token(r.Context(), req)
	if err != nil {

	}

	if err := json.NewEncoder(w).Encode(&token); err != nil {

	}
}

func (h *OAuthHandler) TokenIntrospect(w http.ResponseWriter, r *http.Request) {

}

func (h *OAuthHandler) Authorize(w http.ResponseWriter, r *http.Request) {

}

func (h *OAuthHandler) Keys(w http.ResponseWriter, r *http.Request) {

}
