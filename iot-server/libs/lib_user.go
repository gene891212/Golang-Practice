package libs

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iot/stru"
)

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch r.Method {

	case "POST":
		r.ParseForm()

		user := stru.UserInfo{
			Account:  r.PostFormValue("account"),
			Password: r.PostFormValue("password"),
		}

		InsertData(user)
		now, _ := time.Now().MarshalText()

		message := stru.CreateSuccess{
			Timestamp: string(now),
			Status:    200,
			Message:   "account created successful",
		}

		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			RaiseError(w, 500, err.Error(), r.URL.Path)
		}
	default:
		// t, _ := template.ParseFiles("./create.html")
		// t.Execute(w, nil)
	}
}
