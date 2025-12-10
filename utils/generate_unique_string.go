package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateUniqueString() string {
	uuid := uuid.New().String()
	cleanUUID := strings.ReplaceAll(uuid, "-", "")
	return cleanUUID[:10]
}

func GenerateJWTKey() (string, error) {
	bytes := make([]byte, 64) // 512 bit
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("random read failed: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
