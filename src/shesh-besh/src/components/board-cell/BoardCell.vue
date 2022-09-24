<template>
  <div class="cell">
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
  props: ["cell"],
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
}
.pieces-list {
    display: flex;
    flex-direction: column;
    justify-content: center;
}
</style>
