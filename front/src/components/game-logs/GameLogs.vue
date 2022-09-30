<template>
  <div class="game-logs">
    <div class="log" v-for="log in logs" :key="log">
      {{ log.Tick }}
      {{ log.Msg }}
    </div>
  </div>
</template>

<script>
import { useBoardState } from "../../store/state";
export default {
  name: "GameLogs",
  data() {
    return {
      logs: [],
      state: null,
    };
  },
  mounted() {
    this.state = useBoardState();
    this.logs = this.state.logs;
    this.state.$subscribe(() => {
      this.logs = this.state.logs;
    });
  },
};
</script>

<style>
.game-logs {
    display: flex;
    flex-direction: column;
    margin: 10px;
    width: 300px;
}
.log {
  display: flex;
  margin: 2px;
  font-size: 14px;
}
</style>
