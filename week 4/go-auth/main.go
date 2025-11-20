package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = make(map[string]*User)

var jwtSecret = []byte("mike")

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Go Authentication and Authorization tutorial!",
		})
	})

	router.POST("/register", func(ctx *gin.Context) {
		var user User

		// parse json
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid request payload",
			})
			return
		}

		// check required fields
		if user.Email == "" || user.Password == "" {
			ctx.JSON(400, gin.H{
				"error": "Email and password are required",
			})
			return
		}

		//check user already exist
		if _, exists := users[user.Email]; exists {
			ctx.JSON(400, gin.H{
				"error": "User already exists",
			})
			return
		}

		//hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		user.Password = string(hashedPassword)

		users[user.Email] = &user

		ctx.JSON(200, gin.H{"message": "User registered successfully"})

	})

	router.POST("/login", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid request payload",
			})

			return
		}

		//login logic
		existingUser, ok := users[user.Email]
		if !ok || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
			ctx.JSON(401, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": existingUser.ID,
			"email":   existingUser.Email,
		})

		jwtToken, err := token.SignedString(jwtSecret)

		if err != nil {
			ctx.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
	})

	router.Run()
}
