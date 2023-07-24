document.getElementById("return_to_market").addEventListener("click", goMarket)

function goMarket(event) {
    event.preventDefault();
    document.location.replace("./authorized")
}

let market = document.getElementById("market");
let startListHtml = market.innerHTML;

let upd_btn = document.getElementById("button_update_list")
upd_btn.addEventListener("click", updateList)

function updateList(event){
    if (event !== null){
        event.preventDefault();
    }
    market.innerHTML = startListHtml;
    sendRequestGetMyItems();
}

document.addEventListener("DOMContentLoaded", loadPage);

function loadPage(event){
    event.preventDefault();
    sendRequestGetMyItems();
}

var listItems;
function sendRequestGetMyItems(){
    fetch("/my_items", {}).
    then(response => response.json()).
    then(data => {
        listItems = data.items;
        listItems.forEach(item => {
            let tr = document.createElement("tr");
            tr.className = "sell-item sell-item-description"

            for(let i = 1; i < item.length; i++){
                let td = document.createElement("td");
                td.textContent = item[i];
                tr.appendChild(td);
            }

            let td = document.createElement("td");
            let btn = document.createElement("input");
            btn.type = "button"
            btn.value = "Удалить"
            btn.onclick = () => {
                fetch("/delete_item", {
                    method: "DELETE",
                    headers: {
                        "Sender": "application/json",
                    },
                    body: JSON.stringify({
                        "name": item[1],
                        "description": item[0],
                        "count": item[2],
                        "cost": item[3],
                    })
                }).then(() => {
                    message_window.hidden = true;
                    updateList(null);
                }).catch(error => console.log(error))
            }
            td.appendChild(btn);
            tr.appendChild(td);

            tr.style = "cursor: pointer;";
            tr.onclick = () => {
                document.getElementById("message").innerText = item[0];
                document.getElementById("window_message").hidden = false;
            }

            market.appendChild(tr);
        })
    })
}