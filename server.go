package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)
var count int = 0
func main() {
	port := os.Args[1];
	fmt.Println(port);
	fmt.Println("Web server started on port " + port)	
	
	http.HandleFunc("/IncreaseCount", increaseCount);
	
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/increaseCounter" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	count+=1;
	currentTime := time.Now();

	fmt.Printf("[%v] Request %v has been recieved \n",  currentTime.Format(time.ANSIC), count);
}

func increaseCount(w http.ResponseWriter, r *http.Request) {
	count+=1;
	fmt.Fprintf(w, "%v", count)
}