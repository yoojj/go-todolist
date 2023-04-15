package common

type PageData struct {
	PageTitle    string
	ContentTitle string
	User         User
}

type JsonData struct {
	Status Status `json:"status" binding:"required"`
	Data   Data   `json:"data"`
}

type User struct {
	Firstname string
	Lastname  string
	LoginDate string
}

type Status struct {
	Code int
	Desc string
}

type Data map[string]interface{}
