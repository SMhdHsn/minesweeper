package game

import (
	"strconv"
	"strings"

	"github.com/smhdhsn/minesweeper/content"
	"github.com/smhdhsn/minesweeper/interaction"
)

type Board [][]cell

// countAreaBombs, counts bombs in neighbourhood of a given cell.
func (b *Board) countAreaBombs(c *cell) (areaBombs int) {
	for rowOff := c.row - 1; rowOff <= c.row+1; rowOff++ {
		for colOff := c.col - 1; colOff <= c.col+1; colOff++ {
			if rowOff >= 0 && rowOff < len(*b) && colOff >= 0 && colOff < len((*b)[0]) {
				if (*b)[rowOff][colOff].isBomb {
					areaBombs++
				}
			}
		}
	}

	(*b)[c.row][c.col].areaBombs = areaBombs

	return
}

// PlantBomb, plants given amount of bombs in the board.
func (b *Board) PlantBomb(bombCount int) {
	for {
		row := generateRandom(0, len(*b)-1)
		col := generateRandom(1, len((*b)[0])-1)

		if !(*b)[row][col].isBomb && !(*b)[row][col].isRevealed {
			(*b)[row][col].isBomb = true

			bombCount--
		}

		if bombCount == 0 {
			break
		}
	}
}

// Draw, draws the grid for the board.
func (b *Board) Draw() {
	for _, row := range *b {
		str := strings.Repeat("█", len(row)*3) + "█"
		interaction.Println(str, content.White)

		for _, cell := range row {
			if cell.isRevealed {
				var color string
				var areaBombs string

				switch cell.areaBombs {
				case 9, 8, 7:
					color = content.Red
				case 6, 5, 4:
					color = content.Yellow
				case 3, 2, 1, 0:
					color = content.White
				}

				if cell.areaBombs == 0 {
					areaBombs = " "
				} else {
					areaBombs = strconv.FormatInt(int64(cell.areaBombs), 10)
				}

				interaction.Print("█ ", content.White)
				interaction.Print(areaBombs, color)
			} else {
				interaction.Print("█▓▓", content.White)
			}
		}

		interaction.NewLine(content.White + "█")
	}
	interaction.NewLine(content.White + strings.Repeat("█", len((*b)[0])*3) + "█")
}

// revealeArea, reveals the cells around a given cell with '0' areaBombs.
func (b *Board) revealeArea(c *cell) {
	for rowOff := c.row - 1; rowOff <= c.row+1; rowOff++ {
		for colOff := c.col - 1; colOff <= c.col+1; colOff++ {
			if rowOff >= 0 && rowOff < len(*b) && colOff >= 0 && colOff < len((*b)[0]) {
				(*b)[rowOff][colOff].Reveale()
			}
		}
	}
}

// Over, reveals every cell of the board.
// Must only be executed if the game is over.
func (b *Board) Over() {
	for rowIndex, row := range *b {
		for colIndex := range row {
			(*b)[rowIndex][colIndex].Reveale()
		}
	}
}
