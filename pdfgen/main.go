package main

import (
	"fmt"
	"net/http"

	"github.com/moguchev/playground/pdfgen/pdfGenerator"
)

type EmployeeTemplate struct {
	Fio         string
	PhoneNumber string
	OrgStruct   string
	TabelNumber string
	Position    string
	Outdate     string
}

type PointTemplate struct {
	Number       int
	Title        string
	Description  string
	Approvers    []ApproverTemplate
	ApprovedDate string
	Comment      string
}

type OutlistTemplate struct {
	Title    string
	Employee EmployeeTemplate
	Points   []PointTemplate
}

type ApproverTemplate struct {
	Fio         string
	Email       string
	PhoneNumber string
}

func main() {
	//html template path
	templatePath := "./templates/sample.html"
	//path for download pdf

	m := http.NewServeMux()
	m.HandleFunc("/outlist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	})

	mux := http.NewServeMux()
	mux.Handle("/", m)
	mux.HandleFunc("/pdf", func(w http.ResponseWriter, r *http.Request) {
		req := pdfGenerator.NewRequestPdf("")

		//html template data
		templateData := OutlistTemplate{
			Title: "Обходной лист",
			Employee: EmployeeTemplate{
				Fio:         "Иванов Иван Иванович",
				PhoneNumber: "+79998887766",
				OrgStruct:   "Отдел безопасности",
				TabelNumber: "123456",
				Position:    "Ответсвенный по безопасности",
				Outdate:     "12.08.2099",
			},
			Points: []PointTemplate{
				{
					Number:      1,
					Title:       "Сдача средств защиты и спецодежды",
					Description: "",
					Approvers: []ApproverTemplate{
						ApproverTemplate{
							Fio:         "Иванов Иван Иванович",
							Email:       "ivanov@mts.ru",
							PhoneNumber: "+79999999999",
						},
						ApproverTemplate{
							Fio:         "Иванов Иван Иванович",
							Email:       "ivanov@mts.ru",
							PhoneNumber: "+79999999999",
						},
						ApproverTemplate{
							Fio:         "Иванов Иван Иванович",
							Email:       "ivanov@mts.ru",
							PhoneNumber: "+79999999999",
						},
					},
					ApprovedDate: "",
					Comment:      "",
				},
			},
		}

		err := req.ParseTemplate(templatePath, templateData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		bytes, err := req.GeneratePDF()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})
	fmt.Println("Listen on 5000")
	http.ListenAndServe(":5000", mux)
}
