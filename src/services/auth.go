package services

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Iss      string `json:"iss"`
	Aud      string `json:"aud"`
	AuthTime int64  `json:"auth_time"`
	UserId   string `json:"user_id"`
	Sub      string `json:"sub"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func checkFirebaseJWT(tokenString string) (CustomClaims, error) {
	resp, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		log.Fatalf("Failed to make a request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		log.Fatalf("Failed to json unmarshal: %v", err)
	}

	// JWTのヘッダを解析し署名に用いられている鍵を取得
	parts := strings.Split(tokenString, ".")

	// decode the header
	headerJson, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		log.Fatalf("Error decoding JWT header: %v", err)
		return CustomClaims{}, err
	}

	var header map[string]interface{}

	err = json.Unmarshal(headerJson, &header)
	if err != nil {
		log.Fatalf("Error unmarshalling JWT header: %v", err)
		return CustomClaims{}, err
	}

	kid := header["kid"].(string)
	certString := result[kid].(string)
	block, _ := pem.Decode([]byte(certString))
	if block == nil {
		log.Fatalf("failed to parse PEM block containing the public key")
		return CustomClaims{}, err
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatalf("failed to parse certificate: %v", err)
		return CustomClaims{}, err
	}

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	// 署名を検証
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return rsaPublicKey, nil
	})

	if err != nil {
		log.Fatalf("Error while parsing token: %v\n", err)
		return CustomClaims{}, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if time.Unix(claims.Exp, 0).Before(time.Now()) {
			return CustomClaims{}, errors.New("Token is valid. But token is expired.")
		} else {
			return *claims, nil
		}
	} else {
		return CustomClaims{}, errors.New("Token is not valid")
	}
}
