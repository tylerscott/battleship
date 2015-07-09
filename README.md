# battleship
very bulky and ugly.
Many lines left in that are used for testing and troubleshooting most are commented out.

Current version only plays on a 10x10 board with 3 ship 3 spaces long.
number of ships can be set at line 243


You will need to edit clearscreen() to work with your OS
line 341
Windows cl := exec.Command("cmd", "/c", "cls")
unix base cl := exec.Command("clear")
