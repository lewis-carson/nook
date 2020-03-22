// scroll to the bottom of the buffer
function scrollToBottom(id){
   var element = document.getElementById(id);
   element.scrollTop = element.scrollHeight;
}

document.addEventListener('DOMContentLoaded', function() {
    scrollToBottom("buffer");
});


// add a message to the buffer
// function addMessage() {
//     log = document.getElementById("log");
//     line = document.createElement("li");
//     userName = document.createElement("a");
//     userNameSpan = document.createElement("span");
//     userMessageSpan = document.createElement("span");

//     userText = document.createTextNode("eti");
//     userMessage = document.createTextNode("heyyyyyyy");

//     log.appendChild(line);
//     line.appendChild(userNameSpan).classList.add("name", "msg")

//     userNameSpan.appendChild(userName);
//     userName.appendChild(userText);

//     line.appendChild(userMessageSpan).classList.add("message");
//     userMessageSpan.appendChild(userMessage);

//     scrollToBottom("buffer");
// }

