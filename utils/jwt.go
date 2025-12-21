package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(userId int, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
			"userId": userId,
			"email": email,
			"exp": time.Now().Add(time.Hour * 2).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
		return "", err
    }

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
      return secretKey, nil
   })
  
   if err != nil {
      return nil, err
   }
  
   if !token.Valid {
      return nil, errors.New("Invalid token")
   }

   claims, ok := token.Claims.(jwt.MapClaims)
   if !ok {
      return nil, errors.New("Invalid token claims") 
   }

   return claims, nil
}