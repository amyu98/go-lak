<template>
  <div class="jail-container">
    <div class="jail" @click="jailClicked('black')">
      Black Jail: {{ blackJail }}
    </div>
    <div class="jail" @click="jailClicked('white')">
      White Jail: {{ whiteJail }}
    </div>
  </div>
</template>

<script>
import { useBoardState } from "../../store/state";
export default {
  name: "JailContainer",
  data() {
    return {
      state: useBoardState(),
    };
  },
  mounted() {
    this.state = useBoardState();
    this.state.$subscribe(() => {
      this.state = useBoardState();
    });
  },
  computed: {
    blackJail() {
      return this.state.blackJail;
    },
    whiteJail() {
      return this.state.whiteJail;
    },
  },
  methods: {
    jailClicked(color) {
      this.$emit("jailClicked", color);
    },
  },
};
</script>

<style>
.jail-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 100px;
}
.jail {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 10px;
  border: 1px solid black;
  width: 100px;
  height: 30px;
  padding: 3px;
}
</style>
