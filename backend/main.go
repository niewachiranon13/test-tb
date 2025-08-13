package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret-key")

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// decodeBase64 decodes a base64 string to plain text
func decodeBase64(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	// Example: postgres://user:password@localhost:5432/dbname?sslmode=disable
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=testtb port=5432 sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	app := fiber.New()

	app.Post("/api/register", func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}
		if req.Username == "" || req.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing fields"})
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := User{Username: req.Username, Password: string(hash)}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user exists"})
		}
		return c.JSON(fiber.Map{"message": "registered"})
	})

	app.Post("/api/login", func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
		}
		decodedPassword, err := decodeBase64(req.Password)
		if err != nil {
			fmt.Println("Base64 decode error:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid password encoding"})
		}
		fmt.Println("[DEBUG] Username:", req.Username)
		fmt.Println("[DEBUG] Decoded password:", decodedPassword)
		var user User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			fmt.Println("[DEBUG] User not found")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}
		fmt.Println("[DEBUG] User hash:", user.Password)
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(decodedPassword)) != nil {
			fmt.Println("[DEBUG] Password not match")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}
		fmt.Println("[DEBUG] Login success!")
		exp := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: req.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(exp),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(jwtKey)
		return c.JSON(fiber.Map{"token": tokenString, "username": req.Username})
	})

	app.Get("/api/validate", func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}
		token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}
		claims := token.Claims.(*Claims)
		return c.JSON(fiber.Map{"username": claims.Username})
	})

	app.Listen(":8080")
}
