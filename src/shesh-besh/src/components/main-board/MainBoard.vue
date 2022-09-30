<template>
  <header>
    <!-- <h1>Shesh Besh</h1> -->
  </header>

  <DiceContainer :dice="this.dice" />

  <div class="board">
    <div
      class="row"
      v-for="(row, index) in rows"
      v-bind:class="rowClass(index)"
    >
      <div class="section" v-for="section in row">
        <div class="cell" v-for="cell in section">
          <BoardCell
            :cell="cell"
            @cellClick="cellClick"
            :optinalMove="optionalMoves.includes(cell.Index)"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BoardCell from "../board-cell/BoardCell";
import DiceContainer from "../dice-container/DiceContainer";
import { getPossibleMoves, movePiece } from "../../api";
import { useBoardState } from "../../store/state";

export default {
  name: "MainBoard",
  components: { BoardCell, DiceContainer },
  props: {},
  data() {
    return {
      rows: [],
      dice: [],
      optionalMoves: [],
    };
  },
  mounted() {
    this.state = useBoardState();
    this.rows = this.rawBoardToRows(this.state.board);
    this.state.$subscribe(() => {
      this.rows = this.rawBoardToRows(this.state.board);
      this.optionalMoves = this.state.optionalMoves;
      this.dice = this.state.dice;
    });
    this.state.rollDice();
  },
  computed: {},
  methods: {
    rowClass(index) {
      if (index == 0) {
        return "top-row";
      }
      return "bottom-row";
    },
    rawBoardToRows(rawBoard) {
      const rows = [];
      for (let i = 0; i < 2; i++) {
        const row = [];
        for (let j = 0; j < 2; j++) {
          const section = [];
          for (let k = 0; k < 6; k++) {
            const cell = rawBoard[i * 12 + j * 6 + k];
            section.push(cell);
          }
          row.push(section);
        }
        rows.push(row);
      }
      return rows;
    },
    async cellClick(cell) {
      switch (this.state.moveState) {
        case "selecting":
          this.state.setSelectedCell(cell);
          this.state.indicatePossibleMoves(cell);
          break;
        case "moving":
          await this.movePiece(cell);
          break;
        default:
          break;
      }
    },
    async movePiece(cell) {
      this.state.movePiece(cell)
     },
  },
};
</script>

<style>
.board {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 45vh;
}

.section {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  border: 3px solid black;
  margin: 1px;
}

.row {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 45%;
  margin: 3px;
}

.top-row {
  flex-direction: row;
}

.bottom-row {
  flex-direction: row-reverse;
}

.bottom-row .section {
  flex-direction: row-reverse;
}

.bottom-row .cell {
  justify-content: end;
}
</style>
