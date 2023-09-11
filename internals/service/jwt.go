package service



import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v4"
)


type JwtOptionService interface{
	GetService() JwtService
	SwitchService(kind string) error
}

type JwtOptions struct{
	Symmetric SymJwt
	Asymmetric AsymJwt
	option JwtService
}


func (joptions *JwtOptions) GetService() JwtService {
	return joptions.option
}

func (joptions *JwtOptions) SwitchService(kind string) error {
	// kind can only be of two types `symmetric` or `asymmtric`

	if kind == "symmetric" {
		joptions.option = &joptions.Symmetric
	} else if kind == "asymmtric" {
		joptions.option = &joptions.Asymmetric
	} else {
		return fmt.Errorf("jwt kind does not exist")
	}

	return nil
}

type JwtService interface {
	GenerateToken(id string) string
	DecodeToken(tokenStr string) (string, error)
	ValidationToken(claim Claims) bool
	GenerateRefreshToken(id string) string
	GenerateTokenPair(id string) map[string]interface{}
}




type SymJwt struct {
	accessTokenLife time.Duration
	refreshTokenLife time.Duration
	iss string
	alg jwt.SigningMethod
	secret string
}

func NewSymJwt(accessTokenLife int64, refreshTokenLife int64, iss string, alg string, secret string) SymJwt {

	return SymJwt{
		accessTokenLife: time.Duration(accessTokenLife),
		refreshTokenLife: time.Duration(refreshTokenLife),
		iss: iss,
		alg: jwt.SigningMethodHS256,
		secret: secret,
	}
}

func (symjwt *SymJwt) GenerateToken(id string) string {

	expirationTime := time.Now().Add(symjwt.accessTokenLife * time.Minute)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ID: id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer: symjwt.iss,
		},
	}


	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(symjwt.alg, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(symjwt.secret)

	return tokenString
}


func (symjwt *SymJwt) DecodeToken(tokenStr string) (string, error) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {

		// validate algorithm used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return token.Claims, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		return "", err
		// w.WriteHeader(http.StatusBadRequest)
	}
	if !tkn.Valid {
		return "", err
		// w.WriteHeader(http.StatusUnauthorized)
	}

	if !symjwt.ValidationToken(*claims) {
		return "", jwt.NewValidationError("invalid token", 4)
	}

	return claims.ID, err
}


func (symjwt *SymJwt) ValidationToken(claim Claims) bool {
	return true
}


func (symjwt *SymJwt) GenerateRefreshToken(id string) string {

	expirationTime := time.Now().Add(symjwt.refreshTokenLife * time.Minute)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ID: id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer: symjwt.iss,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(symjwt.alg, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(symjwt.secret)

	return tokenString
}


func (symjwt *SymJwt) GenerateTokenPair(id string) map[string]interface{} {
	return map[string]interface{}{
		"access_token": symjwt.GenerateToken(id),
		"refresh_token": symjwt.GenerateRefreshToken(id),
	}
}


type AsymJwt struct {
	privateKey string
	publicKey string
	accessTokenLife time.Duration
	refreshTokenLife time.Duration
	iss string
	alg jwt.SigningMethod
}


func NewAsymJwt(accessTokenLife int64, refreshTokenLife int64, iss string,
	alg string, privateKey string, publicString string) AsymJwt {

	return AsymJwt{
		privateKey: privateKey,
		publicKey: publicString,
		accessTokenLife: time.Duration(accessTokenLife),
		refreshTokenLife: time.Duration(refreshTokenLife),
		iss: iss,
		alg: jwt.SigningMethodHS256,
	}
}





func (asymjwt *AsymJwt) GenerateToken(id string) string {

	expirationTime := time.Now().Add(asymjwt.accessTokenLife * time.Minute)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ID: id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer: asymjwt.iss,
		},
	}


	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(asymjwt.alg, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(asymjwt.privateKey)

	return tokenString
}


func (asymjwt *AsymJwt) DecodeToken(tokenStr string) (string, error) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {

		// validate algorithm used
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return token.Claims, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// w.WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		return "", err
		// w.WriteHeader(http.StatusBadRequest)
	}
	if !tkn.Valid {
		return "", err
		// w.WriteHeader(http.StatusUnauthorized)
	}

	if !asymjwt.ValidationToken(*claims) {
		return "", jwt.NewValidationError("invalid token", 4)
	}

	return claims.ID, err
}

func (asymjwt *AsymJwt) ValidationToken(claim Claims) bool {
	return true
}

func (asymjwt *AsymJwt) GenerateRefreshToken(id string) string {

	expirationTime := time.Now().Add(asymjwt.refreshTokenLife * time.Minute)
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ID: id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer: asymjwt.iss,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(asymjwt.alg, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(asymjwt.privateKey)

	return tokenString
}

func (asymjwt *AsymJwt) GenerateTokenPair(id string) map[string]interface{} {
	return map[string]interface{}{
		"access_token": asymjwt.GenerateToken(id),
		"refresh_token": asymjwt.GenerateRefreshToken(id),
	}
}
