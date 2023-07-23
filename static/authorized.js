document.getElementById("exit").addEventListener("click", sendExitRequest);
document.getElementById("button_add_new_item").addEventListener("click", openCreatorItem);
document.getElementById("close").addEventListener("click", closeCreatorItem)
document.getElementById("new_item").addEventListener("submit", sendRequestAddItem)

newItem = document.getElementById("new_item");
newItem.hidden = true

const address = "http://localhost:8080"

function sendExitRequest() {
    fetch("/exit", {
        method: "POST",
    }).then(_ => {
        document.location.reload();
    }).catch(err => console.log(err));
}

function openCreatorItem(event) {
    event.preventDefault();
    newItem.hidden = false;
}

function closeCreatorItem(event) {
    event.preventDefault();
    newItem.hidden = true;
}

const nameInput = document.getElementById("name")
const descriptionInput = document.getElementById("description")
const countInput = document.getElementById("count")
const costInput = document.getElementById("cost")

function sendRequestAddItem(event) {
    event.preventDefault();

    let name = nameInput.value;
    let description = descriptionInput.value;
    let count = countInput.value;
    let cost = costInput.value;

    if (allDontEmpty(name, description, count, cost)) {
        if (cost > 0 && count > 0) {
            fetch(address + "/add_item", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Sender": "application/json",
                    "Sender-fields": "name, description, count, cost"
                },
                body: JSON.stringify({
                    "name": name,
                    "description": description,
                    "count": count,
                    "cost": cost
                })
            }).then(_ => {
                newItem.hidden = true;
                updateList()
            }).catch(error => console.log(error));
        } else {
            alert("Цена и количество должны быть больше 0")
        }
    } else {
        alert("Заполните все поля!")
    }
}
//
function allDontEmpty(name, description, count, cost) {
    return name !== "" &&
        description !== "" &&
        count !== "" &&
        cost !== ""
}