//// functions
function scrollToBottom(id){
    var element = document.getElementById(id);
    element.scrollTop = element.scrollHeight;
}

function sendMessage(message) {
    const form = document.querySelector("#input");
    const data = Object.fromEntries(new FormData(form).entries());
    console.log(data);
    return(data);
}

function newMessage(user, message, action) {
    log = document.getElementById("log");
    line = document.createElement("li");
    userName = document.createElement("a");
    userNameSpan = document.createElement("span");
    userMessageSpan = document.createElement("span");
    userText = document.createTextNode(user);
    userMessage = document.createTextNode(message);
    log.appendChild(line);
    line.appendChild(userNameSpan).classList.add("name", action)
    userNameSpan.appendChild(userName);
    userName.appendChild(userText);
    line.appendChild(userMessageSpan).classList.add("message");
    userMessageSpan.appendChild(userMessage);
    scrollToBottom("buffer");
}


//// exec
document.addEventListener('DOMContentLoaded', function() {
    scrollToBottom("buffer");
});

input = document.getElementById("input");
input.addEventListener('submit', e => {
    sendMessage();
    e.preventDefault();
});
