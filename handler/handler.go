package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// cara pertama dalam membuat handler function:
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
    
	// cara mengload file html ke handler di golang
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There's an error", http.StatusInternalServerError)
		return
	}

	// passing value to html
	data := map[string]interface{} {
		"title": "Belajar Golang Web", 
		"content": "Bismillah dapat Ganyu menang rate off low pity hehe :v",
	}
	// ~~~~~~~~~~~~~~~~~~~~~
    
	// cara menampilkannya
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "There's an error", http.StatusInternalServerError)
		return
	}
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
}

func WoiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Woi dunia, saya sedang belajar golang web mugo mugo isok wkwkwk :v"))
}

func BacotHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bacot dunia, saya mau Ganyu & Yae Miko :v"))
}
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product page : %d", idNumb)
}