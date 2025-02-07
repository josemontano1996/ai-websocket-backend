let socket = new WebSocket('ws://localhost:3000/ws');
socket.onmessage = (event) => {
  console.log('received event from server this event', event.data);
};
let socket = new WebSocket('ws://localhost:3000/ws');
