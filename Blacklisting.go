package main

import (
	"regexp"
)

// CheckCommand is a function that check if command is valid for executing or not
func CheckCommand(command string) bool {
	r, _ := regexp.Compile("rm")

	result := r.MatchString(command)

	return !result
}
