package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
    
	// Cara ke 2 (harus berada didalam func main())
	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	} 
	mux.HandleFunc("/about", aboutHandler)
    
	mux.HandleFunc("/", handler.HomeHandler) // root
	mux.HandleFunc("/woi", handler.WoiHandler)
	mux.HandleFunc("/bacot", handler.BacotHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	// cara load folder assets di golang
	fileserver := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

    
	// cara ke 3 (anonymous function)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile"))
	})

	log.Println("Memulai web pada port 8080")
    err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
	
}
