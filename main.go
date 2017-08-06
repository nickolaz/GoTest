package main

import (
	"html/template"
	"strings"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"os"
	"math/rand"
)

type Page struct {
	Title string
	Body template.HTML
	Img	 template.HTML
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/",GetIndexHandle).Methods("GET")
	r.HandleFunc("/focoon.png",GetFocoOnHandle)
	r.HandleFunc("/focooff.png",GetFocoOffHandle)
	server := &http.Server{
		Addr:	":8080",
		Handler:	r,
		ReadHeaderTimeout:	10*time.Second,
		WriteTimeout:	10*time.Second,
		MaxHeaderBytes:	1<<20,
	}
	server.ListenAndServe()
}

func GetFocoOnHandle(w http.ResponseWriter, r *http.Request) {
	f,_ := os.Open("focoon.png")
	defer f.Close()
	fi,_ := f.Stat()
	http.ServeContent(w,r,fi.Name(),fi.ModTime(),f)
}

func GetFocoOffHandle(w http.ResponseWriter, r *http.Request) {
	f,_ := os.Open("focooff.png")
	defer f.Close()
	fi,_ := f.Stat()
	http.ServeContent(w,r,fi.Name(),fi.ModTime(),f)
}

func GetIndexHandle(writer http.ResponseWriter, request *http.Request) {
	tmp1:= template.New("index.html")
	tmp1.Funcs(template.FuncMap{
		"uppercase": uppercase,
	})
	tmp1,_ = tmp1.ParseFiles("index.html")
	_ = tmp1.Execute(writer,Page{
		Title: "MI TITULO",
		Body: "BODY"/*template.HTML(`<script>alert('Hola')</script>`)*/,
		Img:	Estadoluz(),
	})
}

func uppercase(str string) string  {
	return strings.ToUpper(str)
}

func Estadoluz() template.HTML  {
	if(GetEstado()){
		return template.HTML(`<img src="focoon.png" alt="foco on" >`)
	}else {
		return template.HTML(`<img src="focooff.png" alt="foco of" >`)
	}
}
func GetEstado() bool{
	return rand.Intn(2)==1
}