var button = document.getElementById("sendMsgBtn");
var messageDiv = document.getElementById("messageArea");

let socket = new WebSocket("ws://localhost:80/ws");

socket.addEventListener("message", (event) => {
  console.log(event);

  messageDiv.innerHTML += `<p>${event.data}</p>`;
});

socket.onmessage = (event) => {
  console.log(event);
};

document.addEventListener("keypress", (event) => {
  if (event.key == "Enter") {
    var messageInput = document.getElementById("inputMessage");
    socket.send(messageInput.value);
    messageInput.value = "";
  }
});

button.addEventListener("click", (event) => {
  var messageInput = document.getElementById("inputMessage");

  socket.send(messageInput.value);
  messageInput.value = "";
});
