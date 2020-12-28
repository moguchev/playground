package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		tml, err := template.New("").Funcs(
			template.FuncMap{
				"html": func(value interface{}) template.HTML {
					return template.HTML(fmt.Sprint(value))
				},
			}).ParseGlob(path.Join("./", "*.html"))

		// tml, err := template.ParseGlob(path.Join("./", "*.html"))
		if err != nil {
			log.Fatalf("parse %v", err)
		}

		type Data struct {
			String string `json:"string"`
		}
		data := Data{String: "<b>Этот текст будет полужирным, <i>а этот — ещё и курсивным</i>.</b>"}

		// Escape html tags
		buf := new(bytes.Buffer)
		err = tml.ExecuteTemplate(buf, "template.html", data)
		if err != nil {
			log.Fatalf("ExecuteTemplate %v", err)
		}

		//tml.Execute(w, template.HTML(`<b>World</b>`))

		w.Write(buf.Bytes())
	})

	fmt.Println("Listen on 5000")
	http.ListenAndServe(":5000", m)

}
