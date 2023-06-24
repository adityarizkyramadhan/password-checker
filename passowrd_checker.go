package passwordchecker

type PasswordCheckerConfig struct {
	MinLength      int
	MaxLength      int
	AllowSymbol    bool
	AllowSpace     bool
	AllowNumber    bool
	AllowLowerCase bool
	AllowUpperCase bool
	MustUnique     bool
	containsSpace  bool
	containsSymbol bool
	containsNumber bool
	containsLower  bool
	containsUpper  bool
}

type PasswordChecker struct {
	config PasswordCheckerConfig
}

func NewPasswordChecker(config PasswordCheckerConfig) *PasswordChecker {
	return &PasswordChecker{config: config}
}

func (p *PasswordChecker) Check(password string) bool {
	if len(password) < p.config.MinLength || len(password) > p.config.MaxLength {
		return false
	}

	for _, c := range password {
		if !p.checkChar(c) {
			return false
		}
	}

	if p.config.MustUnique {
		return p.checkUnique()
	}

	return true
}

func (p *PasswordChecker) checkChar(c rune) bool {
	if p.config.AllowSymbol && isSymbol(c) {
		p.config.containsSymbol = true
		return true
	}

	if p.config.AllowSpace && isSpace(c) {
		p.config.containsSpace = true
		return true
	}

	if p.config.AllowNumber && isNumber(c) {
		p.config.containsNumber = true
		return true
	}

	if p.config.AllowLowerCase && isLowerCase(c) {
		p.config.containsLower = true
		return true
	}

	if p.config.AllowUpperCase && isUpperCase(c) {
		p.config.containsUpper = true
		return true
	}

	return false
}

func (p *PasswordChecker) checkUnique() bool {
	if p.config.containsSpace != p.config.AllowSpace {
		return false
	}

	if p.config.containsSymbol != p.config.AllowSymbol {
		return false
	}

	if p.config.containsNumber != p.config.AllowNumber {
		return false
	}

	if p.config.containsLower != p.config.AllowLowerCase {
		return false
	}

	if p.config.containsUpper != p.config.AllowUpperCase {
		return false
	}

	return true
}

func isSymbol(c rune) bool {
	return (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~')
}

func isSpace(c rune) bool {
	return c == ' '
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func isLowerCase(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isUpperCase(c rune) bool {
	return c >= 'A' && c <= 'Z'
}
