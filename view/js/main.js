// scroll to the bottom of the buffer when the page has finished loading
function scrollToBottom(id){
    var element = document.getElementById(id);
    element.scrollTop = element.scrollHeight;
}

document.addEventListener('DOMContentLoaded', function() {
    scrollToBottom("buffer");
});

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
