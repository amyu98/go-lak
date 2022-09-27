import axios from "axios"
import { useBoardState } from "./store/state"

export function initHttp() {
    axios.defaults.baseURL = "http://localhost:8080";
    axios.defaults.headers.common["Accept"] = "application/json";
    axios.defaults.headers.common["Content-Type"] = "application/json";
    axios.defaults.transformResponse = [
        function (data) {
            return data ? JSON.parse(data) : data;
        },
    ];
}

export const getNewGame = async () => {
    const data = await axios.get(
        "/new_game"
    );
    return data.data;
};

export const getPossibleMoves = async (selectedCell) => {
    const state = useBoardState();
    const data = await axios.post(
        "/possible_moves?selectedCell=" + selectedCell.Index,
        {
            "CurrentPlayer": state.turn,
            "DiceRoll": state.dice,
            "BoardState": state.board
        }
    );
    return data.data;
};