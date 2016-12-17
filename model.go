package app

import "time"

type User struct {
	Email    string
	UserName string `datastore:"-"`
	Password string `json:"-"`
}

type SessionData struct {
	User
	LoggedIn  bool
	LoginFail bool
}



type Expenses struct {
	CategoryName string
	Amount float64
	Description string
	Date time.Time

}


type Category struct {
	Name        string
	Description string
}


type CategoryData struct {
	User
	Categories []Category
}