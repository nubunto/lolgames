package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const apiURL = "https://br.api.pvp.net/api/lol/br/v1.3/game/by-summoner/7492820/recent?api_key="

type GameData struct {
	Games      []GameRawData `json:"games"`
	SummonerID int64         `json:"summonerId"`
}

type GameRawData struct {
	Stats Stats `json:"stats"`
}

type Stats struct {
	Assists        int `json:"assists"`
	GoldEarned     int `json:"goldEarned"`
	MinionsKilled  int `json:"minionsKilled"`
	PlayerRole     int `json:"playerRole"`
	PlayerPosition int `json:"playerPosition"`
}

//go:generate stringer -type=PlayerRole
type PlayerRole int

const (
	_ PlayerRole = iota
	Duo
	Support
	Carry
	Solo
)

//go:generate stringer -type=PlayerPosition
type PlayerPosition int

const (
	_ PlayerPosition = iota
	Top
	Middle
	Jungle
	Bot
)

func main() {
	apiURLWithKey := apiURL + os.Getenv("RIOT_KEY")
	fmt.Println("Fazendo requisição a", apiURLWithKey)
	req, err := http.NewRequest("GET", apiURLWithKey, nil)
	if err != nil {
		log.Fatal(err)
	}
	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(res.Body)
	var data GameData
	if err := dec.Decode(&data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jogos:", len(data.Games))
	for _, game := range data.Games {
		fmt.Printf("Assistências: %d\nOuro ganho: %d\nMinions mortos: %d\nPosição: %s\nRole: %s\n\n",
			game.Stats.Assists,
			game.Stats.GoldEarned,
			game.Stats.MinionsKilled,
			PlayerPosition(game.Stats.PlayerPosition).String(),
			PlayerRole(game.Stats.PlayerRole).String())
	}
}
