<template>
  <div class="cell" @click="emitClick" :class="{optional: optinalMove}">
    <div v-for="bp in cell.BlackPieces" class="pieces-list">
      <BoardPiece pieceColor="black" />
    </div>

    <div v-for="wp in cell.WhitePieces" class="pieces-list">
      <BoardPiece pieceColor="white" />
    </div>
  </div>
</template>

<script>
import BoardPiece from "../board-pieces/BoardPiece";
export default {
  name: "BoardCell",
  components: { BoardPiece },
  props: ["cell", "optinalMove"],
  methods: {
    emitClick() {
      this.$emit("cellClick", this.cell);
    },
  },
};
</script>

<style>
.cell {
    width: 30px;
    display: flex;
    flex-direction: column;
    background: beige;
    height: 100%;
    border: 1px solid #dbdbdb;
    flex-wrap: wrap;
    overflow: hidden;
}
.pieces-list {
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.optional{
    box-sizing: border-box;
    border: 2px solid brown;
}
</style>
