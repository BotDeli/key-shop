let message = document.getElementById("window_message");
message.addEventListener("click", closeMessage);
message.hidden = true;

function closeMessage(event) {
    event.preventDefault();
    message.hidden = true;
}