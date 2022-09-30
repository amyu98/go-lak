import { useBoardState } from "../store/state";
import { getGame, createNewGame, selectCell, movePiece } from "../api";

export const gameMixin = {
  data() {
    return {
      state: null,
    };
  },
  created: function () {},
  methods: {
    async startGame() {
      this.state = useBoardState();
      const path = window.location.href;

      if (path.includes("slug=")) {
        const slug = path.split("slug=")[1];
        this.state.setGameSlug(slug);
        await this.updateBoardState();
      } else {
        await this.createNewGame();
      }
    },
    async updateBoardState() {
      const gameState = await getGame();
      this.state = useBoardState();
      this.state.setBoardState(gameState);
    },
    async createNewGame() {
      this.state = useBoardState();
      const s = await createNewGame();
      window.history.pushState({}, "", `?slug=${s.Slug}`);
      this.state.setBoardState(s);
    },
    async cellClicked(cell) {
      this.state = useBoardState();
      switch (this.state.actionState()) {
        case "selecting":
          await this.selectCell(cell);
          break;
        case "moving":
          await this.movePiece(cell);
          break;
      }
    },
    async jailClicked(color) {
      this.state = useBoardState();
      await this.selectCell({ Index: color == "white" ? 24 : -1 });
    },
    async selectCell(cell) {
      this.state = useBoardState();
      const s = await selectCell(cell);
      this.state.setBoardState(s);
    },
    async movePiece(cell) {
      this.state = useBoardState();
      const s = await movePiece(cell);
      this.state.setBoardState(s);
    }

  },
};
