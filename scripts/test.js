import http from "k6/http";

export let options = {
    vus: 3,
    duration: "5s",
};

export default function () {
    // http.get("http://test.k6.io");
    http.get("http://host.docker.internal:8000");
}
