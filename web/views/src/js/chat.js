const setMessages = (json) => {
    for (let el in json) {
        let message = json[el];
        let line =  "User: " + message.user + " - " + message.text + "\n";
        chat.innerText += line;
    }
};

window.onload = (e) => {
    let r = fetch("/id/user/" + getCookie("email")).then(r => r.json()).then(r => setCookie("user-id", r.id, {'max-age': 2678400}));
    let r_ = fetch("/id/chat/" + window.location.pathname.split("/")[window.location.pathname.split("/").length - 1]).then(r => r.json()).then(r => setCookie("chat-id", r.id, {'max-age': 2678400}));

    let messages = fetch("/channel/massages/" + getCookie("chat-id") + "/" + "0").then(r => r.json()).then(r => setMessages(r))
};

let url = "ws://" + window.location.host + window.location.pathname + "/ws";
let ws = new WebSocket(url);
let channelId = window.location.pathname.split("/")[window.location.pathname.split("/").length - 1];

let chat = document.getElementById("chatLog");
let text = document.getElementById("input");

ws.onmessage = function (msg) {
    let mes = JSON.parse(msg.data)
    let m = document.createElement('div');
    m.className = '.user2';
    m.innerText = mes.text;

    console.log(m);
    console.log(chat);
    console.log(text);

    chat.appendChild(m);
};

function sendMessage() {
    let msg = {
        user: getCookie('user-id'),
        chat: channelId,
        text: text.value,
    }
    ws.send(JSON.stringify(msg));
    console.log(JSON.stringify(msg))
    text.value = "";
}

text.onkeydown = function (e) {
    if (e.keyCode === 13 && text.value !== "") {
        sendMessage()
    }
};