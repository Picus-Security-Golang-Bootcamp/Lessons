package auth

//User is dummy user concept
type User struct {
	Id       int
	Email    string
	Password string
	Roles    []string
}

func GetUserList() []*User {
	return []*User{
		{
			Id:       1,
			Email:    "admin@mail.com",
			Password: "1234",
			Roles:    []string{"admin", "customer"},
		},
		{
			Id:       2,
			Email:    "customer@mail.com",
			Password: "12345",
			Roles:    []string{"customer"},
		},
	}
}

func GetUser(email, password string) *User {
	for _, v := range GetUserList() {
		if v.Email == email && v.Password == password {
			return v
		}
	}

	return nil
}
