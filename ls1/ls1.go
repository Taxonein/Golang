package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("POD NAME:" + os.Getenv("HOSTNAME") + " version: " + os.Getenv("APPVER"))
	fmt.Println("CURRPWD: " + os.Getenv("PWD"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POD NAME:"+os.Getenv("HOSTNAME")+" version: "+os.Getenv("APPVER"))
		fmt.Fprintln(w, "CURRPWD: "+os.Getenv("PWD"))
	})
	http.ListenAndServe(":80", nil)
}
