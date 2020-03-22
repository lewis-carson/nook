//// functions
function scrollToBottom(id){
    var element = document.getElementById(id);
    element.scrollTop = element.scrollHeight;
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

function currentChannel(server, channel) {
    chans = document.getElementById("chans");
    column = document.createElement("ul");
    line = document.createElement("li");
    serverNameSpan = document.createElement("span");
    serverName = document.createTextNode(server);
    serverChannelsDiv = document.createElement("div");
    channelNameLink = document.createElement("a");
    channelName = document.createTextNode(channel);

    chans.appendChild(column).classList.add("server");

    column.appendChild(serverNameSpan).classList.add("server-name");
    serverNameSpan.appendChild(serverName);

    column.appendChild(serverChannelsDiv).classList.add("server-channels");
    serverChannelsDiv.appendChild(line);
    line.appendChild(channelNameLink).classList.add("active");
    channelNameLink.appendChild(channelName);

    scrollToBottom("buffer");
}


//// exec
document.addEventListener('DOMContentLoaded', function() {
    scrollToBottom("buffer");
});
