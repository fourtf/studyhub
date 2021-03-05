package authorization

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (u *User) PutPID(pid string)           { u.Username = pid }
func (u User) GetPID() string               { return u.Username }
func (u User) GetPassword() string          { return u.Password }
func (u *User) PutPassword(password string) { u.Password = password }

func (u User) Validate() []error {
	return nil
}

/*func (u User) GetArbitrary() (arbitrary map[string]string){
	return map[string] string {}
}

func (u *User) PutArbitrary(arbitrary map[string]string){

	if username, ok = arbitrary["email"]; ok {
		u.Username = username
	}
}*/