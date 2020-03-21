function scrollToBottom(id){
   var element = document.getElementById(id);
   element.scrollTop = element.scrollHeight;
}

document.addEventListener('DOMContentLoaded', function() {
    scrollToBottom("buffer");
});
