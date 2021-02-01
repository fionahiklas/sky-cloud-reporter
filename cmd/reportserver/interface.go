package main

import (
	"net/http"
)

type GrabHandler interface {
	HttpHandler() func(w http.ResponseWriter, r *http.Request)
}
