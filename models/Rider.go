package models

type Rider struct {
    User
}

func NewRider(user User) *Rider {
    return &Rider{User: user}
}


