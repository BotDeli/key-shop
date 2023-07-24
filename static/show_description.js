let message_window = document.getElementById("window_message");
message_window.addEventListener("click", closeMessage);
message_window.hidden = true;

function closeMessage(event) {
    event.preventDefault();
    message_window.hidden = true;
}