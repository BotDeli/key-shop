let action = document.getElementById("window");
let prefix = document.getElementById("prefix");
let close = document.getElementById("close");
let submit = document.getElementById("window_authorization");
let login = document.getElementById("login_inp");
let password = document.getElementById("password_inp");
let output = document.getElementById("output");
let button_submit = document.getElementById("submit");


close.addEventListener("click", closeWindow);
submit.addEventListener("submit", sendRequest);
login.addEventListener("input", updateStatus)
password.addEventListener("input", updateStatus)
document.getElementById("login").addEventListener("click", openLoginWindow);
document.getElementById("signUp").addEventListener("click", openSignUpWindow);

action.hidden = true;
let openedLogin = false;
let openedSignUp = false;
let status = false;

const address = "http://localhost:8080";

function openLoginWindow() {
    if (openedSignUp){
        closeWindow();
    }
    if (!openedLogin){
        action.hidden = false;
        prefix.innerText = "Login";
        openedLogin = true;
        close.classList.add("login");
        close.classList.remove("sign");
    }
}

function openSignUpWindow() {
    if (openedLogin){
        closeWindow();
    }
    if (!openedSignUp){
        action.hidden = false;
        prefix.innerText = "SignUp";
        openedSignUp = true;
        close.classList.remove("login");
        close.classList.add("sign");
    }
}

function closeWindow() {
    action.hidden = true;
    openedLogin = false;
    openedSignUp = false;
    status = false;
    login.value = "";
    password.value = "";
    output.innerText = "";
    wait();
}

function sendRequest(event){
    event.preventDefault();
    if (openedLogin || openedSignUp){
        if (status){
            if (openedLogin){
                sendFetchRequest(address+"/login");
            } else if (openedSignUp){
                sendFetchRequest(address+"/registration");
            }
        } else {
            output.innerText = "Minimum length login and password is 6 chars and started char is letter";
        }
    }
}

function sendFetchRequest(url){
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Sender": "application/json",
            "Sender-fields": "login, password",
            "Accept": "application/json",
            "Accept-fields": "ERR"
        },
        body: JSON.stringify({"login": login.value, "password": password.value})
    }).then(response => response.json()).then(data => {
        if (data.ERR !== "<nil>"){
            output.innerText = data.ERR;
        } else {
            document.location.replace("./authorized")
        }
    }).catch(error => {
        output.innerText = "Ошибка..." + error;
    });
}

function updateStatus(event){
    event.preventDefault();
    status = login.value.length >= 6 && password.value.length >= 6 &&
        isLetter(login.value[0]) && isLetter(password.value[0])
    if (status){
        ready();
    } else {
        wait();
    }
}

function wait(){
    button_submit.style.color = "black";
    button_submit.style.background = "#eff2f7";
}

function ready(){
    button_submit.style.color = "#eff2f7";
    button_submit.style.background = "#57d257";
}

function isLetter(char){
    return /[a-zA-Z]/.test(char);
}