package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

const (
	templatesDir = "templates/"
	templateBase = templatesDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

// IndexHandler is the handler for the root path
func Index(w http.ResponseWriter, r *http.Request) {
	resetValues()
	renderTemplate(w, "index.html", nil)
}

// newGame is the handler for the /newgame path
func NewGame(w http.ResponseWriter, r *http.Request) {
	resetValues()
	renderTemplate(w, "new-game.html", nil)
}

// Game is the handler for the /play path
func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			log.Println(err)
			return
		}
		player.Name = r.Form.Get("name")
	}
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusSeeOther)
	}
	renderTemplate(w, "game.html", player)
}

// Play is the handler for the /play path
func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		http.Error(w, "Error al serializar el resultado", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// About is the handler for the /about path
func About(w http.ResponseWriter, r *http.Request) {
	resetValues()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templatesDir+page))
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func resetValues() {
	player.Name = ""
	rps.PlayerScore = 0
	rps.ComputerScore = 0
}
