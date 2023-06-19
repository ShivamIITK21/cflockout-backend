package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ShivamIITK21/cflockout-backend/db"
	"github.com/ShivamIITK21/cflockout-backend/models"
)

type SignedDetails struct{
	Username string `json:"username" gorm:"PRIMARY_KEY;NOT NULL"`
	CFid     string `json:"cfid" gorm:"NOT NULL"`
	UserType string `json:"userType" gorm:"NOT NULL"`
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateTokens(username string, cfid string, user_type string) (token string, ref_token string, err error){
	claims := &SignedDetails{
		Username: username,
		CFid: cfid,
		UserType: user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(720)).Unix(),
		},
	}

	refresh_claims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(2000)).Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	ref_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims).SignedString([]byte(SECRET_KEY))

	return token, ref_token, err
}

func UpdateAllTokens(token string, username string) error{
	var user models.User
	user.Username = &username
	result := db.DB.Model(&user).Select("token").Updates(map[string]interface{}{"token":token})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedtoken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = "Invalid Token"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Expired Token"
		return
	}

	return claims, msg
}