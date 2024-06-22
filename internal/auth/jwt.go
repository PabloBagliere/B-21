package auth

import (
	"encoding/hex"
	"time"

	"github.com/PabloBagliere/B-21/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/sha3"
)

type jwtClaims struct {
	Username    string                 `json:"username"`
	Role        string                 `json:"role"`
	Email       string                 `json:"email"`
	Permissions []string               `json:"permissions"`
	Data        map[string]interface{} `json:"data"`
	Iss         string                 `json:"iss"`
	Sub         string                 `json:"sub"`
	Aud         string                 `json:"aud"`
	Exp         int64                  `json:"exp"`
	Nbf         int64                  `json:"nbf"`
	Iat         int64                  `json:"iat"`
	Jti         string                 `json:"jti"`
}

type response struct {
	AccessToken  string  `json:"AccessToken"`
	ExpiresIn    float64 `json:"ExpiresIn"`
	RefreshToken string  `json:"RefreshToken"`
	TokenType    string  `json:"TokenType"`
}

type configAuth struct {
	secrectKey      string
	duration        time.Duration
	refreshDuration time.Duration
}

var secrectKey string
var duration time.Duration
var refreshDuration time.Duration

// Implementación de la interfaz jwt.Claims
func (c jwtClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Exp, 0)), nil
}

func (c jwtClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Iat, 0)), nil
}

func (c jwtClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(time.Unix(c.Nbf, 0)), nil
}

func (c jwtClaims) GetIssuer() (string, error) {
	return c.Iss, nil
}

func (c jwtClaims) GetSubject() (string, error) {
	return c.Sub, nil
}

func (c jwtClaims) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{c.Aud}, nil
}

func createToken() (string, error) {
	config, err := getConfigAuth()
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		Username:    "pablo",
		Role:        "admin",
		Email:       "test@test.com",
		Permissions: []string{"read", "write"},
		Data:        map[string]interface{}{"key": "value"},
		Iss:         "auth",
		Sub:         "auth",
		Aud:         "auth",
		Exp:         time.Now().Add(config.duration).Unix(),
		Nbf:         time.Now().Unix(),
		// TODO: Chequear como validar el tiempo de creación del token
		Iat: time.Now().Unix(),
		Jti: "123",
	})
	tokenString, err := t.SignedString([]byte(config.secrectKey))
	if err != nil {
		return "", errors.NewJWTError("Error creating token", err)
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	config, err := getConfigAuth()
	if err != nil {
		return false, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.secrectKey), nil
	})
	if err != nil {
		return false, errors.NewJWTError("Error validate token", err)
	}
	if _, ok := token.Claims.(*jwtClaims); !ok || !token.Valid {
		return false, nil
	}
	return true, nil
}

func refreshToken(tokenString string) (string, error) {
	config, err := getConfigAuth()
	if err != nil {
		return "", err
	}
	hash := sha3.New512()
	_, err = hash.Write([]byte(tokenString + config.secrectKey))
	if err != nil {
		return "", errors.NewJWTError("Error creating refresh token", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil

}

func CreateResponse() (*response, error) {
	config, err := getConfigAuth()
	if err != nil {
		return nil, err
	}
	token, err := createToken()
	if err != nil {
		return nil, err
	}
	refreshToken, err := refreshToken(token)
	if err != nil {
		return nil, err
	}
	return &response{
		AccessToken:  token,
		ExpiresIn:    config.refreshDuration.Seconds(),
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}

func InitJwt(config map[string]interface{}) (bool, error) {
	p, err := parseConfig(config)
	if err != nil {
		return false, err
	}
	secrectKey = p.secrectKey
	duration = p.duration
	refreshDuration = p.refreshDuration
	return true, nil
}

func parseConfig(p map[string]interface{}) (configAuth, error) {
	secret, ok := p["secret"].(string)
	if !ok {
		return configAuth{}, errors.NewJWTError("Error: secret not found in the configuration", nil)
	}
	tokenDurationStr, ok := p["tokenDuration"].(string)
	if !ok {
		return configAuth{}, errors.NewJWTError("Error: tokenDuration not found in the configuration", nil)
	}
	refreshTokenDurationStr, ok := p["refreshTokenDuration"].(string)
	if !ok {
		return configAuth{}, errors.NewJWTError("Error: refreshTokenDuration not found in the configuration", nil)
	}
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		return configAuth{}, errors.NewJWTError("Error: tokenDuration is not possible to parse time", err)
	}
	refreshTokenDuration, err := time.ParseDuration(refreshTokenDurationStr)
	if err != nil {
		return configAuth{}, errors.NewJWTError("Error: refreshTokenDuration is not possible to parse time", err)
	}
	return configAuth{
		secrectKey:      secret,
		duration:        tokenDuration,
		refreshDuration: refreshTokenDuration,
	}, nil
}

func getConfigAuth() (configAuth, error) {
	// check si secrectKey está en la configuración
	if secrectKey == "" {
		return configAuth{}, errors.NewJWTError("Error: secrectKey not init value", nil)
	}
	// check si duration está en la configuración
	if duration == 0 {
		return configAuth{}, errors.NewJWTError("Error: duration not init value", nil)
	}
	// check si refreshDuration está en la configuración
	if refreshDuration == 0 {
		return configAuth{}, errors.NewJWTError("Error: refreshDuration not init value", nil)
	}
	return configAuth{
		secrectKey:      secrectKey,
		duration:        duration,
		refreshDuration: refreshDuration,
	}, nil

}
