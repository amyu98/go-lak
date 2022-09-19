<template>
  <header>
    <h1>Shesh Besh</h1>
  </header>

  Board:
  <div class="board">
    <div class="row" v-for="row in rows">
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
export default {
  name: "MainBoard",
  components: { BoardCell },
  props: {
    rawBoard: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      rows: [],
    };
  },
  mounted() {
    this.rows = this.rawBoardToRows(this.rawBoard);
  },
  computed: {},
  methods: {
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
  },
};
</script>

<style>
.board {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 80vh;
}

.section {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.row {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  height: 45%;
}
</style>
