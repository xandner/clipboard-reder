const net = require("net");

// Replace 'localhost' with your server IP if it's running on a different machine
const client = net.connect({ port: 9999, host: "localhost" }, () => {
  console.log("Connected to server!");

  // Sending data to server
  client.write("Hello from JavaScript client!\n");
});

client.on("data", (data) => {
  console.log("Received:", data.toString());
});

client.on("end", () => {
  console.log("Disconnected from server");
});
