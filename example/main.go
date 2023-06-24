package main

import (
	"fmt"

	passwordchecker "github.com/adityarizkyramadhan/password-checker"
)

func main() {
	passwordChecker := passwordchecker.NewPasswordChecker(passwordchecker.PasswordCheckerConfig{
		MinLength:      8,
		MaxLength:      20,
		AllowSymbol:    true,
		AllowSpace:     false,
		AllowNumber:    true,
		AllowLowerCase: true,
		AllowUpperCase: true,
		MustUnique:     true,
	})

	fmt.Println(passwordChecker.Check("12345678Adf)"))
}
