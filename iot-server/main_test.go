package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createUser(t *testing.T) {
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test addUser",
			args: args{
				w: resp,
				r: &http.Request{
					Method: http.MethodGet,
					RequestURI: "/?account=testA&password=testP",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler.ServeHTTP(resp, tt.args.r)
			if status := resp.Code; status != http.StatusOK{
				t.Errorf("status = %v, want 200}", status)
			}
		})
	}
}

func Test_allUser(t *testing.T) {
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(allUser)
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test allUser",
			args: args{
				w: resp,
				r: httptest.NewRequest("GET", "/", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler.ServeHTTP(tt.args.w, tt.args.r)
			if status := resp.Code; status != http.StatusOK{
				t.Errorf("status = %v, want 200}", status)
			}
		})
	}
}

func Test_insertData(t *testing.T) {
	type args struct {
		account  string
		password string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
