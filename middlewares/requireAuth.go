package middlewares

import (
	"fmt"
	"github/loa212/go-todo/initializers"
	"github/loa212/go-todo/models"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//get cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		println(err)
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//find user with token sub

		var user models.User
		initializers.DB.First(&user, claims["sub"])
		
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		
		//attach user to req

		c.Set("user", user)
		
		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


}