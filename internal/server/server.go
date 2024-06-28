package server

import (
	"fmt"
	"github.com/your_username/serial_manager/internal/serial"
	"github.com/zserge/webview"
	"net/http"
)

// homePage handles the HTTP requests to the home page
func homePage(w http.ResponseWriter, r *http.Request) {
	ports := serial.ListSerialPorts()
	response := "<h1>Available Serial Ports:</h1><ul>"
	for _, port := range ports {
		response += fmt.Sprintf("<li>%s</li>", port)
	}
	response += "</ul>"
	fmt.Fprintf(w, response)
}

// Start initializes the server and webview
func Start() {
	http.HandleFunc("/", homePage)
	go http.ListenAndServe(":8080", nil)

	w := webview.New(webview.Settings{
		URL: "http://localhost:8080",
	})
	defer w.Exit()
	w.Run()
}
