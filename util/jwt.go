package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yangioc/bk_pack/dto"
)

func GenerateToken(serverHash string, data map[string]interface{}, expTime time.Time) (string, error) {
	dataMap := jwt.MapClaims{
		// "jti": "",                     // token 編號
		// "iss": "",                     // token 發行者
		"exp": jwt.NewNumericDate(expTime), //  token 過期時間
		// "iat": time.Now().UnixMilli(), // token 發行時間
		// "nbf": time.Now().UnixMilli(), // token 可使啟用時間
	}

	for k, v := range data {
		dataMap[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataMap)
	tokenString, err := token.SignedString([]byte(serverHash))
	return tokenString, err
}

func ParseToken(serverHash string, tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(serverHash), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}

// GenerateToken 生成 JWT 的函式
func GenerateTokenByStruct(name string, expTime time.Duration, data interface{}, secretKey []byte) (string, error) {
	// 設定聲明
	claims := dto.JwtStruct{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expTime)), // 設定 24 小時後過期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    name, // 你應用程式的名稱
		},
	}

	// 創建一個使用 HS256 加密演算法的 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密鑰簽名並生成最終的 JWT 字串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 驗證和解碼 JWT 的函式
func ParseTokenByStruct(tokenString string, data interface{}, secretKey []byte) (*dto.JwtStruct, error) {
	// 解析並驗證 JWT

	clia := dto.JwtStruct{
		Claims: data,
	}

	token, err := jwt.ParseWithClaims(tokenString, &clia, func(token *jwt.Token) (interface{}, error) {
		// 驗證簽名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid encode: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 確認 token 是否有效並提取聲明
	if claims, ok := token.Claims.(*dto.JwtStruct); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("inverind token")
	}
}
