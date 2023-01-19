package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		resp := GetResp{Apiversion: "1"}
		err := json.NewEncoder(rw).Encode(resp)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/start", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		log.Println("joined game")
	})
	http.HandleFunc("/move", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		resp := GameResp{
			Move: Direction(rand.Intn(4)).String(),
		}
		log.Printf("going %s", resp.Move)
		err := json.NewEncoder(rw).Encode(resp)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/end", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		log.Println("game end")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

type GetResp struct {
	Apiversion string `json:"apiversion"`
	Author     string `json:"author,omitempty"`
	Color      string `json:"color,omitempty"`
	Head       string `json:"head,omitempty"`
	Tail       string `json:"tail,omitempty"`
	Version    string `json:"version,omitempty"`
}

type GameReq struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type Direction int

func (d Direction) String() string {
	switch d {
	case DirUp:
		return "up"
	case DirDown:
		return "down"
	case DirLeft:
		return "left"
	default:
		return "right"
	}
}

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

type GameResp struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

type Game struct {
	ID      string `json:"id"`
	Ruleset struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"ruleset"`
	Map     string `json:"map"`
	Timeout int    `json:"timeout"`
	Source  string `json:"source"`
}
type Battlesnake struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Health         int           `json:"health"`
	Body           []Coordinates `json:"body"`
	Latency        string        `json:"latency"`
	Head           Coordinates   `json:"head"`
	Length         int           `json:"length"`
	Shout          string        `json:"shout"`
	Squad          string        `json:"squad"`
	Customizations struct {
		Color string `json:"color"`
		Head  string `json:"head"`
		Tail  string `json:"tail"`
	} `json:"customizations"`
}
type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Board struct {
	Height  int           `json:"height"`
	Width   int           `json:"width"`
	Food    []Coordinates `json:"food"`
	Hazards []Coordinates `json:"hazards"`
	Snakes  []Battlesnake `json:"snakes"`
}