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
var defenders = make([]string, 0)
var all = make([]string, 0)

func main() {

	args := os.Args[1:]

	teams, _ := strconv.Atoi(args[0])
	keepers, _ := strconv.Atoi(args[1])
	dfs, _ := strconv.Atoi(args[2])
	mids, _ := strconv.Atoi(args[3])
	fws, _ := strconv.Atoi(args[4])
	randomPlayers, _ := strconv.Atoi(args[5])

	// setup reader
	csvIn, err := os.Open("./playerlist.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	// handle header
	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	positionIndex := 0
	nameIndex := 0

	for i := 0; i < len(header); i++ {
		if header[i] == "position" {
			positionIndex = i
		}

		if header[i] == "name" {
			nameIndex = i
		}
	}
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if rec[positionIndex] == "Midfielder" {
			midfielders = append(midfielders, rec[nameIndex])
		}
		if rec[positionIndex] == "Forward" {
			strikers = append(strikers, rec[nameIndex])
		}
		if rec[positionIndex] == "Goalkeeper" {
			goalkeepers = append(goalkeepers, rec[nameIndex])
		}
		if rec[positionIndex] == "Defender" {
			defenders = append(defenders, rec[nameIndex])
		}

		all = append(all, rec[nameIndex])

	}

	for i := 0; i < teams; i++ {
		team := generateTeam(fws, mids, dfs, keepers, randomPlayers)

		fmt.Printf("Team " + strconv.Itoa(i+1) + "\n")
		for j := 0; j < len(team); j++ {
			fmt.Printf(team[j] + " " + "\n")
		}
		fmt.Printf("\n")
	}
}

func generateTeam(fws int, mids int, dfs int, keepers int, randomPlayers int) []string {

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

	for i := 0; i < dfs; i++ {
		index := rand.Intn(len(defenders))
		team = append(team, defenders[index])
		defenders = append(defenders[:index], defenders[index+1:]...)
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
