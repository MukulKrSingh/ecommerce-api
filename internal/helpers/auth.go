package helpers

type Auth struct {
	Secret string
}

func (a Auth) CreateHashedPassword(p string) (string, error) {

	return "", nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {

	return "", nil
}

func (a Auth) VerifyPassword(p string, hp string) error {

	return nil
}
func (a Auth) VerifyToken(t string) error {

	return nil
}
