package server

import (
	"fmt"
	"github.com/chenqi92/serial_manager/internal/serial"
	"net/http"
)

// handleListPorts handles the HTTP requests to list serial ports
func handleListPorts(w http.ResponseWriter, r *http.Request) {
	ports := serial.ListSerialPorts()
	fmt.Fprintf(w, "Available Serial Ports: %v", ports)
}

// Start initializes the web server
func Start() {
	http.HandleFunc("/list-ports", handleListPorts)
	http.ListenAndServe(":8080", nil)
}
