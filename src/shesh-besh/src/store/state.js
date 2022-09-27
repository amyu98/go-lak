import { defineStore } from 'pinia'

export const useBoardState = defineStore('boardState', {
  state: () => {
    return { board: null, turn: "black", dice: null, selectedCell: null, moveState: "selecting" }
  },
  actions: {
    setBoard(board) {
      this.board = board
    },
    setTurn(turn) {
      this.turn = turn
    },
    setDice(dice) {
      this.dice = dice
    },
    setSelectedCell(selectedCell) {
      this.selectedCell = selectedCell
    },
    setMoveState(moveState) {
      this.moveState = moveState
    },
  },
})