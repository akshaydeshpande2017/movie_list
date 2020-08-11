package models

type User struct {
	ID             int
	Name           string
	Password       string
	UserName       string
	HashedPassword []byte
}

type Movie struct {
	ID      int
	Name    string
	Comment string
	Rating  int
	//user    *User
}
