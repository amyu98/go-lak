import { defineStore } from "pinia";
import { getPossibleMoves, movePiece } from "../api"

export const useBoardState = defineStore("boardState", {
  state: () => {
    return {
      board: null,
      turn: "black",
      dice: null,
      movesActed: [],
      selectedCell: null,
      moveState: "selecting",
      selectedCell: null,
      optionalMoves: [],
    };
  },
  actions: {
    setBoard(board) {
      this.board = board;
    },
    setTurn(turn) {
      this.turn = turn;
    },
    setMoveState(moveState) {
      this.moveState = moveState;
    },
    setSelectedCell(selectedCell) {
      this.selectedCell = selectedCell;
    },
    setOptionalMoves(optionalMoves) {
      this.optionalMoves = optionalMoves;
    },
    updateBoard(boardState) {
      this.board = boardState.BoardState;
      this.turn = boardState.CurrentPlayer;
      this.dice = boardState.DiceRoll;
    },
    userMovedPiece(boardState, move) {
      // this.board = boardState.BoardState;
      // this.turn = boardState.CurrentPlayer;
      // this.dice = boardState.DiceRoll;
      // const movesToMake = this.dice[0] === this.dice[1] ? 4 : 2;
      // this.movesActed.push(move);
      // if (this.movesActed.size === movesToMake) {
      //   this.movesActed = []
      //   this.moveState = "selecting";
      //   this.turn = this.turn === "black" ? "white" : "black";
      // }
      // this.selectedCell = null;
    },
    rollDice() {
      // let ds = [];
      // for (let i = 0; i < 2; i++) {
      //   ds.push(Math.floor(Math.random() * 6) + 1);
      // }
      // this.setDice(ds);
    },
    setDice(dice) {
      this.dice = dice;
    },
    async indicatePossibleMoves(cell) {
      // if (this.cellCannotBeClicked(cell)) {
      //   return;
      // }
      // const possibleMoves = await getPossibleMoves(cell);
      // this.setOptionalMoves(possibleMoves);
      // this.setMoveState("moving");
    },
    cellCannotBeClicked(cell) {
      // switch (this.moveState) {
      //   case "selecting":
      //     return (
      //       (cell.BlackPieces == 0 && this.turn == "black") ||
      //       (cell.WhitePieces == 0 && this.turn == "white")
      //     );
      //   case "moving":
      //     return !this.optionalMoves.includes(cell.Index);
      //   default:
      //     break;
      // }
    },
    async movePiece(cell){
      if (this.cellCannotBeClicked(cell)) {
        return;
      }
      const res = await movePiece(this.selectedCell, cell);
      // const move = Math.abs(this.selectedCell.Index - cell.Index)
      // this.userMovedPiece(res, move);
    }
  },
});
