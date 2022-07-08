package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

//go:embed app/dist/*
var appFS embed.FS

var baseDir = "app/dist"

var staticServer http.Handler

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	fSys, err := fs.Sub(appFS, baseDir)
	if err != nil {
		panic(err)
	}
	staticServer = http.FileServer(http.FS(fSys))
	staticServer.ServeHTTP(rw, req)
}

func BlankReq(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	rw.Header().Set("Link", "rel=\"shortcut icon\" href=\"#\"")
}

func StartServer(port string) {
	http.Handle("/favicon.ico", http.HandlerFunc(BlankReq))
	http.Handle("/", http.HandlerFunc(HandleRequest))

	fmt.Println(">> Listening on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	go openbrowser("http://localhost:" + port)
	StartServer(port)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		panic(err)
	}

}
