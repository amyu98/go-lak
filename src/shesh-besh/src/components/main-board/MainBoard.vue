<template>
  <header>
    <h1>Shesh Besh</h1>
  </header>

  <button @click="rollDice">Roll The Dice</button>
  <DiceContainer :dices="dices" />

  <div class="board">
    <div
      class="row"
      v-for="(row, index) in rows"
      v-bind:class="rowClass(index)"
    >
      <div class="section" v-for="section in row">
        <div class="cell" v-for="cell in section">
          <BoardCell :cell="cell" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import BoardCell from "../board-cell/BoardCell";
import DiceContainer from "../dice-container/DiceContainer";
export default {
  name: "MainBoard",
  components: { BoardCell, DiceContainer },
  props: {
    rawBoard: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      rows: [],
      dices: [],
    };
  },
  mounted() {
    this.rows = this.rawBoardToRows(this.rawBoard);
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
    rollDice() {
      this.dices = [];
      for (let i = 0; i < 2; i++) {
        this.dices.push(Math.floor(Math.random() * 6) + 1);
      }
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
