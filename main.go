package main
import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("./home.page.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", home)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Starting Server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	
}