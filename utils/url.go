package utils

import (
	"fmt"
	"hash/fnv"
)

// generateShortURL creates a shortened URL string using a hash function
func GenerateShortURL(originalURL string) string {
 h := fnv.New32a()
 h.Write([]byte(originalURL))
 return fmt.Sprintf("%x", h.Sum32())
}