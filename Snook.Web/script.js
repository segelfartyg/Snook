var button = document.getElementById("sendMsgBtn")

var messageDiv = document.getElementById("messageArea")





let socket = new WebSocket("ws://localhost:3000/ws")

socket.addEventListener("message", (event) => {

    console.log(event)

    messageDiv.innerHTML += `<p>${event.data}</p>`
})

socket.onmessage = (event) => {
    console.log(event)
}



button.addEventListener("click", (event) => {

    var messageInput = document.getElementById("inputMessage")

  
        socket.send(messageInput.value)


})