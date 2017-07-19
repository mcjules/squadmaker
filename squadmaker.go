package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var midfielders = make([]string, 0)
var goalkeepers = make([]string, 0)
var strikers = make([]string, 0)
var all = make([]string, 0)

func main() {

	args := os.Args[1:]

	teams, _ := strconv.Atoi(args[0])
	keepers, _ := strconv.Atoi(args[1])
	mids, _ := strconv.Atoi(args[2])
	fws, _ := strconv.Atoi(args[3])
	randomPlayers, _ := strconv.Atoi(args[4])

	// setup reader
	csvIn, err := os.Open("./premier_league_squads.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	// handle header
	_, err = r.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if rec[1] == "Midfielder" {
			midfielders = append(midfielders, rec[3])
		}
		if rec[1] == "Forward" {
			strikers = append(strikers, rec[3])
		}
		if rec[1] == "Goalkeeper" {
			goalkeepers = append(goalkeepers, rec[3])
		}
		all = append(all, rec[3])

	}

	for i := 0; i < teams; i++ {
		team := generateTeam(fws, mids, keepers, randomPlayers)

		fmt.Printf("Team " + strconv.Itoa(i+1) + "\n")
		for j := 0; j < len(team); j++ {
			fmt.Printf(team[j] + " " + "\n")
		}
		fmt.Printf("\n")
	}
}

func generateTeam(fws int, mids int, keepers int, randomPlayers int) []string {

	rand.Seed(time.Now().Unix())

	team := make([]string, 0)

	for i := 0; i < fws; i++ {
		index := rand.Intn(len(strikers))
		team = append(team, strikers[index])
		strikers = append(strikers[:index], strikers[index+1:]...)
	}

	for i := 0; i < mids; i++ {
		index := rand.Intn(len(midfielders))
		team = append(team, midfielders[index])
		midfielders = append(midfielders[:index], midfielders[index+1:]...)
	}

	for i := 0; i < keepers; i++ {
		index := rand.Intn(len(goalkeepers))
		team = append(team, goalkeepers[index])
		goalkeepers = append(goalkeepers[:index], goalkeepers[index+1:]...)
	}

	for i := 0; i < randomPlayers; i++ {
		index := rand.Intn(len(all))
		team = append(team, all[index])
		all = append(all[:index], all[index+1:]...)
	}

	return team
}
