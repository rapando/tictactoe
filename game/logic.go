package game

import "fmt"

type gameObject int

const (
	blank gameObject = iota
	cross
	circle
)

type Game struct {
	GameID         int64
	GameSpace      [3][3]gameObject
	PlayerXTurn    bool
	EmptyPositions []string
	GameOver       bool
	PlayerXWinner  bool
	PlayerOWinner  bool
}

func (g *Game) NewGame() {
	g.PlayerXTurn = true
}

func (g *Game) PrintGame() {
	var pieces = []rune{'_', 'X', 'O'}
	fmt.Println("---------------")
	for _, x := range g.GameSpace {
		for _, c := range x {
			fmt.Printf("%-4c", pieces[c])
		}
		fmt.Println()
	}
	fmt.Println("---------------")
}

func (g *Game) BoardIsFull() bool {
	var blanks = 0
	var emptyPositions []string
	for i, x := range g.GameSpace {
		for j, c := range x {
			if c == blank {
				blanks++
				emptyPositions = append(emptyPositions, fmt.Sprintf("%d,%d", i, j))
			}
		}
	}
	g.EmptyPositions = emptyPositions
	g.GameOver = blanks == 0
	return blanks == 0
}

func (g *Game) CheckWinner() {
	var winningCombinations = [][]string{
		{"00", "01", "02"},
		{"00", "11", "22"},
		{"00", "10", "20"},
		{"01", "11", "21"},
		{"02", "11", "20"},
		{"02", "12", "22"},
		{"10", "11", "22"},
		{"20", "21", "22"},
	}

	// first collect the positions of x and y
	var crosses, circles []string
	for i, x := range g.GameSpace {
		for j, c := range x {
			if c == circle {
				circles = append(circles, fmt.Sprintf("%d%d", i, j))
			} else if c == cross {
				crosses = append(crosses, fmt.Sprintf("%d%d", i, j))
			}
		}
	}
	if xWIn := checkForWin(crosses, winningCombinations); xWIn {
		g.PlayerXWinner = true
		g.GameOver = true
	}
	if cWin := checkForWin(circles, winningCombinations); cWin {
		g.PlayerOWinner = true
		g.GameOver = true
	}
}

func (g *Game) Turn(i, j int) (err error) {
	g.CheckWinner()
	if full := g.BoardIsFull(); full || g.GameOver {
		return fmt.Errorf("game over")
	}

	var playPiece = circle
	if g.PlayerXTurn {
		playPiece = cross
	}

	// logic here
	if !g.isPositionEmpty(i, j) {
		return fmt.Errorf("position is not empty")
	}
	g.GameSpace[i][j] = playPiece

	g.PlayerXTurn = !g.PlayerXTurn
	g.CheckWinner()
	return nil
}

func (g *Game) isPositionEmpty(i, j int) bool {
	coordinates := fmt.Sprintf("%d,%d", i, j)
	for _, pos := range g.EmptyPositions {
		if pos == coordinates {
			return true
		}
	}
	return false
}

func checkForWin(playerPositions []string, winningCombinations [][]string) bool {
	for _, w := range winningCombinations {
		counter := 0
		// count how many items in  positions appear in [][]string
		for _, p := range playerPositions {
			if existsInArray(p, w) {
				counter++
				if counter == 3 {
					return true
				}
			}
		}
	}
	return false
}

func existsInArray(element string, arr []string) bool {
	exists := false
	for _, a := range arr {
		if element == a {
			exists = true
			break
		}
	}
	return exists
}
