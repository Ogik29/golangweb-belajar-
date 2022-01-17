package handler

import (
	entity "golangweb/entitiy"
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
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There's an error", http.StatusInternalServerError)
		return
	}
    
	// data := map[string]interface{} {
	// 	"content": "Bismillah dapat Ganyu menang rate off low pity hehe :v",
	// }
    
	// passing value to html
	data := []entity.Product{
		{ID: 1, Name: "Asus", Price: 11000000, Stock: 5},
		{ID: 2, Name: "Acer", Price: 10000000, Stock: 8},
		{ID: 3, Name: "Lenovo", Price: 8000000, Stock: 4},
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

	// fmt.Fprintf(w, "Product page : %d", idNumb)
    
	data := map[string]interface{} {
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There's an error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "There's an error", http.StatusInternalServerError)
		return
	}
}


func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("This is GET"))
	case "POST":
		w.Write([]byte("This is POST"))
	default: 
	    http.Error(w, "There's an error", http.StatusBadRequest)
	}
		
}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
		    http.Error(w, "There's an error", http.StatusInternalServerError)
		    return
		}

		err = tmpl.Execute(w, nil) 
		if err != nil {
			log.Println(err)
		    http.Error(w, "There's an error", http.StatusInternalServerError)
		    return
		}

		return
	}

	http.Error(w, "There's an error", http.StatusBadRequest)
}

// menunjukkan hasil post dari user
func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		    http.Error(w, "There's an error", http.StatusInternalServerError)
		    return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{} {
			"name": name,
			"message": message, 
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "There's an error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "There's an error", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "There's an error", http.StatusBadRequest)
}
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~