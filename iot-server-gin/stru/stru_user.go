package stru

type UserInfo struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type IndexData struct {
	AllAccount string
}
