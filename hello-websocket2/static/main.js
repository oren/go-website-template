var sock = new WebSocket("ws://localhost:4000/ws");

sock.onmessage = function(m) {
    console.log("received:", m.data)
}
