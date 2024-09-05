/* Create the game board struct */

package board

type Board struct {
	board [9][9]int
	hints [9][9][9]bool
	playerBoard [9][9]int
	mistakes int
}
