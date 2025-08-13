package main

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestApp() (*fiber.App, *gorm.DB) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&User{})
	app := fiber.New()

	// Register endpoints (copy from main.go)
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
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid password encoding"})
		}
		var user User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(decodedPassword)) != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
		}
		return c.JSON(fiber.Map{"token": "dummy-token", "username": req.Username})
	})

	return app, db
}

func TestAuthCases(t *testing.T) {
	t.Log("Run: TestAuthCases")
	app, _ := setupTestApp()

	// Register with missing fields
	missingBody := `{"username":"","password":""}`
	reqMissing := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(missingBody))
	reqMissing.Header.Set("Content-Type", "application/json")
	respMissing, _ := app.Test(reqMissing)
	if respMissing.StatusCode == http.StatusOK {
		t.Error("register with missing fields should fail")
	}

	// Login with missing fields
	loginMissing := `{"username":"","password":""}`
	reqLoginMissing := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(loginMissing))
	reqLoginMissing.Header.Set("Content-Type", "application/json")
	respLoginMissing, _ := app.Test(reqLoginMissing)
	if respLoginMissing.StatusCode == http.StatusOK {
		t.Error("login with missing fields should fail")
	}

	// Login with invalid base64 password
	loginInvalidBase64 := `{"username":"testuser","password":"not_base64!"}`
	reqInvalidBase64 := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(loginInvalidBase64))
	reqInvalidBase64.Header.Set("Content-Type", "application/json")
	respInvalidBase64, _ := app.Test(reqInvalidBase64)
	if respInvalidBase64.StatusCode == http.StatusOK {
		t.Error("login with invalid base64 password should fail")
	}

	// Register user
	registerBody := `{"username":"testuser","password":"testpass"}`
	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(registerBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("register failed: %d", resp.StatusCode)
	}

	// Register duplicate user
	resp2, _ := app.Test(req)
	if resp2.StatusCode == http.StatusOK {
		t.Error("duplicate user should fail")
	}

	// Login with correct password
	loginBody := `{"username":"testuser","password":"` + base64.StdEncoding.EncodeToString([]byte("testpass")) + `"}`
	reqLogin := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(loginBody))
	reqLogin.Header.Set("Content-Type", "application/json")
	respLogin, _ := app.Test(reqLogin)
	if respLogin.StatusCode != http.StatusOK {
		t.Fatalf("login failed: %d", respLogin.StatusCode)
	}

	// Login with wrong password
	wrongLoginBody := `{"username":"testuser","password":"` + base64.StdEncoding.EncodeToString([]byte("wrongpass")) + `"}`
	reqWrong := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(wrongLoginBody))
	reqWrong.Header.Set("Content-Type", "application/json")
	respWrong, _ := app.Test(reqWrong)
	if respWrong.StatusCode == http.StatusOK {
		t.Error("login with wrong password should fail")
	}
}

func TestDecodeBase64(t *testing.T) {
	t.Log("Run: TestDecodeBase64")
	plain := "1"
	encoded := base64.StdEncoding.EncodeToString([]byte(plain))
	decoded, err := decodeBase64(encoded)
	if err != nil {
		t.Fatalf("decodeBase64 error: %v", err)
	}
	if decoded != plain {
		t.Errorf("expected %s, got %s", plain, decoded)
	}
}

func TestPasswordHashAndCompare(t *testing.T) {
	t.Log("Run: TestPasswordHashAndCompare")
	password := "mysecret"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("hash error: %v", err)
	}
	if bcrypt.CompareHashAndPassword(hash, []byte(password)) != nil {
		t.Error("password should match hash")
	}
	if bcrypt.CompareHashAndPassword(hash, []byte("wrongpass")) == nil {
		t.Error("wrong password should not match hash")
	}
}
