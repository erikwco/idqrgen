package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"html/template"
	"image/png"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

func main() {

	// multiplexor
	mux := http.NewServeMux()

	// entry point
	mux.HandleFunc("/", homeHandler)

	// static file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// generator
	mux.HandleFunc("/generator/", generatorHandler)

	// starting
	log.Printf("Idear QRGenerator Server Started ....")
	if err := http.ListenAndServe(":8181", mux); err != nil {
		log.Panic(err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Generador QR Code Idear"}
	t, _ := template.ParseFiles("templates/generator.html")
	_ = t.Execute(w, p)
}

func generatorHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, qrCode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
