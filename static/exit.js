document.getElementById("exit").addEventListener("click", sendExitRequest);
function sendExitRequest() {
    fetch("/exit", {
        method: "POST",
    }).then(_ => {
        document.location.reload();
    }).catch(err => console.log(err));
}