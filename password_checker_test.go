package passwordchecker_test

import (
	"testing"

	passwordchecker "github.com/adityarizkyramadhan/password-checker"
)

func TestPasswordChecker(t *testing.T) {
	// Buat loop test case
	testCases := []struct {
		name     string
		password string
		config   passwordchecker.PasswordCheckerConfig
		expected bool
	}{
		{
			name:     "Test Tidak Allow Space dan MustUnique [Berhasil]",
			password: "12345678Adf)",
			config: passwordchecker.PasswordCheckerConfig{
				MinLength:      8,
				MaxLength:      20,
				AllowSymbol:    true,
				AllowSpace:     false,
				AllowNumber:    true,
				AllowLowerCase: true,
				AllowUpperCase: true,
				MustUnique:     true,
			},
			expected: true,
		},
		{
			name:     "Test Tidak Allow Space dan tidak MustUnique [Berhasil]",
			password: "12345678Adf)",
			config: passwordchecker.PasswordCheckerConfig{
				MinLength:      8,
				MaxLength:      20,
				AllowSymbol:    true,
				AllowSpace:     false,
				AllowNumber:    true,
				AllowLowerCase: true,
				AllowUpperCase: true,
				MustUnique:     false,
			},
			expected: true,
		},
		{
			name:     "Test Allow Space dan MustUnique [Berhasil]",
			password: "12345678Adf) ",
			config: passwordchecker.PasswordCheckerConfig{
				MinLength:      8,
				MaxLength:      20,
				AllowSymbol:    true,
				AllowSpace:     true,
				AllowNumber:    true,
				AllowLowerCase: true,
				AllowUpperCase: true,
				MustUnique:     true,
			},
			expected: true,
		},
		{
			name:     "Test Allow Space dan MustUnique [Gagal]",
			password: "12345678Adf)",
			config: passwordchecker.PasswordCheckerConfig{
				MinLength:      8,
				MaxLength:      20,
				AllowSymbol:    true,
				AllowSpace:     true,
				AllowNumber:    true,
				AllowLowerCase: true,
				AllowUpperCase: true,
				MustUnique:     true,
			},
			expected: false,
		},
	}

	// Loop test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			passwordChecker := passwordchecker.NewPasswordChecker(tc.config)
			result := passwordChecker.Check(tc.password)

			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
