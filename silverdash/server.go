package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World  %s!", r.URL.Path[1:])
}



func main() {

  http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/", htmlHandler)
  http.HandleFunc("/js/", jsHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Root")
    //http.Redirect(w, r, "/404.html", http.StatusFound)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage( title )
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadHtmlPage( "view/" + title )
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        http.Redirect(w, r, "/view/404.html", http.StatusFound)
    } else {
      fmt.Fprintf(w, "%s",  p.Body)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage( title )
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func loadHtmlPage(title string) (*Page, error) {
    filename := title
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}


func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("---")
        fmt.Println(err)
        fmt.Println("---")
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
