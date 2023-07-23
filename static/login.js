setLogin = document.getElementById("set-login")

sendRequestGetLogin()

function sendRequestGetLogin(){
    fetch("/get_login").
    then(response => response.json()).
    then(data => {
        setLogin.innerText = data.login
    }).catch(err => {
        console.log(err)
        setLogin.innerText = "My account"
    })
}