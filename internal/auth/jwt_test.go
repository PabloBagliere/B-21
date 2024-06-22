package auth_test

import (
	"strings"
	"testing"
	"time"

	"github.com/PabloBagliere/B-21/internal/auth"
)

// TestInitJwtValid tests the InitJwt function with valid configuration
func TestInitJwtValid(t *testing.T) {
	config := map[string]interface{}{
		"secret":               "mySecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	}
	result, err := auth.InitJwt(config)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result {
		t.Fatalf("Expected true, got %v", result)
	}
}

func TestInitJwtMissingSecret(t *testing.T) {
	config := map[string]interface{}{
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	}
	_, err := auth.InitJwt(config)
	if err == nil || !strings.Contains(err.Error(), "secret not found") {
		t.Errorf("Expected error about missing secret, got %v", err)
	}
}

func TestInitJwtMissingTokenDuration(t *testing.T) {
	config := map[string]interface{}{
		"secret":               "mySecret",
		"refreshTokenDuration": "24h",
	}
	_, err := auth.InitJwt(config)
	if err == nil || !strings.Contains(err.Error(), "tokenDuration not found") {
		t.Errorf("Expected error about missing tokenDuration, got %v", err)
	}
}

func TestInitJwtInvalidTokenDuration(t *testing.T) {
	config := map[string]interface{}{
		"secret":               "mySecret",
		"tokenDuration":        "invalid",
		"refreshTokenDuration": "24h",
	}
	_, err := auth.InitJwt(config)
	if err == nil || !strings.Contains(err.Error(), "tokenDuration is not possible to parse") {
		t.Errorf("Expected error about invalid tokenDuration, got %v", err)
	}
}

func TestInitJwtMissingRefreshTokenDuration(t *testing.T) {
	config := map[string]interface{}{
		"secret":        "mySecret",
		"tokenDuration": "2h",
	}
	_, err := auth.InitJwt(config)
	if err == nil || !strings.Contains(err.Error(), "refreshTokenDuration not found") {
		t.Errorf("Expected error about missing refreshTokenDuration, got %v", err)
	}
}

func TestInitJwtInvalidRefreshTokenDuration(t *testing.T) {
	config := map[string]interface{}{
		"secret":               "mySecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "invalid",
	}
	_, err := auth.InitJwt(config)
	if err == nil || !strings.Contains(err.Error(), "refreshTokenDuration is not possible to parse") {
		t.Errorf("Expected error about invalid refreshTokenDuration, got %v", err)
	}
}

// TestCreateResponseSuccess tests the CreateResponse function with valid configuration
func TestCreateResponseSuccess(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "validSecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	})
	refreshTokenDuration, _ := time.ParseDuration("24h")
	resp, err := auth.CreateResponse()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Fatal("Expected response to be non-nil")
	}
	if resp.AccessToken == "" || resp.RefreshToken == "" {
		t.Error("Expected AccessToken and RefreshToken to be non-empty")
	}
	if resp.TokenType != "Bearer" {
		t.Errorf("Expected TokenType to be 'Bearer', got %v", resp.TokenType)
	}
	if resp.ExpiresIn <= 0 {
		t.Errorf("Expected ExpiresIn to be positive, got %v", resp.ExpiresIn)
	}
	// check expiresIn is less than or equal to refresh duration
	if resp.ExpiresIn != refreshTokenDuration.Seconds() {
		t.Errorf("Expected ExpiresIn to be %v, got %v", refreshTokenDuration.Seconds(), resp.ExpiresIn)
	}
}

func TestCreateResponseFailInitJwt(t *testing.T) {
	// Resetting JWT configuration to simulate uninitialized state
	auth.InitJwt(map[string]interface{}{})
	resp, err := auth.CreateResponse()
	if err == nil {
		t.Logf("Response: %v", resp)
		t.Fatal("Expected error, got nil")
	}
	if resp != nil {
		t.Errorf("Expected response to be nil, got %v", resp)
	}

}

func TestCreateResponseWithInvalidSecret(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	})
	_, err := auth.CreateResponse()
	if err == nil {
		t.Fatal("Expected error due to invalid secret, got nil")
	}
}

func TestCreateResponseWithInvalidDuration(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "validSecret",
		"tokenDuration":        "invalid",
		"refreshTokenDuration": "24h",
	})
	_, err := auth.CreateResponse()
	if err == nil {
		t.Fatal("Expected error due to invalid tokenDuration, got nil")
	}
}

// TestVerifyTokenValid tests the VerifyToken function with a valid token

func TestValidateTokenValid(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "validSecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	})
	resp, _ := auth.CreateResponse()
	valid, err := auth.ValidateToken(resp.AccessToken)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !valid {
		t.Error("Expected token to be valid")
	}
}

func TestValidateTokenInvalid(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "validSecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	})
	invalidToken := "invalidTokenString"
	valid, err := auth.ValidateToken(invalidToken)
	if valid || err == nil {
		t.Errorf("Expected token to be invalid, got valid: %v, error: %v", valid, err)
	}
}

func TestValidateTokenExpired(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "validSecret",
		"tokenDuration":        "1ms",
		"refreshTokenDuration": "24h",
	})
	resp, _ := auth.CreateResponse()
	time.Sleep(2 * time.Millisecond) // Ensure token is expired
	valid, err := auth.ValidateToken(resp.AccessToken)
	if valid || err == nil {
		t.Errorf("Expected token to be invalid, got valid: %v, error: %v", valid, err)
	}
}

func TestValidateTokenBeforeNotBefore(t *testing.T) {
	// This test might require adjustments in the implementation to set a future `nbf`
}

func TestValidateTokenInvalidSecret(t *testing.T) {
	auth.InitJwt(map[string]interface{}{
		"secret":               "validSecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	})
	resp, _ := auth.CreateResponse()
	// Re-initialize with a different secret
	auth.InitJwt(map[string]interface{}{
		"secret":               "differentSecret",
		"tokenDuration":        "2h",
		"refreshTokenDuration": "24h",
	})
	valid, err := auth.ValidateToken(resp.AccessToken)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if valid {
		t.Error("Expected token to be invalid")
	}
}

func TestValidateTokenUninitializedConfig(t *testing.T) {
	// Reset or clear the JWT configuration
	auth.InitJwt(map[string]interface{}{}) // Assuming this clears the configuration
	token := "anyToken"
	valid, err := auth.ValidateToken(token)
	if valid || err == nil {
		t.Errorf("Expected token validation to fail due to uninitialized config, got valid: %v, error: %v", valid, err)
	}
}
