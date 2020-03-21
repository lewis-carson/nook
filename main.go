// ???
package main


// import libraries
import (
	// "os"
	"fmt"
	"io/ioutil"
	"github.com/zserge/webview"
)


// variables
var src = "src"
var windowTitle = "nook"
var windowHeight = 400
var windowWidth = 800
var debug = false

// functions
func printFile(file string) (response string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	data := (string(content))
	return(data)
}

func template() (response string) {
	head := `<head>` + printFile(src + "/head.html") + `</head>`
	style := `<style>` + printFile(src + "/css/main.css") + `</style>`

	nav := printFile(src + "/nav.html")
	buffer := printFile(src + "/buffer.html")
	script := printFile(src + "/script.html")
	body := `<body>` + nav + buffer + script + `</body>`

	return(head + style + body)
}

func view() {
	w := webview.New(debug)

	w.SetTitle(windowTitle)
	w.SetSize(windowWidth, windowHeight, webview.HintFixed)
	w.Navigate(`data:text/html, ` + template())
	w.Init(printFile(src + "/js/main.js"))
	w.Run()
}


// main process
func main() {
	view()
}
