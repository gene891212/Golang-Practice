package libs

import (
	"fmt"
	"net/http"
	"text/template"
)

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		fmt.Printf("%+v\n", r)

	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, _ := template.ParseFiles("./create.html")
		t.Execute(w, nil)
	}

}
