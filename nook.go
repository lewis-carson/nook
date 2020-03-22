package main

// importations
import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/zserge/webview"
	ircevent "github.com/thoj/go-ircevent"
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

	//user := "tom nook"
	//message := "sent from golang!"
	//newMessage(user, message)
}

func newMessage(user string, message string) {
	js := `
	log = document.getElementById("log");
	line = document.createElement("li");
	userName = document.createElement("a");
	userNameSpan = document.createElement("span");
	userMessageSpan = document.createElement("span");

	userText = document.createTextNode("` + user + `");
	userMessage = document.createTextNode(" ` + message + ` ");

	log.appendChild(line);
	line.appendChild(userNameSpan).classList.add("name", "msg")

	userNameSpan.appendChild(userName);
	userName.appendChild(userText);

	line.appendChild(userMessageSpan).classList.add("message");
	userMessageSpan.appendChild(userMessage);`

	wv.Dispatch(func() {
		fmt.Println("trying to inject msg...")
		wv.Eval(js)
	})
}

func sendMessage() {
	err := wv.Bind("sendMessage", func() string {
		// fmt.Println()
		return "<message>"
	})

	if err != nil {
		panic(err)
	}
}

// execution
func main() {
	go irc()
	view()
}
