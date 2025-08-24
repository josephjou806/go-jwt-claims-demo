package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/josephjou806/go-jwt-claims-demo/internal/config"
	"github.com/josephjou806/go-jwt-claims-demo/internal/repository"
	"github.com/josephjou806/go-jwt-claims-demo/internal/server"
	"github.com/josephjou806/go-jwt-claims-demo/internal/services"
	"github.com/josephjou806/go-jwt-claims-demo/internal/token"
	"github.com/stretchr/testify/require"
)

func buildTestServer() http.Handler {
	cfg := config.Load()
	repo := repository.NewInMemoryClaimRepository()
	svc := services.NewClaimService(repo)
	tm := token.NewTokenManager("test-secret")
	return server.NewRouter(cfg, svc, tm)
}

func getToken(t *testing.T, h http.Handler) string {
	req := httptest.NewRequest("POST", "/login", strings.NewReader(`{"username":"u","password":"p"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	var resp struct {
		Token string `json:"token"`
	}
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	return resp.Token
}

func TestClaimsLookup_Authorized(t *testing.T) {
	h := buildTestServer()
	tok := getToken(t, h)

	req := httptest.NewRequest("GET", "/claims/1001", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	b, _ := io.ReadAll(rec.Body)
	require.Contains(t, string(b), "\"id\":\"1001\"")
}

func TestClaimsLookup_Unauthorized(t *testing.T) {
	h := buildTestServer()

	req := httptest.NewRequest("GET", "/claims/1001", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	require.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestToken_Expired(t *testing.T) {
	tm := token.NewTokenManager("short")
	tok, err := tm.Generate("u", time.Millisecond*10)
	require.NoError(t, err)
	time.Sleep(time.Millisecond * 20)
	_, err = tm.Validate(tok)
	require.Error(t, err)
}
