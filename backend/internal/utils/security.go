// // backend/internal/utils/utils.go
// package utils

// import (
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"golang.org/x/crypto/bcrypt"
// )

// // HashPassword génère un hash pour le mot de passe donné.
// func HashPassword(password string) (string, error) {
//     bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//     if err != nil {
//         return "", err
//     }
//     return string(bytes), nil
// }

// // VerifyPassword vérifie si le mot de passe donné correspond au hash.
// func VerifyPassword(hashedPassword, password string) error {
//     return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// // JWTClaims représente les claims pour le JWT.
// type JWTClaims struct {
//     UserID uint `json:"user_id"`
//     jwt.StandardClaims
// }

// // GenerateJWT génère un JWT pour l'utilisateur avec un certain temps d'expiration.
// func GenerateJWT(userID uint, secret string) (string, error) {
//     claims := &JWTClaims{
//         UserID: userID,
//         StandardClaims: jwt.StandardClaims{
//             ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 72 heures d'expiration
//         },
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     signedToken, err := token.SignedString([]byte(secret))
//     if err != nil {
//         return "", err // Gérer l'erreur lors de la signature du token
//     }
//     return signedToken, nil
// }

// // ValidateJWT valide le token et retourne les claims.
// func ValidateJWT(tokenStr string, secret string) (*JWTClaims, error) {
//     claims := &JWTClaims{}
//     token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//         return []byte(secret), nil
//     })
//     if err != nil {
//         log.Printf("Error parsing token: %v", err)
//         return nil, err // Gérer l'erreur de parsing
//     }
//     if !token.Valid {
//         log.Println("Token is invalid")
//         return nil, errors.New("invalid token") // Gérer l'erreur d'invalidité
//     }
//     return claims, nil
// }

// // ExtractUserIDFromToken extrait l'ID de l'utilisateur à partir du token JWT.
// func ExtractUserIDFromToken(tokenStr string, secret string) (uint, error) {
//     claims, err := ValidateJWT(tokenStr, secret)
//     if err != nil {
//         return 0, errors.New("unable to extract user ID: invalid token") // Message d'erreur plus explicite
//     }
//     return claims.UserID, nil // Retourne l'ID de l'utilisateur
// }

// // StringSliceToJSON convertit un tableau de chaînes en JSON
// func StringSliceToJSON(slice []string) string {
//     jsonData, err := json.Marshal(slice)
//     if err != nil {
//         return "[]" // Retourne un tableau vide en cas d'erreur
//     }
//     return string(jsonData)
// }

// backend/internal/utils/utils.go

package utils

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword génère un hash pour le mot de passe donné.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// VerifyPassword vérifie si le mot de passe donné correspond au hash.
func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// JWTClaims représente les claims pour le JWT.
type JWTClaims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// GenerateJWT génère un JWT pour l'utilisateur avec un certain temps d'expiration.
func GenerateJWT(userID uint, secret string) (string, error) {
    log.Printf("Generating token for UserID: %d", userID)
    claims := &JWTClaims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}


// ValidateJWT valide le token et retourne les claims.
func ValidateJWT(tokenStr string, secret string) (*JWTClaims, error) {
    log.Println("Starting token validation...")
    
    claims := &JWTClaims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        log.Printf("Error parsing token: %v", err)
        return nil, err // Gérer l'erreur de parsing
    }

    if !token.Valid {
        log.Println("Token is invalid")
        return nil, errors.New("invalid token") // Gérer l'erreur d'invalidité
    }

    log.Printf("Token validated successfully. UserID: %v, ExpiresAt: %v", claims.UserID, claims.ExpiresAt)
    return claims, nil
}

// ExtractUserIDFromToken extrait l'ID de l'utilisateur à partir du token JWT.
func ExtractUserIDFromToken(tokenStr string, secret string) (uint, error) {
    log.Println("Extracting UserID from token...")
    
    claims, err := ValidateJWT(tokenStr, secret)
    if err != nil {
        log.Printf("Failed to extract UserID: %v", err)
        return 0, errors.New("unable to extract user ID: invalid token") // Message d'erreur plus explicite
    }

    log.Printf("UserID extracted successfully: %v", claims.UserID)
    return claims.UserID, nil // Retourne l'ID de l'utilisateur
}

// StringSliceToJSON convertit un tableau de chaînes en JSON
func StringSliceToJSON(slice []string) string {
    jsonData, err := json.Marshal(slice)
    if err != nil {
        return "[]" // Retourne un tableau vide en cas d'erreur
    }
    return string(jsonData)
}
func joinImageURLs(imageURLs []string) string {
	jsonData, err := json.Marshal(imageURLs)
	if err != nil {
		log.Printf("Erreur lors de la conversion des URLs en JSON : %v", err)
		return "[]" // Retourne un tableau vide en cas d'erreur
	}
	return string(jsonData)
}