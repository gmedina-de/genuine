package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	buttonNew, err := gtk.ButtonNew()
	buttonNew.SetLabel("Hello world")
	buttonNew.Connect("button_press_event", func() {
		log.Println("TEST 1234")
	})

	barNew, err := gtk.HeaderBarNew()
	barNew.Add(buttonNew)
	barNew.SetShowCloseButton(true)
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	win.SetTitlebar(barNew)
	win.Add(buttonNew)
	win.Add(l)
	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}
