package ownpack

import (
	"log"
)

func LogInfo(message string) {
	log.Printf("info- %v", message)
}
func LogWarning(message string) {
	log.Printf("warn- %v", message)
}
func LogError(message string) {
	log.Printf("error - %v", message)
}
