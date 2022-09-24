<template>
  <div class="game-dice">
    {{ dice }}
  </div>
</template>

<script>
export default {
  name: "GameDice",
  props: {
    dice: {
      type: Number,
      required: true,
    },
  },
};
</script>

<style>
.game-dice {
  display: inline-block;
  width: 30px;
  height: 30px;
  border: 1px solid black;
  text-align: center;
  vertical-align: middle;
  line-height: 30px;
  margin: 2px;
}
</style>
