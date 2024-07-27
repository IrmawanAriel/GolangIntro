package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func genPass(password string, level string) string {

	if len(password) < 6 {
		fmt.Print("Character password harus lebih dari 6")
		return ""
	}

	var extraCharacter string

	switch strings.ToLower(level) {
	case "low":
		extraCharacter = "qwertyuiopasdfghjklzxcvbnm"
	case "med":
		extraCharacter = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	case "strong":
		extraCharacter = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	default:
		fmt.Print("Level tidak valid. pilih 'low','med' atau 'strong'.")
		return ""
	}

	passwordList := []rune(password)

	for i := len(password); i < 12; i++ {
		randomChar := rune(extraCharacter[rand.Intn(len(extraCharacter))])
		passwordList = append(passwordList, randomChar)
	}

	rand.Shuffle(len(passwordList), func(i, j int) {
		passwordList[i], passwordList[j] = passwordList[j], passwordList[i]
	})

	generatedPassword := string(passwordList)

	fmt.Print(generatedPassword)

	return generatedPassword

}
