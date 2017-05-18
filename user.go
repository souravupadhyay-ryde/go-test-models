package gotestmodels

var globalIdCounter = 0

type User struct {
  ID int
  Name string
  Username string
  Password string
}

func NewUser(name string, username string, password string) (*User) {
  globalIdCounter++
  id := globalIdCounter
  user := &User{ ID: id, Name: name, Username: username, Password: password }
  return user
}

