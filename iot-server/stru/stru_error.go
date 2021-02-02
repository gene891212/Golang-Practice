package stru

type ErrorMessage struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type ErrorResp struct {
	Error ErrorMessage `json:"error"`
}
