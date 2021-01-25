package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_addAccount(t *testing.T) {
	temp := httptest.NewRecorder()
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
			"Test addAccount",
			args{
				temp,
				&http.Request{
					Method: http.MethodGet,
					URL: &url.URL{Path: "/all_user"},
					Body: ,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func Test_allUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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
