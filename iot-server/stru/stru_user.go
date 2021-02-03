package stru

type UserInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CreateSuccess struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}
