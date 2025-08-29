package adapter

type Auth struct {
	ID       string
	Email    string
	Password string
}

func (a Auth) VerifyPassword(plainPassword string) error {
	if a.Password != plainPassword {
		return ErrInvalidPassword
	}
	return nil
}

func (a Auth) GenerateToken() string {
	return "JWT Token"
}
