document.getElementById("button_add_new_item").addEventListener("click", openCreatorItem);
document.getElementById("close").addEventListener("click", closeCreatorItem)
document.getElementById("new_item").addEventListener("submit", sendRequestAddItem)

newItem = document.getElementById("new_item");
newItem.hidden = true

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
            fetch("/add_item", {
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
            }).then(response => {
                if (response.status === 400){
                    alert("Вы превысили лимит добавления предметов, лимит 10!")
                } else {
                    document.location.reload();
                }
            }).catch(error => console.log(error));
        } else {
            alert("Цена и количество должны быть больше 0")
        }
    } else {
        alert("Заполните все поля!")
    }
}

function allDontEmpty(name, description, count, cost) {
    return name !== "" &&
        description !== "" &&
        count !== "" &&
        cost !== ""
}