package authorization

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}


func (u *User) PutPID(pid string) { u.Email = pid }
func (u User) GetPID() string     { return u.Email }
func (u User) GetPassword() string { return u.Password }
func (u *User) PutPassword(password string) { u.Password = password }