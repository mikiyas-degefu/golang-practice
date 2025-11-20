package main

import (
	"fmt"
	"strings"

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

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			ctx.JSON(401, gin.H{"error": "Invalid authorization header"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "Invalid JWT"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("user", claims)
		}

		ctx.Next()

	}
}

func main() {
	router := gin.Default()
	protected := router.Group("/api")
	protected.Use(AuthMiddleware())

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

	protected.GET("/profile", func(ctx *gin.Context) {
		userClaims := ctx.MustGet(("user"))

		ctx.JSON(200, gin.H{
			"message": "Protected route!",
			"user":    userClaims,
		})

	})

	router.Run()
}
