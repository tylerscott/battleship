package main

import  (
  "fmt"
  "math/rand"
  "os"
  "os/exec"
  "strings"
  "strconv"
)

//Makes a new board with no ships
//Returns a 2D array
func Newboard() [6][6]string {
  var a [6][6]string
  a[0][0] = " "
  a[0][1] = "A"
  a[0][2] = "B"
  a[0][3] = "C"
  a[0][4] = "D"
  a[0][5] = "E"


  for i := 1; i < 6; i++ {
    a[i][0] = strconv.Itoa(i)
  }

  for i := 1; i < 6; i++ {
    for j := 1; j < 6; j++ {
      a[i][j] = "-"
    }
  }
  return a
}

//Returns true if match/hit/already played spot
//Return false if no match/no hit
func CheckIfMatch(c []string, board [6][6]string ) bool {
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
  }

  //convert string to int
  row, err := strconv.Atoi(c[1]);
  if err != nil {}

  switch board[row][colum] {
  case "X":
    return true
  case "*":
    return true
  }
  return false
}

func CheckForWinner(gameboard [6][6]string, hidden [6][6]string) bool {
  hits := 0
  ships := 0

  for i := 1; i < 6; i++ {
    for j := 0; j < 6; j++ {
      if gameboard[i][j] == "X"{
        hits = hits + 1
      }
      //fmt.Printf("hits = %d\n", hits)
    }
  }
  for i := 1; i < 6; i++ {
    for j := 1; j < 6; j++ {
      if hidden[i][j] == "X"{
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
func Drawboard(inarray [6][6]string)  {
  for i := 0; i < 6; i++ {
    for j := 0; j < 6; j++ {
      fmt.Printf(inarray[i][j])
    }
    fmt.Printf("\n")
  }
}

//Places a ship on a board
// static now. will be randomized
func Placeships() [6][6]string {
  a := Newboard()
  var row int
  var colum int
  var dir int

  row = rand.Intn(5) + 1
  colum = rand.Intn(5) + 1
  dir = rand.Intn(20)

  switch {
  case dir <= 4: // down
    if row <= 4 {
      //fmt.Println("..down..")
      for i := 0; i < 3; i++ {
        a[row + i][colum] = "X"
      }
    } else {
      //fmt.Println("..down.up..")
      for i := 0; i < 3; i++ {
        a[row - i][colum] = "X"
      }
    }
  case dir > 4 && dir <= 9: // right
    if colum <= 3 {
      //fmt.Println("..right..")
      for i := 0; i < 3; i++ {
        a[row][colum + i] = "X"
      }
    } else {
      //fmt.Println("..right.left..")
      for i := 0; i < 3; i++ {
        a[row][colum - i] = "X"
      }
    }
  case dir > 9 && dir <= 14: // up
    if row >= 3 {
      //fmt.Println("..up..")
      for i := 0; i < 3; i++ {
        a[row - i][colum] = "X"
      }
    } else {
      //fmt.Println("..up.down..")
      for i := 0; i < 3; i++ {
        a[row + i][colum] = "X"
      }
    }
  case dir > 14: // left
    if colum >=3 {
      //fmt.Println("..left..")
      for i := 0; i < 3; i++ {
        a[row][colum - i] = "X"
      }
    } else {
      //fmt.Println("..left.right..")
      for i := 0; i < 3; i++ {
        a[row][colum + i] = "X"
      }
    }
  }


    return a
}

//checks if attack is in range of the board
func ValidAttack(c []string) bool {
  d, err := strconv.Atoi(c[1]);
  if err != nil {}

  switch c[0] {
  case "A", "B", "C", "D", "E":
    if d >= 1 && d <= 5 {
      return true
    }
  }
  return false
}

//Player makes attack and a new gameboard is returned with the attack on it
func Playerattack(gameboard [6][6]string, hidden [6][6]string) [6][6]string {
  a := gameboard
  var b string
  var c []string

  fmt.Println("Enter a space to attack. Range A1 - E5")
  //Checks if valid spot on board
  //If valid continues, Else prompts for new attack
  for  {
    n, err := fmt.Scanf("%s\n", &b)
      if err != nil {
        fmt.Println(n, err)
    }
    c = strings.Split(strings.ToUpper(b), "")

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
  }

  switch CheckIfMatch(c, hidden) {
  case true:
    a[row][colum] = "X"
  case false:
    a[row][colum] = "*"
  }

  return a
}


//Calling clears the terminal screen
// Windows cl := exec.Command("cmd", "/c", "cls")
// unix base cl := exec.Command("clear")
func clearscreen()  {
  cl := exec.Command("cmd", "/c", "cls")
  cl.Stdout = os.Stdout
  cl.Run()
}

func main() {
  fmt.Printf("Welcome to the game!\n")
  board := Newboard()
  //key := Newboard()
  key := Placeships()
  //testing Placeships()
  /*
  Drawboard(key)
  for i := 0; i < 30; i++ {
    key = Placeships()
    //Drawboard(key)
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
