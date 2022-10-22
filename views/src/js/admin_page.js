let ws = new WebSocket("ws://" + window.location.host + "admin/messages/ws")

let stats;

let errors;
let messages;
let users;

window.onload = (e) => {
    let response = fetch('/admin/stats').then(r => r.json).then(r => stats = r)
}

// getMessages  = 1
// getUsers     = 2
// setUsers     = 3
// modifyUsers  = 4

// errorMessage = 5
// newMessage   = 6

ws.onmessage = (msg) => {
    if (msg.type === '5') {
        let error = msg.
        errors.appendChild();
    } else if (msg.type === '6') {

    } else {

    }
}


