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
    event.preventDefault();
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
            item.forEach(subItem => {
                let td = document.createElement("td");
                td.textContent = subItem;
                tr.appendChild(td);
            })
            let td = document.createElement("td");
            let btn = document.createElement("input");
            btn.type = "button"
            btn.value = "Удалить"
            btn.onclick = () => console.log(item)
            td.appendChild(btn);
            tr.appendChild(td);
            market.appendChild(tr);
        })
    })
}