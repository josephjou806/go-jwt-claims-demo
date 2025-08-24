package tests

import (
	"testing"
	"time"

	"github.com/josephjou806/go-jwt-claims-demo/internal/token"
	"github.com/stretchr/testify/require"
)

func TestJWTGenerateAndValidate(t *testing.T) {
	tm := token.NewTokenManager("unit-test-secret")
	tok, err := tm.Generate("user-123", time.Minute)
	require.NoError(t, err)

	claims, err := tm.Validate(tok)
	require.NoError(t, err)
	require.Equal(t, "user-123", claims.Subject)
}
