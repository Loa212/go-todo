package controllers

import (
	"github/loa212/go-todo/initializers"
	"github/loa212/go-todo/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//get email&psw from req.body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return 
	}

	//hash psw

	hashPsw, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error handling password",
		})
		return 
	}

	//create user

	user := models.User{
		Email: body.Email,
		Password: string(hashPsw),
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating the user",
			"details": result.Error.Error(),
		})
		return 
	}

	//return res
	c.JSON(http.StatusOK, gin.H{})
}

func Signin(c *gin.Context) {
	//get email&psw from req.body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return 
	}

	// find requested user in DB
	var user models.User
	// initializers.DB.First("email = ?", body.Email)
	initializers.DB.Where("email = ?", body.Email).First(&user)


	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid credentials",
			"message": "no user found with email " + body.Email,
		})
		return 
	}

	//compare sent psw with saved psw hash

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
			"message": "psw",
		})
		return 
	}	

	//generate JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour *24 *30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create token",
		})
		return 
	}	

	//respond with the generated JWT

	// c.JSON(http.StatusOK, gin.H{
	// 	"token":tokenString,
	// })

	//instead of simply returning JWT let's send it in a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}


func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting user with id=" + id,
		})
		return
	}

	result = initializers.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting user with id=" + id,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted",
	})
}


func Validate(c *gin.Context) {
	user, _ := c.Get("user")


	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}