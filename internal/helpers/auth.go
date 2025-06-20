package helpers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ecommerce-api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func NewAuth(secret string) Auth {
	if secret == "" {
		log.Println("Secret key is empty, using default secret")
		// Use a default secret if none is provided
		secret = "default_secret"
	}
	return Auth{
		Secret: secret}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {

	if len(p) < 8 {
		return "", errors.New("password must be at least 8 characters long")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {

		log.Println("Error hashing password:", err)
		// Return an empty string and the error
		return "", errors.New("failed to hash password")

	}

	return string(hashedPwd), nil

}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	// This function should generate a JWT token using the provided id, email, and role.

	if email == "" || role == "" {
		return "", errors.New("invalid parameters for token generation")

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expires in 30 days
	})

	signedToken, err := token.SignedString([]byte(a.Secret))

	if err != nil {

		// Return an empty string and the error
		// This is a critical error, so we return a more descriptive message.
		// In a real application, you might want to log this error or handle it differently
		log.Println("Error signing token:", err)
		return "", errors.New("failed to sign token: ")
	}

	return signedToken, nil

}

func (a Auth) VerifyPassword(p string, hp string) error {
	if len(p) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
	if err != nil {
		log.Println("Error verifying password:", err)
		return errors.New("invalid password")
	}
	// If the password matches, return nil (no error)
	log.Println("Password verified successfully")

	return nil
}
func (a Auth) VerifyToken(t string) (model.User, error) {
	//Bearer t435627120u121wcdxe34 - token example
	if t == "" {
		return model.User{}, errors.New("token is empty")
	}
	tokenArray := strings.Split(t, " ")
	if len(tokenArray) != 2 {
		return model.User{}, errors.New("invalid token format")
	}
	if tokenArray[0] != "Bearer" {
		return model.User{}, errors.New("token must start with 'Bearer'")
	}
	tokenString := tokenArray[1]
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header)
			return nil, fmt.Errorf("unexpected signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return model.User{}, errors.New("invalid token")
	}
	if !token.Valid {
		log.Println("Token is not valid")
		return model.User{}, errors.New("token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(claims["exp"].(float64)) < float64(time.Now().Unix()) {
			log.Println("Token has expired")
			return model.User{}, errors.New("token has expired")
		}
		u := model.User{}
		// Extract user information from the claims
		u.ID = uint(claims["user_id"].(float64)) // JWT claims are typically float64
		u.Email = claims["email"].(string)
		u.Role = claims["role"].(string)
		log.Println("Token is valid and user information extracted successfully")
		return u, nil
	}
	log.Println("Token claims are not valid")
	return model.User{}, errors.New("invalid token claims")
}

func (a Auth) Authorize(ctx *gin.Context) error {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		log.Println("Authorization header is missing")
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
		return errors.New("authorization header is required")
	}
	user, err := a.VerifyToken(authHeader)
	if err != nil {
		log.Println("Error verifying token:", err)
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
		return err
	}
	// Store the user in the context for later use
	ctx.Set("user", user)
	log.Println("User authorized successfully:", user.Email)
	ctx.Next() // Call the next handler in the chain
	// If you want to return the user information, you can do so here
	// ctx.JSON(200, gin.H{"user": user})
	// But typically, you would just continue to the next handler
	// and let it handle the response.
	log.Println("Authorization successful for user:", user.Email)
	// Return nil to indicate successful authorization
	ctx.Next()
	// No error, so we return nil
	log.Println("Authorization completed successfully")
	return nil
}
func (a Auth) GetCurrentUser(ctx *gin.Context) (model.User, error) {
	user, exists := ctx.Get("user")
	if !exists {
		log.Println("No user found in context")
		return model.User{}, errors.New("no user found in context")
	}
	u, ok := user.(model.User)
	if !ok {
		log.Println("User in context is not of type model.User")
		return model.User{}, errors.New("user in context is not of type model.User")
	}
	log.Println("Current user retrieved successfully:", u.Email)
	// Return the user information
	if u.ID == 0 {
		log.Println("User ID is zero, indicating no user is logged in")
		return model.User{}, errors.New("no user is logged in")
	}
	if u.Email == "" {
		log.Println("User email is empty, indicating no user is logged in")
		return model.User{}, errors.New("no user is logged in")
	}
	if u.Role == "" {
		log.Println("User role is empty, indicating no user is logged in")
		return model.User{}, errors.New("no user is logged in")
	}
	log.Println("Current user:", u.Email, "with role:", u.Role)
	// If everything is fine, return the user
	log.Println("Returning current user:", u.Email)

	return u, nil
}
