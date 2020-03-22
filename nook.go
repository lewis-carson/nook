package main

// importations
import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "os"

    socketio "github.com/googollee/go-socket.io"
    ircevent "github.com/thoj/go-ircevent"
    "github.com/zserge/webview"
)

// variables
var (
    wv            = webview.New(true)
    windowTitle   = "nook"
    windowHeight  = 400
    windowWidth   = 800
    serverPort    = "5007" // this looks like tom nook's face if you look very hard
    serverAddress = "http://localhost:" + serverPort
)

//// functions
// print the content of a file -- useful for "importation"
func printFile(file string) (response string) {
    content, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }
    return (string(content))
}

// returns the path of the current directory
func currentDir() (response string) {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return (dir)
}

func irc() {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        panic(err)
    }

    type Config struct {
        Identity struct {
            Nick     string `json:"nick"`
            RealName string `json:"realName"`
        } `json:"identity"`
        Servers []struct {
            Name        string   `json:"name"`
            Address     string   `json:"address"`
            Port        string   `json:"port"`
            AutoConnect bool     `json:"autoConnect"`
            Channels    []string `json:"channels"`
        } `json:"servers"`
    }

    var config Config
    json.Unmarshal(data, &config)

    var (
        nick       = config.Identity.Nick
        channel    = config.Servers[0].Channels[0]
        serverName = config.Servers[0].Name
        server     = config.Servers[0].Address
        port       = config.Servers[0].Port
    )

    ircobj := ircevent.IRC(nick, nick)
    ircobj.AddCallback("001", func(e *ircevent.Event) {
        ircobj.Join(channel)
        ircobj.AddCallback("PRIVMSG", func(event *ircevent.Event) {
            go newMessage(event.Nick, event.Message(), "message")
        })
        ircobj.AddCallback("JOIN", func(event *ircevent.Event) {
            go currentChannel(serverName, channel) // only diplays the channel that's just been joined
        })
    })
    ircobj.Connect(server + `:` + port)
}

func socket() {
    server, err := socketio.NewServer(nil)
    if err != nil {
        panic(err)
    }
    go server.Serve()
    defer server.Close()

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("view")))
    log.Println("nook: " + serverAddress)
    log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

func newMessage(user string, message string, action string) {
    js := "newMessage(\"" + user + "\", \"" + message + "\", \"" + action + "\");"
    inject(js)
}

func currentChannel(server string, channel string) {
    js := "currentChannel(\"" + server + "\", \"" + channel + "\");"
    inject(js)
}

func inject(js string) {
    wv.Dispatch(func() {
        wv.Eval(js)
    })
}

func view() {
    defer wv.Destroy()
    wv.SetTitle(windowTitle)
    wv.SetSize(windowWidth, windowHeight, webview.HintFixed)
    wv.Navigate(serverAddress)
    wv.Run()
}

// execution
func main() {
    go irc()
    go socket()
    view()
}
