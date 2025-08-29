package entity

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

func (a Auth) GenerateSession() (Session, error) {
	return Session{
		ID: a.ID,
		Token: "User Token",
	}, nil
}


type Session struct{
	ID string
	Token string
}