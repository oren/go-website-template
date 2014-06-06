document.addEventListener('DOMContentLoaded', ready);

function ready() {
  var sock = new WebSocket("ws://localhost:4000/ws");

  sock.onmessage = function(m) {
      console.log("received:", m.data);
  };

  $btn = document.querySelector('#form');
  $msg = document.querySelector('#msg');

  $btn.addEventListener('submit', send);

  function send(e) {
    sock.send(msg.value);
    msg.value = "";
    e.preventDefault();
  }
}
