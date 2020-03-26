package shared

import (
	"fmt"
	"log"
)

// Print gracefully print a message on STDOUT
func Print(str string) {
	fmt.Printf(str + "\n")
}

// PrintErrorIfExists Just prints a error if exits
func PrintErrorIfExists(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}

// FatalError closes the program if a fatal error occurs
func FatalError(err error, msg string) {
	if err != nil {
		log.Fatalf("[%s]: %s", msg, err)
	}
}
