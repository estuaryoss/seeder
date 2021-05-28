package utils

import (
	"crypto/rand"
	"fmt"
)

func GetKeysFromMap(myMap map[string]interface{}) []string {
	keys := make([]string, 0, len(myMap))

	for k, _ := range myMap {
		keys = append(keys, k)
	}

	return keys
}

func GetKeysFromMapInt(myMap map[string]int) []string {
	keys := make([]string, 0, len(myMap))

	for k, _ := range myMap {
		keys = append(keys, k)
	}

	return keys
}

func GenerateRandomId(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
