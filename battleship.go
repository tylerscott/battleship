package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//Makes a new board with no ships
//Returns a 2D array
func Newboard() [10][10]string {
	var a [10][10]string
	a[0][0] = " "
	a[0][1] = "A"
	a[0][2] = "B"
	a[0][3] = "C"
	a[0][4] = "D"
	a[0][5] = "E"
	a[0][6] = "F"
	a[0][7] = "G"
	a[0][8] = "H"
	a[0][9] = "I"

	for i := 1; i < 10; i++ {
		a[i][0] = strconv.Itoa(i)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			a[i][j] = "-"
		}
	}
	return a
}

//Returns true if match/hit/already played spot
//Return false if no match/no hit
func CheckIfMatch(c []string, board [10][10]string) bool {
	var colum int
	//Convert Letter to int
	switch c[0] {
	case "A":
		colum = 1
	case "B":
		colum = 2
	case "C":
		colum = 3
	case "D":
		colum = 4
	case "E":
		colum = 5
	case "F":
		colum = 6
	case "G":
		colum = 7
	case "H":
		colum = 8
	case "I":
		colum = 9
	}

	//convert string to int
	row, err := strconv.Atoi(c[1])
	if err != nil {
	}

	switch board[row][colum] {
	case "X":
		return true
	case "*":
		return true
	}
	return false
}

func CheckForWinner(gameboard [10][10]string, hidden [10][10]string) bool {
	hits := 0
	ships := 0

	for i := 1; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if gameboard[i][j] == "X" {
				hits = hits + 1
			}
			//fmt.Printf("hits = %d\n", hits)
		}
	}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if hidden[i][j] == "X" {
				ships = ships + 1
			}
			//fmt.Printf("ships = %d\n", ships)
		}
	}

	if hits == ships {
		return true
	}
	return false
}

//Draws any board on screen
func Drawboard(inarray [10][10]string) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf(inarray[i][j])
		}
		fmt.Printf("\n")
	}
}

//Places a ship on a board
// static now. will be randomized
func Placeships() [10][10]string {
	a := Newboard()
	var row int
	var colum int
	var dir int
	var restart bool
	ships := 0

	//From stackoverflow to fix same patern with rand.Intn()
	rand.Seed(time.Now().UTC().UnixNano())

	for {

		restart = false
		row = rand.Intn(9) + 1
		colum = rand.Intn(9) + 1
		dir = rand.Intn(20)

		switch {
		case dir <= 4: // down
			if row <= 7 {
				for i := 0; i < 3; i++ { //checks if a ship is already there.
					if a[row+i][colum] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ { //places ship
					a[row+i][colum] = "X"
				}
			} else {
				for i := 0; i < 3; i++ { //checks if open
					if a[row-i][colum] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ { //places ship
					a[row-i][colum] = "X"
				}
			}
		case dir > 4 && dir <= 9: // right
			if colum <= 7 {
				for i := 0; i < 3; i++ {
					if a[row][colum+i] == "X" { // checks if open
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ { // places ship
					a[row][colum+i] = "X"
				}
			} else {
				for i := 0; i < 3; i++ {
					if a[row][colum-i] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ {
					a[row][colum-i] = "X"
				}
			}
		case dir > 9 && dir <= 14: // up
			if row >= 3 {
				for i := 0; i < 3; i++ {
					if a[row-i][colum] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ {
					a[row-i][colum] = "X"
				}
			} else {
				for i := 0; i < 3; i++ {
					if a[row+i][colum] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ {
					a[row+i][colum] = "X"
				}
			}
		case dir > 14: // left
			if colum >= 3 {
				for i := 0; i < 3; i++ {
					if a[row][colum-i] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ {
					a[row][colum-i] = "X"
				}
			} else {
				for i := 0; i < 3; i++ {
					if a[row][colum+i] == "X" {
						restart = true
					}
				}
				if restart {
					continue
				}
				for i := 0; i < 3; i++ {
					a[row][colum+i] = "X"
				}
			}
		}
		ships = ships + 1
		if ships == 4 {
			break
		}
	}

	return a
}

//checks if attack is in range of the board
func ValidAttack(c []string) bool {
	d, err := strconv.Atoi(c[1])
	if err != nil {
	}

	switch c[0] {
	case "A", "B", "C", "D", "E", "F", "G", "H", "I":
		if d >= 1 && d <= 9 {
			return true
		}
	}
	return false
}

//Player makes attack and a new gameboard is returned with the attack on it
func Playerattack(gameboard [10][10]string, hidden [10][10]string) [10][10]string {
	a := gameboard
	var b string
	var c []string
	//Drawboard(hidden)
	fmt.Println("Enter a space to attack. Range A1 - E5")
	//Checks if valid spot on board
	//If valid continues, Else prompts for new attack
	for {
		n, err := fmt.Scanf("%s\n", &b)
		if err != nil {
			fmt.Println(n, err)
		}
		c = strings.Split(strings.ToUpper(b), "")
		// if only 1 char is entered loop contiues loop from begining
		if len(c) < 2 {
			fmt.Println("You only entered part of a space.")
			fmt.Println("Please enter a valid space..")
			continue
		}
		// if already played the spot continues loop
		// breaks loop if new calid spot is played
		if ValidAttack(c) == true {
			if CheckIfMatch(c, a) {
				fmt.Println("You have already played there. Try again.")
				continue
			}
			break
		}
		fmt.Println("Enter a valid space to attack...")
	}

	// after valid attack, marks the gamesboard
	//
	row, err := strconv.Atoi(c[1])
	if err != nil {
	}
	var colum int
	//Convert Letter to int
	switch c[0] {
	case "A":
		colum = 1
	case "B":
		colum = 2
	case "C":
		colum = 3
	case "D":
		colum = 4
	case "E":
		colum = 5
	case "F":
		colum = 6
	case "G":
		colum = 7
	case "H":
		colum = 8
	case "I":
		colum = 9
	}

	// X if hit, * if miss
	switch CheckIfMatch(c, hidden) {
	case true:
		a[row][colum] = "X"
	case false:
		a[row][colum] = "*"
	}

	return a
}

//Calling clears the terminal screen
func clearscreen() {
	var clearScreenCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		clearScreenCmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		clearScreenCmd = exec.Command("clear")
	} else {
		fmt.Fprintf(os.Stderr, "Your OS \"%s\" is not supported.", runtime.GOOS)
		os.Exit(2)
	}

	clearScreenCmd.Stdout = os.Stdout
	clearScreenCmd.Run()
}

func main() {
	fmt.Printf("Welcome to the game!\n")
	board := Newboard()
	key := Placeships()
	//testing Placeships()
	/*
	   Drawboard(key)
	   for i := 0; i < 10; i++ {
	     key = Placeships()
	     Drawboard(key)
	   }
	*/

	// Game play
	for {
		Drawboard(board)
		board = Playerattack(board, key)

		//clear screen here
		clearscreen()

		winner := CheckForWinner(board, key)
		if winner {
			fmt.Println("Congrats, you have won!!!")
			break
		}
	}

}
