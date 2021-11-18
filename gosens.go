package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	s "strings"
	"strconv"
	"time"
	"github.com/pterm/pterm"
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
	{"source", "Source Games", 2.0, 0.022},
}

// matching string input to game dictionary
func getSelectedGame(inputGame string) (name string, prec float64, yaw float64) {
	for _, entry := range gameDic {
		if s.ToLower(entry.shortName) == s.ToLower(inputGame) || s.ToLower(entry.fullName) == s.ToLower(inputGame) {
			return entry.fullName, entry.gamePrecision, entry.yaw
		}
	}
	panic("That game does not exist yet")
}

// generating the random cm/360
func floatRange(min int, max int) float64 {
	fmin := float64(min)
	fmax := float64(max)
	rand.Seed(time.Now().UnixNano())
	return fmin + rand.Float64()*(fmax-fmin)
}

// converts the random cm/360 to game format and rounds to the game precision
// TODO: if games allow different precisions in a config file, have it output both
func generateSens(randValue float64, dpi int, yaw float64, prec float64) (sens float64) {
	fdpi := float64(dpi)
	unrounded := (360.0 * 2.54 / (randValue * fdpi * yaw))
	return float64(int(unrounded* math.Pow(10, prec))) / math.Pow(10, prec)
}

/*
Recalculates the cm/360 for the generated sensitivity since 
the generated sensitivity is sometimes not as accurate as 
the cm/360 initially generated.

Because of that, the program will generate a cm/360 value
that isn't possible to set in game, so it takes the rounded
sens generated in generateSens() and recalculates what the
cm/360 is for that value, not the initially generated one.
*/
func recalcCM(sens float64, dpi int, yaw float64) (cm float64) {
	fdpi := float64(dpi)
	return 914.4 / (sens * fdpi * yaw)
}

func main() {

	// input
	inputGame := flag.String("game", "ow", "Game to generate sensitivity for")
	inputDPI := flag.Int("dpi", 800, "Your DPI")
	inputMin := flag.Int("min", 20, "Lower sensitivity bound")
	inputMax := flag.Int("max", 40, "Upper sensitivity bound")
	inputGameList := flag.Bool("games", false, "Print the available games")
	inputDebug := flag.Bool("debug", false, "Print all variables")
	flag.Parse()

	// processing input
	fullName, precision, yaw := getSelectedGame(*inputGame)
	randNum := floatRange(*inputMin, *inputMax)
	genOutput := generateSens(randNum, *inputDPI, yaw, precision)
	cm360 := recalcCM(genOutput, *inputDPI, yaw)

	// output
	if *inputGameList == true {
		fmt.Println("Games Dictionary:\n")
		for _, entry := range gameDic {
			fmt.Printf("%s - %s\n", entry.fullName, entry.shortName)
		}
	} else {
		pterm.DefaultBasicText.Printf(pterm.Green("%v ") + "in %s (" + pterm.Green("%0.2f cm/360") + ")\n\n", genOutput, fullName, cm360)
		pterm.FgLightCyan.Println("Settings:")
	    pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
	        {"DPI", "Minimum", "Maximum"},
	        {strconv.Itoa(*inputDPI), strconv.Itoa(*inputMin), strconv.Itoa(*inputMax)},
	    }).Render()

		//fmt.Printf("%v in %s setttings (%0.2f cm/360)\n", genOutput, fullName, cm360)
		//fmt.Printf("Settings: %d DPI, %dcm - %dcm\n", *inputDPI, *inputMin, *inputMax)
	}

	// debug
	if *inputDebug == true {
		fmt.Println("\nDebug Info:\n-----------")
		fmt.Printf("Input Game: %v\n", *inputGame)
		fmt.Printf("Input DPI: %v\n", *inputDPI)
		fmt.Printf("Input Min: %v\n", *inputMin)
		fmt.Printf("Input Max: %v\n\n", *inputMax)
		fmt.Printf("fullName: %v\n", fullName)
		fmt.Printf("Precision: %v\n", precision)
		fmt.Printf("Yaw: %v\n", yaw)
		fmt.Printf("RandNum: %v\n", randNum)
		fmt.Printf("genOutput: %v\n", genOutput)
		fmt.Printf("cm360: %v\n", cm360)
	}

}
