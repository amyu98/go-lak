import axios from "axios"
import { useBoardState } from "./store/state"

export function initHttp() {
    axios.defaults.baseURL = "http://localhost:8080/api/v1";
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
    const data = await axios.get(
        `${gameslug}/possible_moves?selectedCell=` + selectedCell.Index
    );
    return data.data;
};

export const movePiece = async (selectedCell, targetCell) => {
    const data = await axios.post(
        `${gameslug}/move_piece?selectedCell=${selectedCell.Index}&targetCell=${targetCell.Index}`,
    );
    return data.data;
};