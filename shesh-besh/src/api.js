import axios from "axios"

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