# battleship
very bulky and ugly.
Many lines left in that are used for testing and troubleshooting most are commented out.

Current version only plays on a 5x5 board with one ship 3 spaces long.


You will need to edit clearscreen() to work with your OS
  line 245
  Windows cl := exec.Command("cmd", "/c", "cls")
  unix base cl := exec.Command("clear")
