package controller

import "net/http"

func registerWelcomeRoutes() {
	http.HandleFunc("/welcome", http.HandlerFunc(welcome))
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("welcome,my friends"))
	if err != nil {
		return
	}
}
