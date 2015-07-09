package main

import  (
  "fmt"
  "math/rand"
  "os"
  "os/exec"
  "runtime"
  "strings"
  "strconv"
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

//Checks if two spots match
//Returns 1 if player hits, 2 if comp hits, -1 if already played, 0 if miss
func CheckIfMatch(c [2]int, board [10][10]string ) int {
  row := c[0]
  colum := c[1]

  switch board[row][colum] {
  case "X":
    return 1
  case "*":
    return -1
  case "O":
    return 2
  }
  return 0
}

// 1 if player wins, -1 if com wins, 0 if no winner yet
func CheckForWinner(gameboard [10][10]string,hboard [10][10]string, hidden [10][10]string) int {
  hits := 0
  ships := 0
  chits := 0

  for i := 1; i < 10; i++ {
    for j := 0; j < 10; j++ {
      if gameboard[i][j] == "X"{
        hits = hits + 1
      }
      //fmt.Printf("hits = %d\n", hits)
    }
  }
  for i := 1; i < 10; i++ {
    for j := 0; j < 10; j++ {
      if hboard[i][j] == "X"{
        chits = chits + 1
      }
      //fmt.Printf("hits = %d\n", hits)
    }
  }
  for i := 1; i < 10; i++ {
    for j := 1; j < 10; j++ {
      if hidden[i][j] == "X"{
        ships = ships + 1
      }
      //fmt.Printf("ships = %d\n", ships)
    }
  }

  if hits == ships {
    return 1
  }
  if chits == ships {
    return -1
  }
  return 0
}


func CompAttack(compcard [10][10]string, lhit [4]int) ([10][10]string, [4]int)  {
  a := compcard
  var row int
  var colum int
  //var restart bool
  var c [2]int
  lasthit := lhit

  rand.Seed( time.Now().UTC().UnixNano())
  for {
    // if it hasnt hit yet
    if lasthit[0] == 0 {
      for {
        row = rand.Intn(9) + 1
        colum = rand.Intn(9) + 1

        c[0] = row
        c[1] = colum
        switch CheckIfMatch(c, a) {
        case -1:
          continue
        case 0:
          a[row][colum] = "*"
          return a, lasthit
        case 1:
          continue
        case 2:
          fmt.Println("You've been hit!")
          a[row][colum] = "X"
          lasthit[0] = row
          lasthit[1] = colum
          return a, lasthit
        }

      }
    }
    // has 1 hit
    if lasthit[2] == 0 && lasthit[3] == 0 {
      for {
        c[0] = lhit[0]
        c[1] = lhit[1]
        row = c[0]
        colum = c[1]

        r := rand.Intn(39) + 4
        switch r % 4 {
        case 0:
          if lasthit[0] != 1 {
            c[0] = c[0] - 1
            switch CheckIfMatch(c, a) {
            case 1, -1:
              continue
            case 0:
              row = c[0]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              row = c[0]
              a[row][colum] = "X"
              lasthit[2] = 1
              lasthit[0] = row
              lasthit[1] = colum
              lasthit[3] = lasthit[3] + 1
              return a, lasthit
            }
          }
          continue
        case 1:
          if lasthit[0] != 9 {
            c[0] = c[0] + 1
            switch CheckIfMatch(c, a) {
            case 1, -1:
              continue
            case 0:
              row = c[0]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              row = c[0]
              a[row][colum] = "X"
              lasthit[2] = 2
              lasthit[0] = row
              lasthit[1] = colum
              lasthit[3] = lasthit[3] + 1
              return a, lasthit
            }
          }
          continue
        case 2:
          if lasthit[1] != 1 {
            c[1] = c[1] - 1
            switch CheckIfMatch(c, a) {
            case 1, -1:
              continue
            case 0:
              colum = c[1]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              colum = c[1]
              a[row][colum] = "X"
              lasthit[2] = 3
              lasthit[0] = row
              lasthit[1] = colum
              lasthit[3] = lasthit[3] + 1
              return a, lasthit
            }
          }
          continue
        case 3:
          if lasthit[1] != 9 {
            c[1] = c[1] + 1
            switch CheckIfMatch(c, a) {
            case 1, -1:
              continue
            case 0:
              colum = c[1]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              colum = c[1]
              a[row][colum] = "X"
              lasthit[2] = 4
              lasthit[0] = row
              lasthit[1] = colum
              lasthit[3] = lasthit[3] + 1
              return a, lasthit
            }
          }
          continue
        }
      }
    }

    if lasthit[2] > 0 && lasthit[3] < 2 {
      for {
        c[0] = lasthit[0]
        c[1] = lasthit[1]
        row = c[0]
        colum = c[1]

        switch lasthit[2] {
        case 1:
          if lasthit[0] > 1 {
            c[0] = row - 1
            switch CheckIfMatch(c, a) {
            case -1, 1:
              c[0] = row + 2
              if c[0] <= 9 {
                switch CheckIfMatch(c, a) {
                case -1, 1:
                  lasthit[2] = 0
                  lasthit[3] = 0
                  break
                case 0:
                  row = c[0]
                  a[row][colum] = "*"
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                case 2:
                  row = c[0]
                  a[row][colum] = "X"
                  lasthit[0] = 0
                  lasthit[1] = 0
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                }
              }
              lasthit[2] = 0
              lasthit[3] = 0
              break
            case 0:
              row = c[0]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              row = c[0]
              a[row][colum] = "X"
              lasthit[0] = 0
              lasthit[1] = 0
              lasthit[2] = 0
              lasthit[3] = 0
              return a, lasthit

            }
          }
          row = c[0] + 1
          lasthit[0] = row
          lasthit[2] = 2
          continue
        //if last time was down
        case 2:
          if lasthit[0] < 9 {
            c[0] = row + 1
            switch CheckIfMatch(c, a) {
            case -1, 1:
              c[0] = row - 2
              if c[0] >= 1 {
                switch CheckIfMatch(c, a) {
                case -1, 1:
                  lasthit[2] = 0
                  lasthit[3] = 0
                  break
                case 0:
                  row = c[0]
                  a[row][colum] = "*"
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                case 2:
                  row = c[0]
                  a[row][colum] = "X"
                  lasthit[0] = 0
                  lasthit[1] = 0
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                }
              }
              lasthit[2] = 0
              lasthit[3] = 0
              break
            case 0:
              row = c[0]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              row = c[0]
              a[row][colum] = "X"
              lasthit[0] = 0
              lasthit[1] = 0
              lasthit[2] = 0
              lasthit[3] = 0
              return a, lasthit

            }
          }
          row = c[0] - 1
          lasthit[0] = row
          lasthit[2] = 1
          continue
        case 3:
          if lasthit[1] > 1 {
            c[1] = colum - 1
            switch CheckIfMatch(c, a) {
            case -1, 1:
              c[1] = colum + 2
              if c[1] <= 9 {
                switch CheckIfMatch(c, a) {
                case -1, 1:
                  lasthit[2] = 0
                  lasthit[3] = 0
                  break
                case 0:
                  colum = c[1]
                  a[row][colum] = "*"
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                case 2:
                  colum = c[1]
                  a[row][colum] = "X"
                  lasthit[0] = 0
                  lasthit[1] = 0
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                }
              }
              lasthit[2] = 0
              lasthit[3] = 0
              break
            case 0:
              colum = c[1]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              colum = c[1]
              a[row][colum] = "X"
              lasthit[0] = 0
              lasthit[1] = 0
              lasthit[2] = 0
              lasthit[3] = 0
              return a, lasthit

            }
          }
          colum = c[1] + 1
          lasthit[1] = colum
          lasthit[2] = 4
          continue
        case 4:
          if lasthit[1] < 9 {
            c[1] = colum + 1
            switch CheckIfMatch(c, a) {
            case -1, 1:
              c[1] = colum - 2
              if c[1] >= 1 {
                switch CheckIfMatch(c, a) {
                case -1, 1:
                  lasthit[2] = 0
                  lasthit[3] = 0
                  break
                case 0:
                  colum = c[1]
                  a[row][colum] = "*"
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                case 2:
                  colum = c[1]
                  a[row][colum] = "X"
                  lasthit[0] = 0
                  lasthit[1] = 0
                  lasthit[2] = 0
                  lasthit[3] = 0
                  return a, lasthit
                }
              }
              lasthit[2] = 0
              lasthit[3] = 0
              break
            case 0:
              colum = c[1]
              a[row][colum] = "*"
              return a, lasthit
            case 2:
              colum = c[1]
              a[row][colum] = "X"
              lasthit[0] = 0
              lasthit[1] = 0
              lasthit[2] = 0
              lasthit[3] = 0
              return a, lasthit

            }
          }
          colum = c[1] - 1
          lasthit[1] = colum
          lasthit[2] = 2
          continue


        } //lasthit switch

      }
      continue
    }

  }
  return a, lhit
}


//Draws any board on screen
func Drawboard(inarray [10][10]string)  {
  for i := 0; i < 10; i++ {
    for j := 0; j < 10; j++ {
      fmt.Printf(inarray[i][j])
    }
    fmt.Printf("\n")
  }
}

// Allows user to enter a space and verifies its validity
func EnterSpace(a [10][10]string) [2]int {
  var b string
  var c []string
  var out [2]int

  for  {
    n, err := fmt.Scanf("%s\n", &b)
      if err != nil {
        fmt.Println(n, err)
    }
    c = strings.Split(strings.ToUpper(b), "")
    // if only 1 char is entered loop contiues loop from begining
    if len(c) < 2 {
      fmt.Println("You only entered part of a space.")
      fmt.Println("Please enter a valid space...")
      continue
    }
    // breaks loop if new calid spot is played
    if ValidSpace(c) == true {
      break
    }
    fmt.Println("Enter a valid space...")
  }

  // converts strings to int
  row, err := strconv.Atoi(c[1]);
  if err != nil {}
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

  out[1] = colum
  out[0] = row

  return out
}

// gets direction for ship, and verifies it fits on the board
// returns 1 for up, 2 for down, 3 for left, 4 for right
func GetDir(spot [2]int) int {
  var dirstring string
  var c []string

  for {
    n, err := fmt.Scanf("%s\n", &dirstring)
      if err != nil {
        fmt.Println(n, err)
    }
    c = strings.Split(strings.ToUpper(dirstring), "")
    switch c[0] {
    case "U":
      if (spot[0] - 2) <= 0 {
        fmt.Println("That ship doesn't fit there, try again.")
        continue
      }
      return 1
    case "D":
      if (spot[0] + 2) >= 10 {
        fmt.Println("That ship doesn't fit there, try again.")
        continue
      }
      return 2
    case "L":
      if (spot[1] - 2) <= 0 {
        fmt.Println("That ship doesn't fit there, try again.")
        continue
      }
      return 3
    case "R":
      if (spot[1] + 2) >= 10 {
        fmt.Println("That ship doesn't fit there, try again.")
        continue
      }
      return 4
    }
    fmt.Println("Please enter up, down, left, or right")
  }
}

//Places a ship on a board
// static now. will be randomized
func Placeships(s int) [10][10]string {
  a := Newboard()
  var row int
  var colum int
  var dir int
  var restart bool
  ships := 0

  //From stackoverflow to fix same patern with rand.Intn()
  rand.Seed( time.Now().UTC().UnixNano())

  for {

    restart = false
    row = rand.Intn(9) + 1
    colum = rand.Intn(9) + 1
    dir = rand.Intn(20)

    switch {
    case dir <= 4: // down
      if row <= 7 {
        for i := 0; i < 3; i++ { //checks if a ship is already there.
          if a[row + i][colum] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ { //places ship
          a[row + i][colum] = "X"
        }
      } else {
        for i := 0; i < 3; i++ { //checks if open
          if a[row - i][colum] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ { //places ship
          a[row - i][colum] = "X"
        }
      }
    case dir > 4 && dir <= 9: // right
      if colum <= 7 {
        for i := 0; i < 3; i++ {
          if a[row][colum + i] == "X" { // checks if open
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ { // places ship
          a[row][colum + i] = "X"
        }
      } else {
        for i := 0; i < 3; i++ {
          if a[row][colum - i] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ {
          a[row][colum - i] = "X"
        }
      }
    case dir > 9 && dir <= 14: // up
      if row >= 3 {
        for i := 0; i < 3; i++ {
          if a[row - i][colum] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ {
          a[row - i][colum] = "X"
        }
      } else {
        for i := 0; i < 3; i++ {
          if a[row + i][colum] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ {
          a[row + i][colum] = "X"
        }
      }
    case dir > 14: // left
      if colum >=3 {
        for i := 0; i < 3; i++ {
          if a[row][colum - i] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ {
          a[row][colum - i] = "X"
        }
      } else {
        for i := 0; i < 3; i++ {
          if a[row][colum + i] == "X" {
            restart = true
          }
        }
        if restart {
          continue
        }
        for i := 0; i < 3; i++ {
          a[row][colum + i] = "X"
        }
      }
    }
    ships = ships + 1
    if ships == s {
      break
    }
  }

  return a
}

//checks if attack is in range of the board
func ValidSpace(c []string) bool {
  d, err := strconv.Atoi(c[1]);
  if err != nil {}

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
  var spot [2]int
  //Drawboard(hidden)
  fmt.Println("Enter a space to attack. Range A1 - E5")
  //Checks if valid spot on board
  //If valid continues, Else prompts for new attack
  for  {
    spot = EnterSpace(a)
    //Drawboard(a)
    if CheckIfMatch(spot, a) == 1 || CheckIfMatch(spot, a) == -1 {
        fmt.Println("You have already played there. Try again.")
        continue
    }
    break
  }

  row := spot[0]
  colum := spot[1]
  // X if hit, * if miss
  switch CheckIfMatch(spot, hidden) {
  case 1:
    a[row][colum] = "X"
  case 0:
    a[row][colum] = "*"
  }

  return a
}


func PlayerBoard(s int) [10][10]string  {
  a := Newboard()
  ships := 0
  var row int
  var colum int
  var spot [2]int
  var dir int
  var restart bool
  Drawboard(a)

  for {
    restart = false
    fmt.Println("Enter a space to place your ship. ie. b4 ")
    spot = EnterSpace(a)
    fmt.Println("Select a direction to place your ship. up, down, left, right\nYour ship will take up three spaces starting")
    dir = GetDir(spot)
    row = spot[0]
    colum = spot[1]
    //test if a ship is already then, if not a new ship will be placed
    switch dir {
    case 1:
      for i := 0; i < 3; i++ {
        if a[row - i][colum] == "O" {
          restart = true
        }
      }
      if restart {
        fmt.Println("A ship was already played in one of those spaces.\nTry a new space.")
        continue
      }
      for i := 0; i < 3; i++ {
        a[row - i][colum] = "O"
      }
    case 2:
      for i := 0; i < 3; i++ {
        if a[row + i][colum] == "O" {
          restart = true
        }
      }
      if restart {
        fmt.Println("A ship was already played in one of those spaces.\nTry a new space.")
        continue
      }
      for i := 0; i < 3; i++ {
        a[row + i][colum] = "O"
      }
    case 3:
      for i := 0; i < 3; i++ {
        if a[row][colum - i] == "O" {
          restart = true
        }
      }
      if restart {
        fmt.Println("A ship was already played in one of those spaces.\nTry a new space.")
        continue
      }
      for i := 0; i < 3; i++ {
        a[row][colum - i] = "O"
      }
    case 4:
      for i := 0; i < 3; i++ {
        if a[row][colum + i] == "O" {
          restart = true
        }
      }
      if restart {
        fmt.Println("A ship was already played in one of those spaces.\nTry a new space.")
        continue
      }
      for i := 0; i < 3; i++ {
        a[row][colum + i] = "O"
      }
    }
    ships = ships + 1
    if ships == s {
      break
    }
    clearscreen()
    fmt.Printf("You have placed %d of 3 ships\n", ships)
    Drawboard(a)
    fmt.Println("Place your next ship.")
  }
  return a
}


//Calling clears the terminal screen
// Windows cl := exec.Command("cmd", "/c", "cls")
// unix base cl := exec.Command("clear")
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
  ships := 6
  var complogic [4]int
  fmt.Printf("Welcome to the game!\n")
  board := Newboard()
  key := Placeships(ships)
  humankey := PlayerBoard(ships)

  // Game play
  for {
    fmt.Println("Attack board\n")
    Drawboard(board)
    fmt.Println("\nYour Ships\n")
    Drawboard(humankey)
    board = Playerattack(board, key)
    humankey, complogic = CompAttack(humankey, complogic)

    //clear screen here
    clearscreen()

    if CheckForWinner(board, humankey, key) == 1 {
      fmt.Println("Congrats, you've won!")
      break
    } else if CheckForWinner(board, humankey, key) == -1 {
      fmt.Println("I'm sorry, you have lost.")
      break
    }
  }

}
