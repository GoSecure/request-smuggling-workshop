// Lightly modified example from: https://github.com/thrawn01/h2c-golang-example
package main

import (
    "fmt"
    "golang.org/x/net/http2"
    "golang.org/x/net/http2/h2c"
    "net/http"
    "os"
	"html/template"
)

func checkErr(err error, msg string) {
    if err == nil {
        return
    }
    fmt.Printf("ERROR: %s: %s\n", msg, err)
    os.Exit(1)
}

func main() {
    H2CServerUpgrade()
}


type HomePage struct {
}
type LoginPage struct {
    Username string
    Password string
}

// This server supports "H2C upgrade" and "H2C prior knowledge" along with
// standard HTTP/2 and HTTP/1.1 that golang natively supports.
func H2CServerUpgrade() {
    h2s := &http2.Server{}

    handler := http.NewServeMux()
    handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "text/html")
        
        t, _ := template.ParseFiles("index.htm")
		var p = &HomePage{}
        t.Execute(w, p)
    })

    handler.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "text/html")

        t, _ := template.ParseFiles("admin.htm")
		var p = &LoginPage{Username: "",Password:""}
        t.Execute(w, p)

    })

    server := &http.Server{
        Addr:    "0.0.0.0:80",
        Handler: h2c.NewHandler(handler, h2s),
    }

    fmt.Printf("Listening [0.0.0.0:80]...\n")
    checkErr(server.ListenAndServe(), "while listening")
}
