package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dramamask/password-police-go/password"
	"github.com/dramamask/password-police-go/policy"
)

func main() {
	policy.LoadFromFile("password-police.json")

	policyDescription := policy.Describe()
	fmt.Println(policyDescription)

	myPassword := getPwdFromCli()
	fmt.Printf("Password entered: \"%s\".\n", myPassword)
	err := password.Validate(myPassword)
	if err == nil {
		fmt.Println("Password is valid.")
	} else {
		fmt.Println("Password is invalid.")
	}
}

func getPwdFromCli() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter password: ")
	pwd, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading password.")
	}

	pwd = pwd[:len(pwd)-1] // Remove newline character from password.

	return pwd
}
