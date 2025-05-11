package utils

import (
	"fmt"
	"hash/crc32"

	"github.com/golang-jwt/jwt/v5"
	// "github.com/inidaname/mosque/api_gateway/pkg/constant"
	// "github.com/inidaname/mosque/api_gateway/pkg/store/cache"
	"github.com/inidaname/mosque/api_gateway/pkg/types"
)

func GenerateChecksum(b []byte) uint32 {
	hasher := crc32.NewIEEE()
	hasher.Write(b)

	return hasher.Sum32()
}

// func GenerateAPIKeys(env string) (publicKey, secretKey string, err error) {
// 	var publicPrefix, secretPrefix string

// 	if env == constant.EnvironmentLive {
// 		publicPrefix = "flk"
// 		secretPrefix = "fsk"
// 	} else {
// 		publicPrefix = "ftk"
// 		secretPrefix = "ftsk"
// 	}

// 	// Generate public key
// 	publicBytes := make([]byte, 16)
// 	if _, err := rand.Read(publicBytes); err != nil {
// 		return "", "", err
// 	}
// 	publicBaseKey := hex.EncodeToString(publicBytes)
// 	publicChecksum := GenerateChecksum([]byte(publicBaseKey))
// 	publicKey = fmt.Sprintf("%s_%s_%d", publicPrefix, publicBaseKey, publicChecksum)

// 	// Generate secret key
// 	secretBytes := make([]byte, 32)
// 	if _, err := rand.Read(secretBytes); err != nil {
// 		return "", "", err
// 	}
// 	secretBaseKey := hex.EncodeToString(secretBytes)
// 	secretChecksum := GenerateChecksum([]byte(secretBaseKey))
// 	secretKey = fmt.Sprintf("%s_%s_%d", secretPrefix, secretBaseKey, secretChecksum)

// 	return publicKey, secretKey, nil
// }

// func ValidatePublicKey(publicKey string) (bool, error) {
// 	// Split the public key into parts
// 	publicKeyParts := strings.Split(publicKey, "_")
// 	if len(publicKeyParts) != 3 {
// 		fmt.Println("Invalid public key format:", publicKey)
// 		return false, fmt.Errorf("invalid public key format")
// 	}

// 	// Validate the public key prefix (flk or ftk)
// 	prefix := publicKeyParts[0]
// 	if prefix != "flk" && prefix != "ftk" {
// 		fmt.Println("Invalid public key prefix:", prefix)
// 		return false, fmt.Errorf("invalid public key prefix")
// 	}

// 	checksum, err := strconv.ParseUint(publicKeyParts[2], 10, 32)
// 	if err != nil {
// 		fmt.Println("Invalid checksum in public key:", publicKeyParts[2])
// 		return false, fmt.Errorf("invalid checksum in public key: %v", err)
// 	}

// 	// Validate that the middle part of the public key is alphanumeric
// 	middlePart := publicKeyParts[1]
// 	match, err := regexp.MatchString("^[0-9a-zA-Z]+$", middlePart)
// 	if err != nil {
// 		fmt.Println("Regex error for public key middle part:", err)
// 		return false, err
// 	} else if !match {
// 		fmt.Println("Invalid format for public key middle part:", middlePart)
// 		return false, fmt.Errorf("invalid format for public key middle part")
// 	}

// 	// Validate checksum correctness for the public key
// 	calculatedChecksum := GenerateChecksum([]byte(middlePart))
// 	if uint32(checksum) != calculatedChecksum {
// 		fmt.Printf("Checksum mismatch: calculated %d, provided %d\n", calculatedChecksum, checksum)
// 		return false, fmt.Errorf("checksum mismatch for public key")
// 	}

// 	fmt.Println("Public key is valid:", publicKey)
// 	return true, nil
// }

// func ValidateApiKey(key string) (bool, error) {
// 	livePrefix := "flk"
// 	testPrefix := "ftk"
// 	secretLivePrefix := "fsk"
// 	secretTestPrefix := "ftsk"

// 	keyParts := strings.Split(key, "_")
// 	if len(keyParts) != 3 {
// 		return false, fmt.Errorf("invalid key format: expected 3 parts separated by underscores")
// 	}

// 	// Validate the prefix
// 	prefix := keyParts[0]
// 	validPrefixes := []string{livePrefix, testPrefix, secretLivePrefix, secretTestPrefix}
// 	isValidPrefix := false
// 	for _, validPrefix := range validPrefixes {
// 		if prefix == validPrefix {
// 			isValidPrefix = true
// 			break
// 		}
// 	}

// 	if !isValidPrefix {
// 		return false, fmt.Errorf("invalid key prefix: '%s'. Expected one of [%s, %s, %s, %s]",
// 			prefix, livePrefix, testPrefix, secretLivePrefix, secretTestPrefix)
// 	}

// 	// Validate checksum
// 	checksum, err := strconv.ParseUint(keyParts[2], 10, 32)
// 	if err != nil {
// 		return false, fmt.Errorf("invalid checksum: '%s'. It must be a valid integer", keyParts[2])
// 	}

// 	// Validate middle part format
// 	middlePart := keyParts[1]
// 	match, err := regexp.MatchString("^[0-9a-zA-Z]+$", middlePart)
// 	if err != nil {
// 		return false, fmt.Errorf("regex error while validating middle part of the key: %v", err)
// 	} else if !match {
// 		return false, fmt.Errorf("invalid format for the middle part of the key: '%s'. It must contain only alphanumeric characters", middlePart)
// 	}

// 	// Validate checksum correctness
// 	calculatedChecksum := GenerateChecksum([]byte(middlePart))
// 	if uint32(checksum) != calculatedChecksum {
// 		return false, fmt.Errorf("checksum mismatch: provided '%d', but calculated '%d'", checksum, calculatedChecksum)
// 	}

// 	return true, nil
// }

// func IsAPIKeyCache(cache cache.CacheService, key string) (bool, *types.ApiKeyValidityCache) {
// 	if cached, found := cache.Get(key); found {
// 		return true, cached.(*types.ApiKeyValidityCache)
// 	}
// 	return false, nil
// }

// func CacheAPIKey(cache cache.CacheService, key string, valid types.ApiKeyValidityCache) {
// 	cache.SetDefault(key, valid)
// 	return
// }

type JWTAuthenticator struct {
	secret string
	aud    string
	iss    string
}

func NewJWTAuthenticator(secret, aud, iss string) *JWTAuthenticator {
	return &JWTAuthenticator{secret, iss, aud}
}

func (a *JWTAuthenticator) GenerateToken(claims types.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(a.secret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(a.aud),
		jwt.WithIssuer(a.aud),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
