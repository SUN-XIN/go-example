package main

import (
	"encoding/json"
	"log"
)

type Toto struct {
	Vint    int
	Vstring string
}

type Tutu struct {
	Vint    int    `json:"v_int"`
	Vstring string `json:"v_string"`
}

type Tata struct {
	Vok         int `json:"vint"`
	vOverlooked int `json:"v_overlooked"`
}

func main() {
	// Marshal without/with specifies json name
	to := Toto{
		Vint:    1,
		Vstring: "a",
	}
	b, err := json.Marshal(to)
	if err != nil {
		log.Printf("Failed Marshal (%+v): %+v", to, err)
		return
	}
	log.Printf("Marshal without specifies json name -> (%s)", b)

	tu := Tutu{
		Vint:    1,
		Vstring: "a",
	}
	b, err = json.Marshal(tu)
	if err != nil {
		log.Printf("Failed Marshal (%+v): %+v", tu, err)
		return
	}
	log.Printf("Marshal with specifies json name -> (%s)", b)

	// Unmarshal into struct without specifies json name
	var totoTest Toto
	err = json.Unmarshal(b, &totoTest)
	if err != nil {
		log.Printf("Failed Unmarshal: %+v", err)
		return
	}
	log.Printf("Unmarshal into the struct without specifies json name -> (%+v)", totoTest)

	// Marshal/Unmarshal with overlooked champ
	ta := Tata{
		Vok:         1,
		vOverlooked: 11,
	}
	b, err = json.Marshal(ta)
	if err != nil {
		log.Printf("Failed Marshal (%+v): %+v", ta, err)
		return
	}
	log.Printf("Marshal with with overlooked champ -> (%s)", b)

	b = []byte(`{"vint":2,"v_overlooked":22}`)
	var tataTest Tata
	err = json.Unmarshal(b, &tataTest)
	if err != nil {
		log.Printf("Failed Unmarshal: %+v", err)
		return
	}
	log.Printf("Unmarshal with with overlooked champ -> (%+v)", tataTest)

	// Unmarshal
	var tataPointer *Tata
	err = json.Unmarshal(b, tataPointer)
	if err != nil {
		log.Printf("Failed Unmarshal: %+v", err)
		return
	}
}
