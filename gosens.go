package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	s "strings"
	"time"
)

type gameEntry struct {
	shortName     string
	fullName      string
	gamePrecision float64
	yaw           float64
}

var gameDic = []gameEntry{
	{"ow", "Overwatch", 2.0, 0.0066},
	{"fn", "Fortnite", 3.0, 0.005555},
	{"csgo", "Counter-Strike", 2.0, 0.022},
	{"qc", "Quake Champions", 6.0, 0.022},
}

func getSelectedGame(inputGame string) (name string, prec float64, yaw float64) {
	for _, entry := range gameDic {
		if s.ToLower(entry.shortName) == s.ToLower(inputGame) || s.ToLower(entry.fullName) == s.ToLower(inputGame) {
			return entry.fullName, entry.gamePrecision, entry.yaw
		}
	}
	panic("That game does not exist yet")
}

func floatRange(min int, max int) float64 {
	fmin := float64(min)
	fmax := float64(max)
	rand.Seed(time.Now().UnixNano())
	return fmin + rand.Float64()*(fmax-fmin)
}

func generateSens(randValue float64, dpi int, yaw float64, precision float64) string {
	fdpi := float64(dpi)
	sens := (4572.0 / (5.0 * randValue * fdpi * yaw))
	return strconv.FormatFloat(sens, 'f', int(precision), 64)
}

func main() {

	// input
	inputGame := flag.String("game", "ow", "Game to generate sensitivity for")
	inputDPI := flag.Int("dpi", 800, "Your DPI")
	inputMin := flag.Int("min", 20, "Lower sensitivity bound")
	inputMax := flag.Int("max", 40, "Upper sensitivity bound")
	flag.Parse()

	// processing input
	fullName, precision, yaw := getSelectedGame(*inputGame)
	randNum := floatRange(*inputMin, *inputMax)
	genOutput := generateSens(randNum, *inputDPI, yaw, precision)

	// output
	fmt.Printf("%s in %s setttings (%0.2f cm/360)\n", genOutput, fullName, randNum)
	fmt.Printf("Settings: %d DPI, %dcm - %dcm", *inputDPI, *inputMin, *inputMax)
}
