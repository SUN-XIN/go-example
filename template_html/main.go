package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func handlerSimple(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("toto").Parse(htmlSimple)
	if err != nil {
		log.Printf("failed template Parse: %+v", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("failed Execute: %+v", err)
		return
	}
}

type Person struct {
	Name   string
	City   string
	Number int
}

type Persons struct {
	Data       []Person
	HideNumber int
}

func handlerWithParams(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	city := r.FormValue("city")

	p := Person{
		Name: name,
		City: city,
	}
	tmpl, err := template.New("tutu").Parse(htmlParams)
	if err != nil {
		log.Printf("failed template Parse: %+v", err)
		return
	}

	err = tmpl.Execute(w, &p)
	if err != nil {
		log.Printf("failed Execute: %+v", err)
		return
	}
}

func handlerList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	city := r.FormValue("city")

	ps := make([]Person, 0, 5)
	for i := 0; i < 5; i++ {
		ps = append(ps, Person{
			Name:   fmt.Sprintf("%s-%d", name, i),
			City:   fmt.Sprintf("%s-%d", city, i),
			Number: i,
		})
	}

	tmpl, err := template.New("tata").Parse(htmlList)
	if err != nil {
		log.Printf("failed template Parse: %+v", err)
		return
	}

	err = tmpl.Execute(w, &Persons{Data: ps})
	if err != nil {
		log.Printf("failed Execute: %+v", err)
		return
	}
}

func handlerCheckList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	city := r.FormValue("city")
	hideNum := r.FormValue("h_num")

	var n int
	var err error
	if hideNum == "" {
		n = 0
	} else {
		n, err = strconv.Atoi(hideNum)
		if err != nil {
			log.Printf("failed Atoi (%s): %+v", hideNum, err)
		}
	}

	ps := make([]Person, 0, 5)
	for i := 0; i < 5; i++ {
		ps = append(ps, Person{
			Name:   fmt.Sprintf("%s-%d", name, i),
			City:   fmt.Sprintf("%s-%d", city, i),
			Number: i,
		})
	}

	funcMod := template.FuncMap{"mod": func(a, b int) int {
		return a % b
	}}

	tmpl, err := template.New("titi").Funcs(funcMod).Parse(htmlCheckList)
	if err != nil {
		log.Printf("failed template Parse: %+v", err)
		return
	}

	err = tmpl.Execute(w, &Persons{
		Data:       ps,
		HideNumber: n})
	if err != nil {
		log.Printf("failed Execute: %+v", err)
		return
	}
}

func main() {
	http.HandleFunc("/simple", handlerSimple)
	http.HandleFunc("/params", handlerWithParams)
	http.HandleFunc("/list", handlerList)
	http.HandleFunc("/check_and_list", handlerCheckList)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

const (
	htmlSimple = `
<html>
    <body>
    This is a simple example
    </body>
</html>
    `

	htmlParams = `
<html>
    <body>
    This is an example with Params <br/>
    Name: {{.Name}} <br/>
    City: {{.City}} <br/>
    </body>
</html>
    `

	htmlList = `
<html>
    <body>
    This is an example with List <br/> 
    <ul>
    {{range .Data}} 
        <li>{{.Name}}</li>
        <li>{{.City}}</li> 
        <li>{{.Number}}</li>
        <br/>    
    {{end}}
</ul>
</html>
    `

	htmlCheckList = `
<html>
    <body>
    {{$hide_num := .HideNumber}} 
    This is an example with List and no display if Number equals to HideNumber ({{$hide_num}}) <br/> 
    <ul>
    {{range .Data}} 
        {{if eq .Number $hide_num}} 
            <li>equals to HideNumber->do not display</li>
            <li>Number: {{.Number}}</li>
        {{else}}
            {{$res := mod .Number 2}}
            {{if eq $res 0}}
                <li>paired number->display city</li>
                <li>{{.City}}</li> 
                <li>Number: {{.Number}}</li>
            {{else}}
                <li>unpaired number->display name</li>
                <li>{{.Name}}</li> 
                <li>Number: {{.Number}}</li>
            {{end}}
        {{end}}
        <br/>    
    {{end}}
</ul>
</html>
    `
)
