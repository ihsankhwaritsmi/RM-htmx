package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Character struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Image   string `json:"image"`
}

type CharacterPage struct {
	Results []Character `json:"results"`
}

func main() {
	charApi := "https://rickandmortyapi.com/api/character/?page=5"

	f1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	var allCharacter []Character

	response, err := http.Get(charApi)
	if err != nil {
		fmt.Println("Error making HTTP request", err)
		return
	}
	defer response.Body.Close()

	var characterPage CharacterPage
	err = json.NewDecoder(response.Body).Decode(&characterPage)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	allCharacter = append(allCharacter, characterPage.Results...)

	for _, character := range allCharacter {
		fmt.Println(character.Name)
		fmt.Println(character.Image)

	}
	f2 := func(w http.ResponseWriter, r *http.Request) {

		for _, character := range allCharacter {
			htmlStr := fmt.Sprintf(` <div class="card p-3 m-sm-5" style="width: 18rem;">
										<img src="%s" class="card-img-top">
										<div class="card-body">
											<p class="card-text">Name		: %s</p>
											<p class="card-text">Status		: %s</p>
											<p class="card-text">Species	:%s</p>
										</div>
									</div>`, character.Image, character.Name, character.Status, character.Species)
			tmpl, _ := template.New("t").Parse(htmlStr)
			tmpl.Execute(w, nil)
		}
	}

	http.HandleFunc("/", f1)
	http.HandleFunc("/first-act", f2)
	log.Fatal(http.ListenAndServe(":7000", nil))
}
