package stru

type UserInfo struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type CreateSuccess struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}
