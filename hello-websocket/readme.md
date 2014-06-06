# Hello World WebSocket

Server that emit 'hello' message as soon as a client connects.

* [Run](#run)

## Run

    go run server.go
    open localhost:4000
    enter in the browser's console:
    var sock = new WebSocket("ws://localhost:4000/ws"); sock.onmessage = function(m) { console.log("received:", m.data) }
    => received: hello

