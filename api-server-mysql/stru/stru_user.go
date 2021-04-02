package stru

type User struct {
	ID     int    `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type Message struct {
	Content string
	Sender  User
	Reciver User
	Time    string
}
