package user

type Users []User
type User struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
}

func NewUser(id int, name string, email string, phoneNumber string) *User {
	return &User{
		ID:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}
