package main

// importations
import (
    "fmt"
    "io/ioutil"
    "os"
    "time"

    ircevent "github.com/thoj/go-ircevent"
    "github.com/zserge/webview"
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

func config() {

}

func view() {
    defer wv.Destroy()
    wv.SetTitle(title)
    wv.SetSize(windowWidth, windowHeight, webview.HintFixed)
    wv.Navigate("file://" + currentDir() + "/view/index.html")

    err := wv.Bind("sendMessage", func(message string) {
        fmt.Println("message received")
    })

    if err != nil {
        fmt.Println(err)
    }

    wv.Run()
}

func irc() {
    time.Sleep(time.Second * 1) // sleep for 1 second otherwise messages won't load -- needs work

    ircobj := ircevent.IRC("nook", "nook")
    ircobj.AddCallback("001", func(e *ircevent.Event) {
        ircobj.Join("#letirc")
        ircobj.Privmsg("#letirc", "send with <3 from nook")
        ircobj.AddCallback("PRIVMSG", func(event *ircevent.Event) {
            go newMessage(event.Nick, event.Message(), "message")
        })
    })
    ircobj.Connect("irc.rizon.net:7000")
}

func newMessage(user string, message string, action string) {
    js := "newMessage(\"" + user + "\", \"" + message + "\", \"" + action + "\");"
    inject("message", js)
}

func sendMessage(message string) {
}

func changeChannel(server string, channel string) {
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
    // go irc()
    view()
}
