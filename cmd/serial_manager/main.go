package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/chenqi92/serial_manager/internal/serial"
	"github.com/chenqi92/serial_manager/internal/server"
)

func main() {
	// Start the web server in a goroutine
	go server.Start()

	// Create a new fyne application
	a := app.New()
	w := a.NewWindow("Serial Port Manager")

	// List available serial ports
	ports := serial.ListSerialPorts()
	portList := widget.NewList(
		func() int {
			return len(ports)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ports[i])
		},
	)

	// Set up the GUI layout
	w.SetContent(container.NewVBox(
		widget.NewLabel("Available Serial Ports:"),
		portList,
	))

	// Show and run the application
	w.ShowAndRun()
}
