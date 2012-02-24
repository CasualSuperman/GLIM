package main

import (
	"fmt"
	gtk "github.com/mattn/go-gtk/gtk"
	"os"
)

func main(){
	fmt.Println("Hello, world.")
	initGTK()
}

func initGTK() {
	gtk.Init(&os.Args)
}

func showLogin() {
	window := gtk.Window(gtk.GTK_WINDOW_POPUP)
	window.Connect("destroy", gtk.MainQuit)
	window.SetSizeRequest()
	window.ShowAll()

	gtk.Main()
}
