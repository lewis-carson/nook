package main


// importations
import (
	"os"
	"fmt"
	"github.com/zserge/webview"
)


// functions
func currentDir() (response string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return(dir)
}

func irc() {
	// start the irc backend here
}

func callback() {
	// put js callbacks here (eg.: on new irc msg, print this, etc)
}

func view() {
	windowTitle := "nook"
	windowHeight := 400
	windowWidth := 800

	w := webview.New(true)

	w.SetTitle(windowTitle)
	w.SetSize(windowWidth, windowHeight, webview.HintFixed)
	w.Navigate("file://" + currentDir() + "/view/index.html")
	w.Run()
}


// execution
func main() {
	go irc()
	go callback()
	view()
}
