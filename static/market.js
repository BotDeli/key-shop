lastPage = document.getElementById("last_page");
nextPage = document.getElementById("next_page");
lastPage.addEventListener("click", lastPage);
nextPage.addEventListener("click", nextPage);

function lastPage(event){
    event.preventDefault();
    page--;
}

function nextPage(event){
    event.preventDefault();
    page++;
}

let page = 1;
let countPages = 0;
getCountPages()
function getCountPages() {
    fetch(address+"/count_pages", {
        method: "GET",
        headers: {
            "Accept": "application/json",
        }
    }).then(response => response.json()).
    then(data => {
        countPages = data.pages;
    }).catch(err => console.log(err));
}

document.addEventListener("DOMContentLoaded", loadPage);

function loadPage(event){
    event.preventDefault();
    hideSelectorsPage();
    sendRequestGetItems();
}



function hideSelectorsPage(){
    lastPage.hidden = true;
    nextPage.hidden = true;
}

market = document.getElementById("market");
startListHtml = market.innerHTML;

function sendRequestGetItems(){
    fetch(address+"/items", {
        method: "GET",
        headers: {
            "Accept": "application/json"
        }
    }).then(response => response.json()).
    then(data => {
        const listItems = data.items;
        listItems.forEach(item => {
            let mainLi = document.createElement("li");
            mainLi.className = "sell-item"
            let ul = document.createElement("ul");
            ul.className = "sell-item-description"
            item.forEach(subItem => {
                let li = document.createElement("li");
                li.textContent = subItem;
                ul.appendChild(li);
            })
            mainLi.appendChild(ul)
            market.appendChild(mainLi);
        })
        if (listItems.length >= 20){
            showSelectorsPage();
        }
    })
}

function showSelectorsPage(){
    if (page > 1){
        lastPage.hidden = false;
    }
    if (page < countPages){
        nextPage.hidden = false;
    }
}

document.getElementById("button_update_list").addEventListener("click", updateList)

function updateList(event){
    event.preventDefault();
    market.innerHTML = startListHtml;
    sendRequestGetItems();
}