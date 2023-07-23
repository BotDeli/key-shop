let page = 1;

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
    fetch("/items", {
        method: "POST",
        headers: {
            "Accept": "application/json"
        },
        body: JSON.stringify({"page": page})
    }).then(response => response.json()).
    then(data => {
        const listItems = data.items;
        if (listItems != null){
            listItems.forEach(item => {
                let tr = document.createElement("tr");
                tr.className = "sell-item sell-item-description"
                item.forEach(subItem => {
                    let td = document.createElement("td");
                    td.textContent = subItem;
                    tr.appendChild(td);
                })
                market.appendChild(tr);
            })
        }
    })
    showSelectorsPage();
}

function showSelectorsPage(){
    fetch("/count_pages", {
        method: "GET",
        headers: {
            "Accept": "application/json",
        }
    }).then(response => response.json()).
    then(data => {
        const countPages = data.pages;
        lastPage.hidden = page <= 1;
        nextPage.hidden = page >= countPages;
    }).catch(err => console.log(err));
}

function updateList(event){
    event.preventDefault();
    market.innerHTML = startListHtml;
    sendRequestGetItems();
}
lastPage = document.getElementById("last_page");
nextPage = document.getElementById("next_page");
lastPage.addEventListener("click", setLastPage);
nextPage.addEventListener("click", setNextPage);

let upd_btn = document.getElementById("button_update_list")
upd_btn.addEventListener("click", updateList)

function setLastPage(event){
    event.preventDefault();
    page--;
    upd_btn.click();
}

function setNextPage(event){
    event.preventDefault();
    page++;
    upd_btn.click();
}