document.getElementById("button_my_items").addEventListener("click", goMyItems)

function goMyItems(event){
    event.preventDefault()
    document.location.replace("./account")
}