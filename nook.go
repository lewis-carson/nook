package main

// importations
import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/zserge/webview"
	"github.com/thoj/go-ircevent"
)

// variables
var wv = webview.New(true)
var title = "nook"
var windowHeight = 400
var windowWidth = 800

//// functions
// print the content of a file -- useful for "importation"
func printFile(file string) (response string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	data := (string(content))
	return (data)
}

// returns the path of the current directory
func currentDir() (response string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return (dir)
}

func view() {
	defer wv.Destroy()
	wv.SetTitle(title)
	wv.SetSize(windowWidth, windowHeight, webview.HintFixed)
	wv.Navigate("file://" + currentDir() + "/view/index.html")
	wv.Run()
}

func irc() {
	time.Sleep(time.Second * 1) // sleep for 1 second otherwise messages won't load -- needs work

	ircobj := ircevent.IRC("nook", "nook")
	ircobj.AddCallback("001", func(e *ircevent.Event) {
		ircobj.Join("#letirc")
		ircobj.Privmsg("#letirc", "send with <3 from nook")
		ircobj.AddCallback("PRIVMSG", func(event *ircevent.Event) {
			go newMessage(event.Nick, event.Message())
		});
	})
	ircobj.Connect("irc.rizon.net:7000")
}

func newMessage(user string, message string, action string) {
    js := "newMessage(\"" + user + "\", \"" + message + "\", \"" + action + "\");"
    inject("message", js)
}

func changeChannel(server string, channel string) {
    wv.Bind("changeChannel", func() {
        fmt.Println(server, channel)
    })
}

func inject(action string, js string) {
    wv.Dispatch(func() {
        switch action {
        case "message":
            wv.Eval(js)
        }
    // add other injections here
    })
}

// execution
func main() {
	go irc()
	view()
}
